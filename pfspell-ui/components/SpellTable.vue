<template>
    <div class="overflow-x-auto pt-4" v-if="spells.length > 0">
        <table class="table table-zebra table-compact w-full">
            <thead>
                <tr>
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
    <Spell :spell="selectedSpell" v-model:show="showSpell" />
</template>
<script setup lang="ts">
import _startCase from 'lodash/startCase'
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