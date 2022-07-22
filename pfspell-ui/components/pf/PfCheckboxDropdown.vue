<template>
    <div ref="dropdown" class="dropdown">
        <label ref="button" tabindex="0" :class="['btn', 'm-1', {'btn-disabled': disabled, 'btn-primary': modelValue.length > 0}]" >
            {{ name }}
            <ChevronDownIcon class="ml-1 hidden h-4 w-4 opacity-70 sm:inline-block"/>
            </label>
        <ul
            tabindex="0"
            role="listbox"
            class="dropdown-content menu p-2 shadow bg-base-100 rounded-box min-w-52 max-h-screen-4 overflow-y-auto">
            <li>
                <span class="whitespace-nowrap" @click="toggleAll">
                    Delselect / Select All
                </span>
            </li>
            <li
                v-for="(option) in optionValues"
                :key="option.opt.key">
                <PfCheckbox v-model="option.check"  @update:modelValue="(val) => checkboxUpdate(option.opt, val)">
                    <template v-slot:text>
                        <div class="w-full pl-2 gap-2 flex justify-between">
                            <span>{{ option.opt.name }}</span>
                            <span v-if="option.opt.count > 0" class="badge badge-primary">{{ option.opt.count }}</span>
                        </div>
                    </template>
                </PfCheckbox>
            </li>
        </ul>
    </div>
</template>
<script setup lang="ts">
import { ChevronDownIcon } from '@heroicons/vue/outline'
import _remove from 'lodash/remove'
import _find from 'lodash/find'
import { CheckboxItem, CheckboxState, SelectOption } from '~/types';

interface Option {
    opt: SelectOption,
    check: CheckboxItem
}

const { modelValue, options, name, disabled = false, multiple = false } = defineProps<{
    modelValue: SelectOption[],
    options: SelectOption[]
    name: string,
    disabled?: boolean
    multiple?: boolean
}>();
const dropdown = ref<HTMLDivElement>(null)
const button = ref<HTMLLabelElement>(null)

const optionValues = computed<Option[]>(() => {
    var opts: Option[] = []
    options.forEach(o => {
        const val = _find(modelValue, (v) => v.key === o.key)
        opts.push({ opt: o, check: { selected: val !== undefined ? CheckboxState.Checked: CheckboxState.Unchecked, text: ''} })
    });
    
    return opts
})

const emit = defineEmits(['update:modelValue'])
function checkboxUpdate(opt: SelectOption, check: CheckboxItem) {    
    if (check.selected == CheckboxState.Checked)
        modelValue.push({ name: opt.name, key: opt.key, count: opt.count })
    else if (check.selected == CheckboxState.Unchecked)
        _remove(modelValue, (o: SelectOption) => o.key == opt.key)
    
    emit('update:modelValue', modelValue)
}

watch(() => disabled, () => {
    dropdown.value?.blur()
    button.value?.blur()
})

function toggleAll() {
    if (modelValue.length > 0) {
        modelValue.splice(0)
    } else {
        options.forEach((opt) => {
            modelValue.push({ name: opt.name, key: opt.key, count: opt.count })
        })
    }
    emit('update:modelValue', modelValue)
}

</script>
<style scoped>

.max-h-screen-4 {
    max-height: 50vh;
}

</style>