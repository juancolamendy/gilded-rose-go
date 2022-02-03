package main

import(
	"fmt"

	"github.com/juancolamendy/gilded-rose-go/service/grprocessorsvc"
)

func main(){
	item := grprocessorsvc.New("normal", 5, 10)
	grprocessorsvc.Process(item)
	fmt.Printf("Result item: %s\n", item.String())
}