package hw14

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	{Id: "1", Title: "Title1", Desc: "Description1", Content: "Content1"},
	{Id: "2", Title: "Title2", Desc: "Description2", Content: "Content2"},
	{Id: "3", Title: "Title3", Desc: "Description3", Content: "Content3"},
}
