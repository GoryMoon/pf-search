<template>
    <div class="bg-base-100 h-full min-h-screen">
        <Navbar />
        <div class="container mx-auto px-2">
            <h1 class="text-6xl text-center mt-4 mb-6 title">Spells</h1>
            <div class="w-full grid justify-items-center">
                <input v-model.trim="searchText" type="text" placeholder="Search…" class="input input-bordered md:input-lg max-w-lg w-full" />
            </div>
            <div class="flex flex-col lg:flex-row lg:flex-wrap gap-y-4 items-center pt-6">
                <PfCheckboxDropdown name="Range" v-model="range" @update:model-value="refreshData" :options="availableRanges" />
                <div class="divider divider-horizontal"></div>
                <PfSelect placeholder="Class" v-model="classes" @update:model-value="updateClasses" :options="availableClasses" />
                <PfCheckboxDropdown name="Class Level" class="ml-2" v-model="classLevels" @update:model-value="updateClassLevels" :disabled="availableClassLevels.length <= 0" :options="availableClassLevels" />
                <div class="divider divider-horizontal"></div>
                <PfCheckbox v-model="spellResistance" indeterminable @update:model-value="refreshData" />
                <div class="divider divider-horizontal"></div>
                <PfCheckboxDropdown name="Saving Throw" v-model="savingThrow" @update:model-value="refreshData" :options="availableSavingThrows" />
                <div class="divider divider-horizontal"></div>
                <PfSelect placeholder="School" v-model="schools" @update:model-value="refreshData" :options="availableSchools" />
                <div class="divider divider-horizontal"></div>
            </div>
            <div class="divider"></div>
            <div v-if="pending" class="flex items-center justify-center">
                <PfSpinner />
            </div>
            <div v-if="!pending && searchResult != null" class="pb-10">
                <span class="text-sm">Matches: {{ searchResult.estimatedTotalHits }}</span>
                <SpellTable class="pt-4" :spells="searchResult.hits" :class-filter="classes.length > 0 ? classes[0].key.substring(8): ''"/>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import _throttle from 'lodash/throttle'
import _filter from 'lodash/filter'
import _startCase from 'lodash/startCase'
import _unionWith from 'lodash/unionWith'
import { CheckboxItem, CheckboxState, FacetValue, Params, Result, SelectOption } from './types';

const searchText = ref<string>('')
const range = ref<SelectOption[]>([])
const classes = ref<SelectOption[]>([])
const classLevels = ref<SelectOption[]>([])
const spellResistance = ref<CheckboxItem>({selected: CheckboxState.Unchecked, text: 'Spell Resistance'})
const savingThrow = ref<SelectOption[]>([])
const schools = ref<SelectOption[]>([])

const availableClasses = ref<SelectOption[]>([])
const availableClassLevels = ref<SelectOption[]>([])

const offset = ref(0)
const limit = ref(100)

let prevClass = ref<string>('')
let prevLevels: SelectOption[] = []

function filterParams(params: Params): Params {
    const out = {}

    for (const key in params) {
        const element = params[key];
        if (element != null && ((typeof(element) == "string" && element != '') || !(typeof(element) == "string")))
            out[key] = element
    }

    return out
}


const config = useRuntimeConfig()
const { pending, data: searchResult } = useLazyAsyncData<Result>('search', () => {        
    return $fetch('/v1/spells/search', { baseURL: config.API_BASE_URL, params: filterParams({
        s: searchText.value,
        class: classes.value.length > 0 ? classes.value[0].key.substring(8): '',
        class_levels: classLevels.value.length > 0 ? classLevels.value.map((v) => v.key).join(',') : '',
        range: range.value.length > 0 ? range.value.map((v) => v.key).join(',') : '',
        spell_resistance: spellResistance.value.selected == CheckboxState.Checked ? 1: spellResistance.value.selected == CheckboxState.Unchecked ? '': 0,
        saving_throw: savingThrow.value.length > 0 ? savingThrow.value.map((v) => v.key).join(',') : '',
        school: schools.value.length > 0 ? schools.value[0].key: '',
        sub_school: '',
        descriptor: '',
        limit: limit.value,
        offset: offset.value
    })})
})

const throttledSearch = _throttle(() => refreshNuxtData('search'), 500, { leading: false, trailing: true })
watch(searchText, () => throttledSearch())
onMounted(() => throttledSearch())

function refreshData() {
    refreshNuxtData('search')
}

function updateClasses() {
    prevClass.value = ''
    availableClassLevels.value = classLevels.value = []
    refreshData()
}

function updateClassLevels() {
    refreshData()
}

