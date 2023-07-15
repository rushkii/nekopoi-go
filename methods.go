package nekopoi

import (
	"encoding/json"
	"fmt"
)

func GetRecent() (*RecentT, error) {
	data := &RecentT{}

	body, err := Request("/recent")
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	for i := range data.Carousel {
		if len(data.Carousel[i].Slug) > 0 {
			data.Carousel[i].WebUrl = SlugWebURL(data.Carousel[i].Slug)
		} else {
			data.Carousel[i].WebUrl = SlugWebURL(data.Carousel[i].Title)
		}
	}

	for i := range data.Posts {
		for n := range data.Posts[i].Data {
			if len(data.Posts[i].Data[n].Slug) > 0 {
				data.Posts[i].Data[n].WebUrl = SlugWebURL(data.Posts[i].Data[n].Slug)
			} else {
				data.Posts[i].Data[n].WebUrl = SlugWebURL(data.Posts[i].Data[n].Title)
			}
		}
	}

	return data, nil
}

func Search(query string, page uint8) (*SearchT, error) {
	data := &SearchT{}

	body, err := Request(fmt.Sprintf("/search?q=%s&page=%d", query, page))
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	if len(data.Result) > 0 {
		for i := range data.Result {
			if len(data.Result[i].Slug) > 0 {
				data.Result[i].WebUrl = SlugWebURL(data.Result[i].Slug)
			} else {
				data.Result[i].Slug = Slugify(data.Result[i].Title)
				data.Result[i].WebUrl = SlugWebURL(data.Result[i].Title)
			}
		}
	}

	return data, nil
}

type termidT interface {
	uint64 | []uint64
}

func SearchByGenre[T termidT](term_id T) (*SearchT, error) {
	data := &SearchT{}

	var params string

	if terms, ok := any(term_id).([]uint64); ok {
		for _, term := range terms {
			params += fmt.Sprintf("term=%d&", term)
		}
		params = params[0 : len(params)-1]
	} else {
		params = fmt.Sprintf("term=%v", term_id)
	}

	body, err := Request(fmt.Sprintf("/searchByGenre?%s", params))
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	if len(data.Result) > 0 {
		for i := range data.Result {
			if len(data.Result[i].Slug) > 0 {
				data.Result[i].WebUrl = SlugWebURL(data.Result[i].Slug)
			} else {
				data.Result[i].Slug = Slugify(data.Result[i].Title)
				data.Result[i].WebUrl = SlugWebURL(data.Result[i].Title)
			}
		}
	}

	return data, nil
}

func GetDetail(id uint64) (*DetailT, error) {
	data := &DetailT{}

	body, err := Request(fmt.Sprintf("/post?id=%d", id))
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	if data.Slug == "" {
		data.Slug = Slugify(data.Title)
		data.WebUrl = SlugWebURL(data.Title)
	} else {
		data.WebUrl = SlugWebURL(data.Slug)
	}

	return data, nil
}

func GetByIndex(letter string, filter string) (*SearchT, error) {
	data := &SearchT{}

	body, err := Request(fmt.Sprintf("/listall?letter=%s&type=%s&page=1", letter, filter))
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	for i := range data.Result {
		if data.Result[i].Slug == "" {
			data.Result[i].Slug = Slugify(data.Result[i].Title)
			data.Result[i].WebUrl = SlugWebURL(data.Result[i].Title)
		} else {
			data.Result[i].WebUrl = SlugWebURL(data.Result[i].Slug)
		}
	}

	return data, nil
}

func GetSeries(id uint64) (*SeriesT, error) {
	data := &SeriesT{}

	body, err := Request(fmt.Sprintf("/series?id=%d", id))
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	if data.Slug == "" {
		data.Slug = Slugify(data.Title)
		data.WebUrl = SlugWebURL(data.Title)
	} else {
		data.WebUrl = SlugWebURL(data.Slug)
	}

	for i := range data.Episodes {
		if data.Episodes[i].Slug == "" {
			data.Episodes[i].Slug = Slugify(data.Episodes[i].Title)
			data.Episodes[i].WebUrl = SlugWebURL(data.Episodes[i].Title)
		} else {
			data.Episodes[i].WebUrl = SlugWebURL(data.Episodes[i].Slug)
		}
	}

	return data, nil
}

func GetComments(slug string) (*DiscussionT, error) {
	data := &DiscussionT{}

	body, err := Request(fmt.Sprintf("/service/disqus/post?slug=%s", slug))
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	return data, nil
}

func GetComingSoon() (*ComingSoonT, error) {
	data := &ComingSoonT{}

	body, err := Request("/comingsoon")
	if err != nil {
		return nil, err
	}

	defer body.Close()

	if json.NewDecoder(body).Decode(&data) != nil {
		return nil, err
	}

	return data, nil
}
