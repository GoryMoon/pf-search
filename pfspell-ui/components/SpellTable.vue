<template>
    <div class="overflow-x-auto pt-4" v-if="spells.length > 0">
        <table class="table table-zebra table-compact w-full">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Range</th>
                    <th class="max-w-xs">Target</th>
                    <th class="max-w-xl">Classes</th>
                    <th class="max-w-xs">Spell Resistance</th>
                    <th class="max-w-xs">Saving Throw</th>
                    <th>School</th>
                    <th>Subschool</th>
                </tr>
            </thead>
            <tbody>
                <tr
                    v-for="spell of sortedSpells"
                    :key="spell.id"
                    class="hover cursor-pointer"
                    @click="openSpellInfo(spell)"
                    >
                    <td class="title text-lg">{{ spell.name }}</td>
                    <td>{{ _capitalize(spell.effect.range) }}</td>
                    <td class="whitespace-normal max-w-xs">{{ _capitalize(spell.effect.target) }}</td>
                    <td class="whitespace-normal max-w-xl">{{ formatClasses(spell.classes) }}</td>
                    <td class="whitespace-normal max-w-xs">{{ spell.spell_resistance.description }}</td>
                    <td class="whitespace-normal max-w-xs">{{ spell.saving_throw.description }}</td>
                    <td>{{ _capitalize(spell.school.school) }}</td>
                    <td>{{ _capitalize(spell.school.sub_school) }}</td>
                </tr>
            </tbody>
        </table>
    </div>
    <Spell :spell="selectedSpell" v-model:show="showSpell" />
</template>
<script setup lang="ts">
import _startCase from 'lodash/startCase'
import _capitalize from 'lodash/capitalize'
import { Classes, Spell as SpellType } from '~/types';

const { spells, classFilter = '' } = defineProps<{
    spells: SpellType[],
    classFilter?: string
}>()

const selectedSpell = ref<SpellType>(null)
const showSpell = ref<boolean>(false)

const sortedSpells = computed<SpellType[]>(() => {
    if (classFilter !== '') {
        spells.sort((a: SpellType, b: SpellType) => a.classes[classFilter] - b.classes[classFilter])
    }
    return spells
})

function openSpellInfo(spell: SpellType) {
    selectedSpell.value = spell
    showSpell.value = true
}

function formatClasses(classes: Classes) {
    const out: string[] = []
    Object.entries(classes).forEach(([key, level]) => {
        out.push(`${_startCase(key)}: ${level}`)
    })
    return out.join(', ')
}


</script>
<style>

.max-w-class {
    max-width: 35rem;

}

.title, th {
    font-family: 'Balthazar', 'Inter', 'Tahoma', Geneva, Verdana, sans-serif;
}

</style>