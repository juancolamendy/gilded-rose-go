package grprocessorsvc

import (
	"log"
)

func Process(items ...*Item) {
	log.Printf("--- Start processing items. length %d", len(items))

	// loop through the items
	for _, item := range items {
		rule := getProcessingRule(item)
		rule(item)
	}
}