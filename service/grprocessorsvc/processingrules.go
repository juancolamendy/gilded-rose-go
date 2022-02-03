package grprocessorsvc

import (
	"strings"
	"log"

	"github.com/juancolamendy/gilded-rose-go/utils/mathutils"
)

const (
	AGED_BRIE = "aged brie"
	LEGENDARY = "legendary sulfuras, hand of ragnaros"
	BACKSTAGE = "backstage passes to a tafkal80etc concert"
)

type ruleFn func(item *Item)

var (
	rules = map[string]ruleFn {
		AGED_BRIE: defaultRule,
		LEGENDARY: defaultRule,
		BACKSTAGE: defaultRule,
	}
)

func adjust(value, adjust, lowBound, highBound int) int {
	calc := value + adjust
	calc = mathutils.Min(calc, highBound)
	calc = mathutils.Max(calc, lowBound)
	return calc
}

func getProcessingRule(item *Item) ruleFn {
	// get key
	key := strings.ToLower(item.name)
	// look up the rule
	if rule, ok := rules[key]; ok {
		return rule
	}
	// fallback to default rule
	return defaultRule
}

func defaultRule(item *Item) {
	log.Printf("--- Applying default rule to item: %s", item.String())
}