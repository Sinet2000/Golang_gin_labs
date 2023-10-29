// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
	return articleList
}

// getArticleByID retrieves an article by its ID from the articleList.
// It takes an integer 'id' as a parameter and returns a pointer to an 'article' and an error.
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			// If a matching article is found, return a pointer to the article and no error.
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
