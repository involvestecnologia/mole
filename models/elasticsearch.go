package models

type SearchResults struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Querys []Querys `json:"hits"`
}

type Querys struct {
	Source Source `json:"_source"`
}

type Source struct {
	Timestamp string `json:"timestamp"`
}
