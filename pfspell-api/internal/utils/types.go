package utils

type School struct {
	School      string   `json:"school"`
	SubSchool   *string  `json:"sub_school"`
	Descriptors []string `json:"descriptors"`
}

type Components struct {
	Verbal      bool    `json:"verbal"`
	Somatic     bool    `json:"somatic"`
	Material    *string `json:"material"`
	Focus       *string `json:"focus"`
	DivineFocus bool    `json:"divine_focus"`
}

type Effect struct {
	Range       *string `json:"range"`
	Area        *string `json:"area"`
	Target      *string `json:"target"`
	Duration    *string `json:"duration"`
	Dismissible bool    `json:"dismissible"`
}

type SavingThrow struct {
	Fortitude   bool    `json:"fortitude"`
	Reflex      bool    `json:"reflex"`
	Will        bool    `json:"will"`
	None        bool    `json:"none"`
	Description *string `json:"description"`
}

type SpellResistance struct {
	Applies     bool    `json:"applies"`
	NoneSet     bool    `json:"none_set"`
	Description *string `json:"description"`
}

type Spell struct {
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	Link              string          `json:"url"`
	School            School          `json:"school"`
	Classes           map[string]int  `json:"classes"`
	CastingTime       string          `json:"casting_time"`
	Components        Components      `json:"components"`
	Effect            Effect          `json:"effect"`
	SavingThrow       SavingThrow     `json:"saving_throw"`
	SpellResistance   SpellResistance `json:"spell_resistance"`
	Description       string          `json:"description"`
	SourceBook        string          `json:"source_book"`
	RelatedSpellNames []string        `json:"related_spell_names"`
}
