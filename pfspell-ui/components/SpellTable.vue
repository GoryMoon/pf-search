<template>
    <div class="overflow-x-auto" v-if="spells.length > 0">
        <table class="table table-zebra table-compact w-full">
            <thead>
                <tr>
                    <th class="hidden"></th>
                    <th>Name</th>
                    <th>Range</th>
                    <th class="max-w-class">Classes</th>
                    <th>Spell Resistance</th>
                    <th>Saving Throw</th>
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
                    <td class="hidden"></td>
                    <td class="title text-lg">{{ spell.name }}</td>
                    <td>{{ _startCase(spell.effect.range) }}</td>
                    <td class="whitespace-normal max-w-class">{{ formatClasses(spell.classes) }}</td>
                    <td>{{ spell.spell_resistance.description }}</td>
                    <td>{{ spell.saving_throw.description }}</td>
                    <td>{{ _startCase(spell.school.school) }}</td>
                    <td>{{ _startCase(spell.school.sub_school) }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script setup lang="ts">
import _startCase from 'lodash/startCase'
import { Classes, Spell } from '~/types';

const { spells, classFilter = '' } = defineProps<{
    spells: Spell[],
    classFilter?: string
}>()

const sortedSpells = computed<Spell[]>(() => {
    if (classFilter !== '') {
        spells.sort((a: Spell, b: Spell) => a.classes[classFilter] - b.classes[classFilter])
    }
    return spells
})

function openSpellInfo(spell: Spell) {
    
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