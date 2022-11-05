package models

type defaultResponseFields struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
}

type CharaterResponse struct {
	Docs []Character `json:"docs"`
	defaultResponseFields
}

type BookResponse struct {
	Docs []Book `json:"docs"`
	defaultResponseFields
}

type ChapterResponse struct {
	Docs []Chapter `json:"docs"`
	defaultResponseFields
}

type MovieResponse struct {
	Docs []Movie `json:"docs"`
	defaultResponseFields
}

type QuoteResponse struct {
	Docs []Quote `json:"docs"`
	defaultResponseFields
}
