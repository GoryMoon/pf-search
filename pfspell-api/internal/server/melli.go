package server

import (
	"encoding/json"
	"github.com/gobeam/stringy"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog/log"
	"os"
	"pfspell-api/internal/utils"
	"regexp"
)

var noneSavingThrowRegex = regexp.MustCompile("(no(ne)?)")

func (s *Server) initMelli() {
	log.Info().Str("ctx", "meili").Msg("Loading spell file")
	f, err := os.ReadFile("spells.json")
	if err != nil {
		log.Fatal().Str("ctx", "meili").Err(err).Msg("Error reading file")
	}

	var spells []utils.Spell
	err = json.Unmarshal(f, &spells)
	if err != nil {
		log.Fatal().Str("ctx", "meili").Err(err).Msg("Error parsing file")
	}
	log.Info().Str("ctx", "meili").Msgf("Spell count: %d", len(spells))

	for i, spell := range spells {
		spell.Id = utils.CleanString(stringy.New(spell.Name).SnakeCase().ToLower())
		if spell.SpellResistance.Description == nil {
			spell.SpellResistance.NoneSet = true
		}
		if !spell.SavingThrow.Will &&
			!spell.SavingThrow.Fortitude &&
			!spell.SavingThrow.Reflex &&
			(spell.SavingThrow.Description == nil || noneSavingThrowRegex.MatchString(*spell.SavingThrow.Description)) {
			spell.SavingThrow.None = true
		}
		spells[i] = spell
	}

	log.Info().Str("ctx", "meili").Msg("Loading spell into meili")
	index := s.meili.Index("spells")
	_, _ = index.UpdateFilterableAttributes(&[]string{
		"school.school",
		"school.sub_school",
		"school.descriptors",
		"classes",
		"effect.range",
		"effect.dismissible",
		"spell_resistance.applies",
		"spell_resistance.none_set",
		"spell_resistance.description",
		"saving_throw.fortitude",
		"saving_throw.reflex",
		"saving_throw.will",
		"saving_throw.none",
		"source_book",
		"related_spell_names",
	})

	task, err := index.AddDocuments(spells, "id")
	if err != nil {
		log.Fatal().Str("ctx", "meili").Err(err).Msg("Error loading spells into meili")
	}
	log.Info().Str("ctx", "meili").Msg("Done loading spells into meili")
	str, _ := json.Marshal(task)
	log.Info().Str("ctx", "meili").Msg(string(str))
}

func setupMelli(melliHost string, melliKey string) *meilisearch.Client {
	return meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   melliHost,
		APIKey: melliKey,
	})
}
