package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meilisearch/meilisearch-go"
	"strconv"
	"strings"
)

var _rangeOtherQuery string

func rangeOtherQuery() string {
	if _rangeOtherQuery == "" {
		validRanges := []string{"close", "medium", "long", "personal", "touch"}
		subFilter := make([]string, len(validRanges))
		for i, check := range validRanges {
			subFilter[i] = fmt.Sprintf("effect.range != '%s'", check)
		}
		_rangeOtherQuery = strings.Join(subFilter, " AND ")
	}
	return _rangeOtherQuery
}

func (s *Server) setupRoutes() {
	s.fiber.Static("/", "./public", fiber.Static{
		Compress: true,
	})
	api := s.fiber.Group("/api")
	api.Get("/v1/spells/search", s.search)
}

func (s *Server) search(ctx *fiber.Ctx) error {
	index := s.meili.Index("spells")
	q := ctx.Query("s")

	limitString := ctx.Query("limit", "1000")
	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		limit = 1000
	}

	offsetString := ctx.Query("offset", "0")
	offset, err := strconv.ParseInt(offsetString, 10, 64)
	if err != nil {
		limit = 0
	}

	class := ctx.Query("class", "")
	rawClassLevels := ctx.Query("class_levels", "")
	rangeQuery := ctx.Query("range", "")
	spellResistance := ctx.Query("spell_resistance", "")
	savingThrowQuery := ctx.Query("saving_throw", "")
	school := ctx.Query("school", "")
	subSchool := ctx.Query("sub_school", "")
	descriptor := ctx.Query("descriptor", "")
	sourceBook := ctx.Query("source_book", "")

	var classLevels []int
	for _, rawLevel := range strings.Split(rawClassLevels, ",") {
		level, err := strconv.Atoi(rawLevel)
		if err == nil {
			classLevels = append(classLevels, level)
		}
	}

	ranges := strings.Split(rangeQuery, ",")
	savingThrow := strings.Split(savingThrowQuery, ",")

	stringFilter := [][]string{
		{"school.school", school},
		{"school.sub_school", subSchool},
		{"school.descriptors", descriptor},
		{"source_book", sourceBook},
	}

	var filter [][]string
	for _, str := range stringFilter {
		if str[1] != "" {
			filter = append(filter, []string{fmt.Sprintf("%s = '%s'", str[0], str[1])})
		}
	}

	if spellResistance != "" {
		state, _ := strconv.ParseInt(spellResistance, 10, 64)
		spellResistanceFilter := ""
		if state == -1 {
			spellResistanceFilter = "spell_resistance.applies = false AND spell_resistance.none_set = true"
		} else if state == 0 {
			spellResistanceFilter = "spell_resistance.applies = false"
		} else if state == 1 {
			spellResistanceFilter = "spell_resistance.applies = true"
		}

		if spellResistanceFilter != "" {
			filter = append(filter, []string{spellResistanceFilter})
		}
	}

	if savingThrowQuery != "" {
		var savingCheck []string
		for _, r := range savingThrow {
			if r == "any" {
				savingCheck = append(savingCheck, fmt.Sprintf("saving_throw.none = false"))
			} else {
				savingCheck = append(savingCheck, fmt.Sprintf("saving_throw.%s = true", r))
			}
		}
		filter = append(filter, []string{strings.Join(savingCheck, " OR ")})
	}

	if rangeQuery != "" {
		var rangeCheck []string
		for _, r := range ranges {
			if r == "other" {
				rangeCheck = append(rangeCheck, rangeOtherQuery())
			} else {
				rangeCheck = append(rangeCheck, fmt.Sprintf("effect.range = '%s'", r))
			}
		}
		filter = append(filter, []string{strings.Join(rangeCheck, " OR ")})
	}

	if class != "" {
		if len(classLevels) > 0 {
			subFilter := make([]string, len(classLevels))
			for i, level := range classLevels {
				subFilter[i] = fmt.Sprintf("classes.%s = %d", class, level)
			}
			filter = append(filter, subFilter)
		} else {
			filter = append(filter, []string{fmt.Sprintf("classes.%s > -1", class)})
		}
	}

	search, err := index.Search(q, &meilisearch.SearchRequest{
		Filter: filter,
		Limit:  limit,
		Offset: offset,
		Facets: []string{
			"classes",
			"school.school",
			"school.sub_school",
			"school.descriptors",
			"effect.range",
			"saving_throw.fortitude",
			"saving_throw.reflex",
			"saving_throw.will",
			"saving_throw.none",
			"source_book",
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(search)
}
