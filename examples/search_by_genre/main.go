package main

import (
	"fmt"
	"log"

	"github.com/rushkii/nekopoi-go"
	"github.com/rushkii/nekopoi-go/filters"
)

func main() {
	// term_id: uint64 or []uint64 (e.g: 36 or []uint64{36, 37} or use defined value with filters.LOLI)
	res, err := nekopoi.SearchByGenre([]uint64{filters.LOLI, filters.SCHOOLGIRL})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(nekopoi.Prettify(&res))
}
