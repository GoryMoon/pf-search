<template>
    <div :class="['dropdown', {'dropdown-open': open}]">
        <div
            class="input input-bordered cursor-text w-fit input-search inline-flex items-center"
            @mousedown="(event) => toggleDropdown(event)">
            <slot
                v-for="option in selectedValue"
                name="selected-options"
                :option="option"
                :deselect="deselect"
                :multiple="multiple">
                <span :key="option.key" class="badge p-3 mr-2 w-fit whitespace-nowrap">
                    {{ option.name }}
                    <button
                        v-if="multiple"
                        ref="deselectButtons"
                        type="button"
                        @click="deselect(option)">
                        <XIcon class="inline-block w-4 h-4 stroke-current" />
                    </button>
                </span>
            </slot>
            <input
                type="text"
                ref="searchInput"
                :placeholder="placeholder"
                @compositionstart="isComposing = true"
                @compositionend="isComposing = false"
                @blur="onSearchBlur"
                @focus="onSearchFocus"
                @keydown="onSearchKeyDown"
                v-model="search"
                class="input input-ghost max-w-input" />
            <button
                v-show="showClearButton"
                ref="clearButton"
                type="button"
                title="Clear Selected"
                aria-label="Clear Selected"
                class="cursor-pointer"
                @click="clearSelection">
                <XIcon class="inline-block w-4 h-4 stroke-current" />
            </button>
        </div>
        <ul
            v-if="open"
            @mousedown.prevent="mousedown = true"
            @mouseup="mousedown = false"
            tabindex="-1"
            role="listbox"
            class="dropdown-content menu w-full p-2 shadow bg-base-100 rounded-box max-h-screen-4 overflow-y-auto">
            <li
                v-for="(option, index) in filteredOptions"
                :key="option.key">
                <a @click.prevent.stop="select(option)" class="max-w-full gap-2 justify-between" >
                    <span>{{ option.name }}</span>
                    <span v-if="option.count > 0" class="badge badge-primary">{{ option.count }}</span>
                </a>
            </li>
            <li v-if="!taggable && filteredOptions.length === 0">
                Sorry, no matching options.
            </li>
        </ul>
    </div>
</template>
<script setup lang="ts">
import { XIcon } from '@heroicons/vue/outline'
import { SelectOption } from '~/types';

var deselectButtons = ref<HTMLButtonElement[]>([])
var clearButton = ref<HTMLButtonElement>()
var searchInput = ref<HTMLInputElement | null>(null)

var search = ref<string>('')
var open = ref(false)
var isComposing = ref(false)
var pushedTags = ref<SelectOption[]>([])
var _value = ref<SelectOption[]>([])

var searching = ref(false)
var mousedown = ref(false)


const { placeholder, modelValue, options, taggable = false,  pushTags = false, multiple = false } = defineProps<{
    placeholder: string,
    modelValue: SelectOption[],
    options: SelectOption[]
    taggable?: boolean
    pushTags?: boolean
    multiple?: boolean
}>();

const emit = defineEmits(['update:modelValue'])
const selectedValue = computed<SelectOption[]>({
    get() {
        let value = modelValue

        if (value !== undefined && value !== null)
            return [].concat(modelValue)
        return []
    },
    set(value: SelectOption[]) {
        if (typeof(modelValue) === 'undefined')
            _value.value = value
        
        emit('update:modelValue', value)
    }
})
const optionList = computed(() => options.concat(pushTags ? pushedTags.value: []))
const isValueEmpty = computed(() => selectedValue.value.length === 0)
const showClearButton = computed(() => !multiple && !isValueEmpty.value)

function optionExists(opt: SelectOption) {
    return optionList.value.some((_opt) => optionComparator(_opt, opt))
}

function optionComparator(a: SelectOption, b: SelectOption) {
    return a.key === b.key
}

function isOptionSelected(option: SelectOption) {
    return selectedValue.value.some((v) => optionComparator(v, option))
}

function select(option: SelectOption) {
    if (!isOptionSelected(option)) {
        if (taggable && !optionExists(option))
            pushedTags.value.push(option)

        if (multiple)
            selectedValue.value = selectedValue.value.concat(option)
        else
            selectedValue.value = [].concat(option)
        
    } else if ((multiple && selectedValue.value.length > 1)) {
        deselect(option)
    }

    open.value = !open.value
    searchBlur()
    search.value = ""
}

function deselect(option: SelectOption) {
    selectedValue.value = selectedValue.value.filter((v) => !optionComparator(v, option))
}

function clearSelection() {
    selectedValue.value = []
}

function maybeDeleteValue() {
    if (!search.value.length && selectedValue.value != null && selectedValue.value.length > 0) {
        selectedValue.value = [
            ...selectedValue.value.slice(0, selectedValue.value.length - 1)
        ]
    }
}

function toggleDropdown(event: MouseEvent) {
    const targetIsNotSearch = event.target !== searchInput.value
    if (targetIsNotSearch)
        event.preventDefault()

    const ignoreButtons = [
        ...(deselectButtons.value || []),
        ...([clearButton.value] || [])
    ]    

    if (ignoreButtons.filter(Boolean).some((r) => (event.target instanceof Element && r.contains(event.target)) || r === event.target)) {
        event.preventDefault()
        return
    }

    if (open.value && targetIsNotSearch)
        searchBlur()
    else {
        open.value = true
        searchFocus()
    }
}

function closeSearchOptions() {
    open.value = false
}

function onSearchKeyDown(e: KeyboardEvent) {
    const preventAndSelect = (e: KeyboardEvent) => {
        e.preventDefault()
        return !isComposing
    }

    const eventHandlers = {
        "Backspace": (e: KeyboardEvent) => maybeDeleteValue(),
        "Enter": preventAndSelect,
        "Escape": (e: KeyboardEvent) => onEscape(),
    }

    if (typeof(eventHandlers[e.key]) == "function")
        return eventHandlers[e.key](e)
}

function onEscape() {
    if (!search.value.length)
        searchBlur()
    else
        search.value = ''

}

function searchBlur() {
    searchInput.value?.blur()
}

function searchFocus() {
    searchInput.value?.focus()
}

function onSearchBlur() {
    if (mousedown.value && !searching.value)
        mousedown.value = false
    else {
        search.value = ''
        closeSearchOptions()
        return
    }

    if (search.value.length === 0 || options.length === 0)
        closeSearchOptions()
}

function onSearchFocus() {
    open.value = true
}

const filteredOptions = computed<SelectOption[]>(() => {
    const optionsList = [].concat(optionList.value)

    let options = search.value.length > 0
        ? optionsList.filter((opt) => (opt.name || '')
                .toLocaleLowerCase()
                .indexOf(search.value.toLocaleLowerCase()) > -1)
        : optionsList

    if (taggable && search.value.length) {
        const newVal: SelectOption = { name: search.value, key: search.value, count: 0}
        if (!optionExists(newVal))
            options.unshift(newVal)
    }

    return options
})

</script>
<style scoped>

.max-h-screen-4 {
    max-height: 50vh;
}

.max-w-input {
    max-width: 8rem;
}


.input-ghost {
    padding-left: 0;
    padding-right: 0;
}

.input-ghost:focus {
    outline: none;
    --tw-bg-opacity: 0;
}

.input-search {
    min-width: 16rem;
}

</style>