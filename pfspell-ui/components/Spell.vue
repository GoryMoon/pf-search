<template>
    <div class="modal modal-open" v-if="show && spell !== null" @mousedown="hideModal">
        <div class="modal-box w-11/12 max-w-7xl relative" ref="modalBox">
            <label for="my-modal-3" class="btn btn-sm btn-circle absolute right-2 top-2" @click="close">✕</label>
            <a class="link" :href="spell.url" target="_blank">
                <h1 class="title text-4xl">{{ spell.name }}</h1>
            </a>
            <div class="text-base">{{ spell.description }}</div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { Spell as SpellType } from '~/types';


const props = defineProps<{
    spell?: SpellType,
    show: boolean
}>()
const emit = defineEmits(['close'])

const show = computed<boolean>({
    get() {
        return props.show
    }, set(val: boolean) {
        emit('close', val)
    }
})
const modalBox = ref<HTMLDivElement>(null)

function hideModal(e: MouseEvent) {
    if (e.target == modalBox.value || modalBox.value.contains(e.target))
        return

    close()
}

function close() {
    show.value = false
}

</script>
<style>

.title {
    font-family: 'Balthazar', 'Inter', 'Tahoma', Geneva, Verdana, sans-serif;
}

</style>