const rangeTypes = ['close', 'medium', 'long', 'personal', 'touch']
const rangeDefaults: SelectOption[] = []
rangeTypes.forEach((v) => {
    rangeDefaults.push({ name: _startCase(v), key: v, count: 0 })
})

const availableRanges = computed(() => {
    if (searchResult.value == null)
        return []

    let ranges: SelectOption[] = []
    const rangesData: FacetValue = searchResult.value.facetDistribution['effect.range']
    
    for (const key of rangeTypes) {
        const name = _startCase(key)
        const count = rangesData[key]
        ranges.push({ name: name, key, count })
    }
    
    const otherKeys: string[] = _filter(Object.keys(rangesData), (s: string) => !rangeTypes.includes(s))


    let otherCount = 0
    otherKeys.forEach((k: string) => otherCount += rangesData[k])
    ranges.push({ name: 'Other', key: 'other', count: otherCount})
    
    ranges = _unionWith(ranges, rangeDefaults, (a: SelectOption, b: SelectOption) => a.key === b.key)

    return ranges
})


const savingThrowTypes = ['none', 'will', 'reflex', 'fortitude']
const savingThrowDefaults: SelectOption[] = []
savingThrowTypes.forEach((v) => {
    savingThrowDefaults.push({ name: _startCase(v), key: v, count: 0 })
})
const availableSavingThrows = computed(() => {
    if (searchResult.value == null)
        return []

    let anyCount = 0
    let savingThrows: SelectOption[] = []
    for (const key of savingThrowTypes) {
        if (searchResult.value.facetDistribution['saving_throw.' + key] == undefined)
            continue

        const name = _startCase(key)
        const count = searchResult.value.facetDistribution['saving_throw.' + key]['true']
        if (key !== 'none')
            anyCount += count
        savingThrows.push({ name: name, key, count })
    }

    savingThrows.push({ name: 'Any set', key: 'any', count: anyCount})
    savingThrows = _unionWith(savingThrows, savingThrowDefaults, (a: SelectOption, b: SelectOption) => a.key === b.key)

    return savingThrows
})

const schoolTypes = ['abjuration', 'conjuration', 'divination', 'enchantment', 'evocation', 'illusion', 'necromancy', 'transmutation']
const schoolDefaults: SelectOption[] = []
schoolTypes.forEach((v) => {
    schoolDefaults.push({ name: _startCase(v), key: v, count: 0 })
})
const availableSchools = computed(() => {
    if (searchResult.value == null)
        return []

    let schools: SelectOption[] = []
    const rangesData: FacetValue = searchResult.value.facetDistribution['school.school']
    
    for (const key of schoolTypes) {
        const name = _startCase(key)
        const count = rangesData[key]
        schools.push({ name: name, key, count })
    }

    schools = _unionWith(schools, schoolDefaults, (a: SelectOption, b: SelectOption) => a.key === b.key)

    return schools
})

watch(searchResult, () => {
    if (searchResult.value == null) {
        availableClasses.value = []
        availableClassLevels.value = []
        return
    }

    let classList: SelectOption[] = []
    const classNames: string[] = _filter(Object.keys(searchResult.value.facetDistribution), (v) => v.startsWith('classes.'))
    const data = searchResult.value.facetDistribution

    for (const key of classNames) {
        const name = _startCase(key.substring(8))        
        const count = Object.values(data[key]).reduce((p, c) => p + c, 0)
        classList.push({ name: name, key, count })
    }
    
    availableClasses.value = classList

    // Class levels
    if (classes.value.length == 0) {
        availableClassLevels.value = prevLevels = []
        prevClass.value = ''
        return
    }

    var defaults = []
    const clazz = classes.value[0].key
    if (clazz == prevClass.value)
        defaults = prevLevels

    let levels: SelectOption[] = []
    const rangesData: FacetValue = searchResult.value.facetDistribution[clazz]

    for (const key in rangesData) {
        const name = key
        const count = rangesData[key]
        levels.push({ name: name, key, count })
    }

    console.log(defaults, levels);
    
    defaults.forEach((v: SelectOption) => v.count = 0)
    levels = _unionWith(levels, defaults, (a: SelectOption, b: SelectOption) => a.key === b.key)
    levels.sort((a: SelectOption, b: SelectOption) => a.name.localeCompare(b.name))

    prevClass.value = clazz
    prevLevels = levels
    availableClassLevels.value = levels
})

useHead({
    title: 'Pathfinder Search',
    meta: [
        { name: 'description', content: 'A site to search Pathfinder spells' }
    ]
})

</script>
<style>

.title {
    font-family: 'Balthazar', 'Inter', 'Tahoma', Geneva, Verdana, sans-serif;
}

</style>