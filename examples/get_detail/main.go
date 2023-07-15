package main

import (
	"fmt"
	"log"

	"github.com/rushkii/nekopoi-go"
)

func main() {
	// id: uint64 (e.g: 27110) this is hentai ID, not series ID.
	res, err := nekopoi.GetDetail(27110)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(nekopoi.Prettify(&res))
}
