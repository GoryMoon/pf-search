<template>
    <div class="modal modal-open" v-if="show && spell !== null" @mousedown="hideModal">
        <div class="modal-box w-11/12 max-w-7xl relative" ref="modalBox">
            <label for="my-modal-3" class="btn btn-sm btn-circle absolute right-2 top-2" @click="close">✕</label>
            <span class="title text-4xl font-bold">{{ spell.name }}</span>
            <div class="tooltip" data-tip="Open on d20PFSRD">
                <a class="btn btn-primary btn-xs ml-2" :href="spell.url" target="_blank">
                    d20PFSRD<ExternalLinkIcon class="ml-1 inline-block w-4 h-4 stroke-current"/>
                </a>
            </div>
            <div class="divider"></div>
            <div class="w-fit">
                <span class="font-bold">School:</span>
                <span class="hover:underline underline-offset-1">
                    <span class="ml-1">{{ _capitalize(spell.school.school) }}</span>
                    <span v-if="spell.school.sub_school !== null" class="ml-1">({{ _capitalize(spell.school.sub_school) }})</span>
                    <span v-if="spell.school.descriptors !== null && spell.school.descriptors.length > 0" class="ml-1">[{{ spell.school.descriptors.join(', ') }}]</span>
                </span>
            </div>
            <div class="w-fit pb-2">
                <span class="font-bold">Level:</span>
                <span class="ml-1 underline-offset-1">
                    <span v-for="name, index in Object.keys(spell.classes)" :key="name">
                        <template v-if="index > 0">, </template>
                        <span class="underline">{{ _startCase(name) }}</span> {{ spell.classes[name] }}
                    </span>
                </span>
            </div>
            <div v-if="hasCasting" class="mb-2">
                <span class="text-lg border-b-2 border-black">Casting</span>
                <div v-if="spell.casting_time.trim() !== ''">
                    <span class="font-bold">Casting time:</span> {{ spell.casting_time }}
                </div>
                <div v-if="hasComponent">
                    <span class="font-bold">Component:</span>
                    <span class="ml-1">
                        <span v-for="comp, index in components">
                            <template v-if="index > 0">, </template>
                            <div class="tooltip underline decoration-dotted hover:decoration-solid underline-offset-1" :data-tip="comp.name">
                                {{ comp.label }}
                            </div>
                        </span>
                    </span>
                </div>
            </div>
            <div v-if="hasEffect" class="mb-2">
                <span class="text-lg border-b-2 border-black">Effect</span>
                <div v-if="spell.effect.range !== null">
                    <span class="font-bold">Range:</span> {{ _capitalize(spell.effect.range) }}
                </div>
                <div v-if="spell.effect.area !== null">
                    <span class="font-bold">Area:</span> {{ _capitalize(spell.effect.area) }}
                </div>
                <div v-if="spell.effect.target !== null">
                    <span class="font-bold">Target:</span> {{ _capitalize(spell.effect.target) }}
                </div>
                <div v-if="spell.effect.duration !== null">
                    <span class="font-bold">Duration:</span> {{ _capitalize(spell.effect.duration) }}
                </div>
                <div v-if="spell.saving_throw.description !== null">
                    <span class="font-bold">Saving Throw:</span> {{ _capitalize(spell.saving_throw.description) }}
                </div>
                <div v-if="spell.spell_resistance.description !== null">
                    <span class="font-bold">Allow Spell Resistance:</span> {{ _capitalize(spell.spell_resistance.description) }}
                </div>
            </div>

            <span class="text-lg border-b-2 border-black w-fit">Description</span>
            <span class="block">{{ spell.description }}</span>

            <div v-if="spell.source_book.trim() !== ''" class="mt-2">
                <span class="text-lg border-b-2 border-black w-fit">Source</span>
                <span class="block">{{ spell.source_book }}</span>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { Spell as SpellType } from '~/types';
import { ExternalLinkIcon } from '@heroicons/vue/outline'
import _startCase from 'lodash/startCase';
import _capitalize from 'lodash/capitalize';

interface Component {
    label: string
    state: boolean
    name: string
}

const props = defineProps<{
    spell?: SpellType,
    show: boolean
}>()
const emit = defineEmits(['update:show'])
const modalBox = ref<HTMLDivElement>(null)

const show = computed<boolean>({
    get() {
        return props.show
    }, set(val: boolean) {
        emit('update:show', val)
    }
})
const components = computed<Component[]>(() => {
    return props.spell !== null ? [
        {label: 'V', state: props.spell.components.verbal, name: 'Verbal'},
        {label: 'S', state: props.spell.components.somatic, name: 'Somatic'},
        {label: 'DF', state: props.spell.components.divine_focus, name: 'Divine Focus'},
        {label: props.spell.components.material, state: props.spell.components.material !== null, name: 'Material'},
        {label: props.spell.components.focus, state: props.spell.components.focus !== null, name: 'Focus'}
    ]
    .filter((c) => c.state): []
})
const hasComponent = computed<boolean>(() => {
    return props.spell !== null && (props.spell.components.verbal || props.spell.components.divine_focus || props.spell.components.somatic)
})
const hasCasting = computed<boolean>(() => {
    return props.spell !== null && (props.spell.casting_time.trim() !== '' || hasComponent.value)
})
const hasEffect = computed<boolean>(() => {
    return props.spell !== null && (
        props.spell.effect.range != null ||
        props.spell.effect.area !== null ||
        props.spell.effect.target != null ||
        props.spell.effect.duration !== null ||
        props.spell.effect.dismissible ||
        props.spell.saving_throw.description !== null ||
        props.spell.spell_resistance.description !== null
    )
})


function hideModal(e: MouseEvent) {
    if (e.target == modalBox.value || (e.target instanceof Element && modalBox.value.contains(e.target)))
        return

    close()
}

function close() {
    show.value = false
}

function onKeyDown(e: KeyboardEvent) {
    if (show.value && e.key == 'Escape') {
        close()
    }
}

onMounted(() => document.addEventListener('keydown', onKeyDown))
onBeforeUnmount(() => document.removeEventListener('keydown', onKeyDown))

</script>
<style>

.title {
    font-family: 'Balthazar', 'Inter', 'Tahoma', Geneva, Verdana, sans-serif;
}

</style>