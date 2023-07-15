package nekopoi

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gosimple/slug"
)

var nekopoiWebURL *url.URL = &url.URL{Scheme: "https", Host: "nekopoi.care"}

func Slugify(text string) string {
	return slug.Make(text)
}

func SlugWebURL(text string) string {
	return nekopoiWebURL.JoinPath(Slugify(text)).String()
}

func Prettify(vStruct any) string {
	v, err := json.MarshalIndent(&vStruct, "", "   ")
	if err != nil {
		log.Println(err)
	}

	return string(v)
}
