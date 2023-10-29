package controllers

import (
	"semaphore_gin_labs/models"

	"errors"
)

var ArticleList = []models.Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func List() []models.Article {
	return ArticleList
}

// GetArticleByID retrieves an article by its ID from the articleList.
// It takes an integer 'id' as a parameter and returns a pointer to an 'article' and an error.
func GetArticleByID(id int) (*models.Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			// If a matching article is found, return a pointer to the article and no error.
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// DeleteArticleByID deletes an article by its ID from the articleList.
// It takes an integer 'id' as a parameter and returns an error.
func DeleteArticleByID(id int) error {
	for i, a := range ArticleList {
		if a.ID == id {
			// If a matching article is found, append the articleList up to (but excluding) that article,
			// then append the rest of the articleList, starting after the article.
			ArticleList = append(ArticleList[:i], ArticleList[i+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}

// UpdateArticleByID updates an article by its ID from the articleList.
// It takes an integer 'id', a pointer to an 'article', and returns an error.
func UpdateArticleByID(id int, newArticle *models.Article) error {
	for i, a := range ArticleList {
		if a.ID == id {
			// If a matching article is found, update the articleList at that index.
			ArticleList[i].Title = newArticle.Title
			ArticleList[i].Content = newArticle.Content
			return nil
		}
	}
	return errors.New("Article not found")
}

// CreateArticle creates a new article and adds it to the articleList.
// It takes a pointer to an 'article' and returns the new article's ID.
func CreateArticle(newArticle *models.Article) int {
	// Set the ID of a new article to one more than the number of articles in the articleList.
	// This will make the ID of the new article unique.
	newArticle.ID = len(ArticleList) + 1
	// Add the new article to the articleList.
	ArticleList = append(ArticleList, *newArticle)
	// Return the ID of the new article.
	return newArticle.ID
}
