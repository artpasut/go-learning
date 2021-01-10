package main

type article struct {
	ID      int    `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

var articleList = []article{
	article{ID: 1, Title: "Golang", Content: "Learn how to dev in Golang."},
	article{ID: 2, Title: "Webserver", Content: "How to make webserver with Go"},
}

func getAllArticles() []article {
	return articleList
}
