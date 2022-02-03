package main

import(
	"fmt"

	"github.com/juancolamendy/gilded-rose-go/service/grprocessorsvc"
)

var things = []struct {
	name    string
	sellIn  int
	quality int
}{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	{"Conjured Mana Cake", 3, 6},
}

func main(){
	var items []*grprocessorsvc.Item

	for _, thing := range things {
		item := grprocessorsvc.New(thing.name, thing.sellIn, thing.quality)
		items = append(items, item)
		fmt.Println(item)
	}

	grprocessorsvc.Process(items...)

	for _, item := range items {
		fmt.Println(item)
	}
}