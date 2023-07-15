package nekopoi

type PreviewT struct {
	ID          uint64 `json:"id"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	WebUrl      string `json:"web_url"`
}

type recentdataT struct {
	Category string     `json:"category"`
	Data     []PreviewT `json:"data"`
}

type RecentT struct {
	Carousel []PreviewT    `json:"carousel"`
	Posts    []recentdataT `json:"posts"`
}

type SearchT struct {
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Result     []PreviewT `json:"result"`
}

type linkT struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type downloadT struct {
	Type  string  `json:"type"`
	Links []linkT `json:"links"`
}

type DetailT struct {
	ID       uint64      `json:"id"`
	Date     string      `json:"date"`
	Title    string      `json:"title"`
	Image    string      `json:"image"`
	Content  string      `json:"content"`
	Series   string      `json:"series"`
	Note     string      `json:"note"`
	Slug     string      `json:"slug"`
	WebUrl   string      `json:"web_url"`
	Stream   []linkT     `json:"stream"`
	Download []downloadT `json:"download"`
}

type genreT struct {
	TermID uint64 `json:"term_id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
}

type infometaT struct {
	Aliases  string   `json:"aliases"`
	Episode  string   `json:"episode"`
	Status   string   `json:"status"`
	Durasi   string   `json:"durasi"`
	Skor     string   `json:"skor"`
	Produser string   `json:"produser"`
	Tayang   string   `json:"tayang"`
	Genre    []genreT `json:"genre"`
}

type SeriesT struct {
	ID       uint64     `json:"id"`
	Date     string     `json:"date"`
	Title    string     `json:"title"`
	Image    string     `json:"image"`
	Slug     string     `json:"slug"`
	WebUrl   string     `json:"web_url"`
	Meta     infometaT  `json:"info_meta"`
	Episodes []PreviewT `json:"episode"`
}

type authorcommentT struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type likecommentT struct {
	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`
}

type messagecommentT struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

type commentT struct {
	ID        string          `json:"id"`
	Author    authorcommentT  `json:"author"`
	Like      likecommentT    `json:"like"`
	Message   messagecommentT `json:"message"`
	CreatedAt string          `json:"created_at"`
}

type DiscussionT struct {
	Thread      string     `json:"thread"`
	Hasnext     bool       `json:"has_next"`
	HasPrevious bool       `json:"has_previous"`
	Total       uint64     `json:"total"`
	Result      []commentT `json:"result"`
}

type ComingSoonT struct {
	Result string `json:"result"`
}
