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
		AGED_BRIE: defaultAgedBrie,
		LEGENDARY: defaultLegendary,
		BACKSTAGE: defaultBackstage,
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
	log.Printf("--- Applying 'default' rule to item: %s", item.String())	
	// default adjustment
	adjustment := -1
	if item.days <= 0 {
		adjustment = -2
	}
	// apply rules
	// adjust day
	item.days = item.days - 1
	// adjust quality
	item.quality = adjust(item.quality, adjustment, 0, 50)
}

func defaultAgedBrie(item *Item) {
	log.Printf("--- Applying 'aged brie' rule to item: %s", item.String())	
	// default adjustment
	adjustment := 1
	if item.days <= 0 {
		adjustment = 2
	}
	// apply rules
	// adjust day
	item.days = item.days - 1
	// adjust quality
	item.quality = adjust(item.quality, adjustment, 0, 50)
}

func defaultLegendary(item *Item) {
	log.Printf("--- Applying 'legendary' rule to item: %s", item.String())
	// noop
}

func defaultBackstage(item *Item) {
	log.Printf("--- Applying 'backstage' rule to item: %s", item.String())
	// apply rules
	// default adjustment
	adjustment := 0
	if item.days > 10 {
		adjustment = 1
	} else if item.days > 5 {
		adjustment = 2
	} else if item.days > 0 {
		adjustment = 3
	}
	// adjust day
	item.days = item.days - 1
	// adjust quality
	if adjustment == 0 {
		item.quality = 0
		return
	}	
	item.quality = adjust(item.quality, adjustment, 0, 50)
}