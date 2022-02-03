package main

import(
	"fmt"

	"github.com/juancolamendy/gilded-rose-go/service/grprocessorsvc"
)

func main(){
	grprocessorsvc.Process()
	fmt.Println("Hello world")
}