<template>
    <div class="py-4">
        <a class="link" :href="spell.url" target="_blank">
            <h1 class="title text-xl">{{ spell.name }}</h1>
        </a>
        <div class="text-base">{{ spell.description }}</div>
        
    </div>
</template>
<script setup lang="ts">
import { Spell } from '~/types';


const props = defineProps<{
    spell: Spell
}>()

</script>
<style>

.title {
    font-family: 'Balthazar', 'Inter', 'Tahoma', Geneva, Verdana, sans-serif;
}

</style>