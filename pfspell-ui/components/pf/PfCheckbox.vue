<template>
    <div class="form-control" @click.self="check = !check">
        <label class="flex cursor-pointer w-full items-center">
            <input type="checkbox" v-model="check" ref="checkbox" class="checkbox" />
            <span class="label-text w-full">
                <slot name="text">
                    <span class="pl-2">{{ modelValue.text }}</span>
                </slot>
            </span> 
        </label>
    </div>
</template>
<script setup lang="ts">
import { CheckboxItem, CheckboxState } from '~/types'

const { modelValue = {selected: CheckboxState.Unchecked, text: ''}, indeterminable = false, reverse = false } = defineProps<{
    modelValue?: CheckboxItem,
    indeterminable?: boolean
    reverse?: boolean
}>();

const emit = defineEmits(['update:modelValue'])
const checkbox = ref<HTMLInputElement>(null)
const state = ref<CheckboxState>(CheckboxState.Unchecked)
const check = computed<boolean>({
    get() {
        return state.value == CheckboxState.Checked || state.value == CheckboxState.Indeterminate
    },
    set() {
        const val = state.value
        const checked = val == CheckboxState.Checked
        const unchecked =  val == CheckboxState.Unchecked
        const indeterminate = val == CheckboxState.Indeterminate

        if (indeterminable && ((checked && !reverse) || (unchecked && reverse))) {
            state.value = CheckboxState.Indeterminate
            console.log("Set: Indeterminate");
        } else {
            state.value = checked || (indeterminate && !reverse) ? CheckboxState.Unchecked: CheckboxState.Checked
        }
        
        if (state.value == CheckboxState.Indeterminate) {
            checkbox.value.indeterminate = true
        } else {
            checkbox.value.indeterminate = false
            if (state.value == CheckboxState.Checked) {
                checkbox.value.checked = true
            }
        }

        modelValue.selected = state.value
        emit('update:modelValue', modelValue)
    }
})

watch(() => modelValue, () => {
    state.value = modelValue.selected
})

onMounted(() => {
    state.value = modelValue.selected
    if (state.value == CheckboxState.Indeterminate) {
        checkbox.value.indeterminate = true
    }
})

</script>