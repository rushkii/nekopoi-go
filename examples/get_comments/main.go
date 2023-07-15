package main

import (
	"fmt"
	"log"

	"github.com/rushkii/nekopoi-go"
)

func main() {
	// slug: string (e.g: "araiya-san-ore-to-aitsu-ga-onnayu-de-episode-1-subtitle-indonesia")
	res, err := nekopoi.GetComments("araiya-san-ore-to-aitsu-ga-onnayu-de-episode-1-subtitle-indonesia")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(nekopoi.Prettify(&res))
}
