// models.article_test.go

package main

import (
	"semaphore_gin_labs/controllers"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := controllers.List()

	// Check that the length of the list of articles returned is the same as the length of the global variable holding the list of articles.
	if len(alist) != len(controllers.List()) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != controllers.ArticleList[i].Content || v.ID != controllers.ArticleList[i].ID || v.Title != controllers.ArticleList[i].Title {

			t.Fail()
			break
		}
	}
}
