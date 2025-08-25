package main

type RawExtractBook struct {
	ASIN       string         `json:"asin"`
	Title      string         `json:"title"`
	Authors    string         `json:"authors"`
	Highlights []RawHighlight `json:"highlights"`
}

type RawHighlight struct {
	Text       string `json:"text"`
	IsNoteOnly bool   `json:"isNoteOnly"`
	Location   struct {
		URL   string `json:"url"`
		Value int    `json:"value"`
	} `json:"location"`
	Note string `json:"note"`
}

type Book struct {
	ID         string      `json:"id"`
	ASIN       string      `json:"asin"`
	Title      string      `json:"title"`
	Authors    string      `json:"authors"`
	UserID     string      `json:"user_id"`
	Highlights []Highlight `json:"highlights"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
}

type Highlight struct {
	ID        string `json:"id"`
	BookID    string `json:"book_id"`
	Text      string `json:"text"`
	Location  int    `json:"location"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_at"`
}
