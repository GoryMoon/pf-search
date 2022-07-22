<template>
    <div class="navbar bg-base-300">
        <div class="flex-1">
            <a class="pl-4 font-bold normal-case text-2xl title">Pathfinder&nbsp;Search<span class="align-super text-sm pl-1">Alpha</span></a>
        </div>
        <div class="flex-none">
            <div title="Change Theme" class="dropdown dropdown-end">
                <div tabindex="0" class="btn gap-1 normal-case btn-ghost">
                    <ColorSwatchIcon class="inline-block h-5 w-5 stroke-current md:h-6 md:w-6" />
                    <span class="hidden md:inline">Theme</span>
                    <ChevronDownIcon class="ml-1 hidden h-4 w-4 opacity-70 sm:inline-block"/>
                </div>
                <div
                    class="dropdown-content bg-base-200 text-base-content rounded-t-box rounded-b-box top-px max-h-96 h-[70vh] w-52 overflow-y-auto shadow-2xl mt-16">
                    <div class="grid grid-cols-1 gap-3 p-3" tabindex="0">
                        <div
                            v-for="theme of themes"
                            :key="theme"
                            v-on:click="colorMode.preference = theme"
                            :class="['outline-base-content overflow-hidden rounded-lg outline-2 outline-offset-2', { outline: colorMode.preference == theme }]"
                            :data-set-theme="theme" data-act-class="outline">
                            <div :data-theme="theme"
                                class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                                <div class="grid grid-cols-5 grid-rows-3">
                                    <div class="col-span-5 row-span-3 row-start-1 flex gap-1 py-3 px-4">
                                        <div class="flex-grow text-sm font-bold">{{ theme }}</div>
                                        <div class="flex flex-shrink-0 flex-wrap gap-1">
                                            <div class="bg-primary w-2 rounded"></div>
                                            <div class="bg-secondary w-2 rounded"></div>
                                            <div class="bg-accent w-2 rounded"></div>
                                            <div class="bg-neutral w-2 rounded"></div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ColorModeInstance } from '@nuxtjs/color-mode/dist/runtime/types'
import { ColorSwatchIcon, ChevronDownIcon } from '@heroicons/vue/outline'
import { Ref } from 'vue';

const themes: Ref<string[]> = ref()
const colorMode: Ref<ColorModeInstance> = ref()

colorMode.value = useColorMode()
themes.value = ["light", "dark", "cupcake", "bumblebee", "emerald", "corporate", "synthwave", "retro", "cyberpunk", "valentine", "halloween", "garden", "forest", "aqua", "lofi", "pastel", "fantasy", "wireframe", "black", "luxury", "dracula", "cmyk", "autumn", "business", "acid", "lemonade", "night", "coffee", "winter"]
</script>
<style>

.title {
    font-family: 'Balthazar', 'Inter', 'Tahoma', Geneva, Verdana, sans-serif;
}

</style>