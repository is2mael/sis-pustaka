package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Publisher_Year   int    `json:"publisher_year"`
}