package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Publisher_Year   string    `json:"publisher_year"`
}

// file ini nantinya akan kita gunakan untuk mendefinisikan struktur data buku yang akan kita simpan di dalam database.