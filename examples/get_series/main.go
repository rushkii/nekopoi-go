package main

import (
	"fmt"
	"log"

	"github.com/rushkii/nekopoi-go"
)

func main() {
	// id: uint64 (e.g: 9690) this is series ID, not hentai ID.
	res, err := nekopoi.GetSeries(9690)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(nekopoi.Prettify(&res))
}
