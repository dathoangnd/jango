package routers

import (
	"net/http"
	"jango/models"
	"jango/utils"
)

func getResults(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	searchTerm := r.URL.Query()["s"][0]

	type SearchResult struct {
		Path string `json:"path"`
		Title string `json:"title"`
		Group string `json:"group"`
	}
	articleChan := make(chan SearchResult, 5)
	go func() {
		rows, err := db.Limit(5).Table("articles").Select("id, slug, title").Where("publish = ? AND title LIKE ?", 1, "%" + searchTerm + "%").Order("updated_at desc").Rows()
		if err != nil {
		} else {
			for rows.Next() {
				var id string
				var slug string
				var title string

				rows.Scan(&id, &slug, &title)
				articleChan <- SearchResult{
					slug + "-" + id,
					title,
					"article",
				}
			}
			close(articleChan)
		}
	}()

	documentChan := make(chan SearchResult, 5)
	go func() {
		rows, err := db.Limit(5).Table("documents").Select("id, slug, title").Where("publish = ? AND title LIKE ?", 1, "%" + searchTerm + "%").Order("updated_at desc").Rows()
		if err != nil {
		} else {
			for rows.Next() {
				var id string
				var slug string
				var title string

				rows.Scan(&id, &slug, &title)
				documentChan <- SearchResult{
					slug + "-" + id,
					title,
					"document",
				}
			}
			close(documentChan)
		}
	}()

	discussionChan := make(chan SearchResult, 5)
	go func() {
		rows, err := db.Limit(5).Table("threads").Select("id, content").Where("content LIKE ?", "%" + searchTerm + "%").Order("updated_at desc").Rows()
		if err != nil {
		} else {
			for rows.Next() {
				var id string
				var content string

				rows.Scan(&id, &content)
				discussionChan <- SearchResult{
					"chu-de-" + id,
					content,
					"discussion",
				}
			}
			close(discussionChan)
		}
	}()

	results := []SearchResult {}
	for article := range articleChan {
		results = append(results, article)
	}

	for document := range documentChan {
		results = append(results, document)
	}

	for discussion := range discussionChan {
		results = append(results, discussion)
	}
	
	utils.JSONResponse(w, 200, utils.JSON {
		"results": results,
	})
}