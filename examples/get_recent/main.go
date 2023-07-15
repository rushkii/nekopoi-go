package main

import (
	"fmt"
	"log"

	"github.com/rushkii/nekopoi-go"
)

func main() {
	res, err := nekopoi.GetRecent()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(nekopoi.Prettify(&res))
}
