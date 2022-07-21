
interface School {
    school: string,
    sub_school: string,
    descriptors: string
}

interface Classes {
    [key: string]: number
}

interface Components {
    verbal: boolean,
    somatic: boolean,
    material: string|null,
    focus: string|null,
    divine_focus: boolean
}

interface Effect {
    range: string|null,
    area: string|null,
    target: string|null,
    duration: string|null,
    dismissible: boolean
}

interface SavingThrow {
    fortitude: boolean,
    reflex: boolean,
    will: boolean,
    description: string|null
}

interface SpellResistance {
    applies: boolean,
    description: string|null
}

interface Spell {
    id: string
    name: string,
    url: string,
    school: School,
    classes: Classes,
    casting_time: string,
    components: Components,
    effect: Effect,
    saving_throw: SavingThrow,
    spell_resistance: SpellResistance,
    description: string,
    source_book: string,
    related_spell_names: Array<string>
}

interface Result {
    hits: Array<Spell>,
    estimatedTotalHits: number,
    offset: number,
    limit: number,
    processingTimeMs: number,
    query: string,
    facetDistribution: Facet
}

interface Facet {
    [key: string]: FacetValue
}

interface FacetValue {
    [key: string]: number
}

interface Params {
    [key: string]: any
}

interface SelectOption {
    name: string,
    key: any,
    count: number
}

enum CheckboxState {
    Checked,
    Indeterminate,
    Unchecked
}

interface CheckboxItem {
    selected: CheckboxState,
    text: string
}

export { Params, SelectOption, CheckboxItem, CheckboxState, Spell, Classes, Result, Facet, FacetValue }