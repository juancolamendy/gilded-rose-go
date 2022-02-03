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
	CONJURED  = "conjured mana cake"
)

type ruleFn func(item *Item)

var (
	rules = map[string]ruleFn {
		AGED_BRIE: defaultAgedBrie,
		LEGENDARY: defaultLegendary,
		BACKSTAGE: defaultBackstage,
		CONJURED:  conjuredtRule,
	}
)

func adjust(value, adjust, lowBound, highBound int) int {
	calc := value + adjust
	calc = mathutils.Min(calc, highBound)
	calc = mathutils.Max(calc, lowBound)
	return calc
}

func adjustDayQuality(item *Item, adjustment int) {
	// adjust day
	item.days -= 1
	// adjust quality
	item.quality = adjust(item.quality, adjustment, 0, 50)	
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
	// apply rules
	// default adjustment
	adjustment := -1
	// calc adjustment
	if item.days <= 0 {
		adjustment = -2
	}
	// adjust day and quality
	adjustDayQuality(item, adjustment)
}

func defaultAgedBrie(item *Item) {
	log.Printf("--- Applying 'aged brie' rule to item: %s", item.String())	
	// apply rules
	// default adjustment
	adjustment := 1
	// calc adjustment
	if item.days <= 0 {
		adjustment = 2
	}
	// adjust day and quality
	adjustDayQuality(item, adjustment)
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
	// calc adjustment
	switch {
	case item.days > 10:
		adjustment = 1
	case item.days > 5:
		adjustment = 2
	case item.days > 0:
		adjustment = 3
	}
	// adjust quality for zero case
	if adjustment == 0 {
		item.days -= 1
		item.quality = 0
		return
	}
	// adjust day and quality
	adjustDayQuality(item, adjustment)
}

func conjuredtRule(item *Item) {
	log.Printf("--- Applying 'conjured' rule to item: %s", item.String())
	// apply rules
	// default adjustment
	adjustment := -2
	// calc adjustment
	if item.days <= 0 {
		adjustment = -4
	}
	// adjust day and quality
	adjustDayQuality(item, adjustment)
}