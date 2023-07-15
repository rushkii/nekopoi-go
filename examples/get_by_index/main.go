package main

import (
	"fmt"
	"log"

	"github.com/rushkii/nekopoi-go"
	"github.com/rushkii/nekopoi-go/filters"
)

func main() {
	// letter: string (e.g: "0-9" or "A" to "Z")
	// filter: string (e.g: "3d_hentai" or use defined value with filters.TYPE_3D_HENTAI)
	res, err := nekopoi.GetByIndex("0-9", filters.TYPE_JAV_COSPLAY)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(nekopoi.Prettify(&res))
}
