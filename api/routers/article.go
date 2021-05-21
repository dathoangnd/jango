package routers

import (
	"os"
	"net/http"
	"strconv"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

type Response struct {
	Title string `json:"title"`
	Id string `json:"id"`
	Slug string `json:"slug"`
	UpdatedAt string `json:"updatedAt"`
	Avatar string `json:"avatar"`
	Categories string `json:"categories"`
	Content string `json:"content"`
	Draft string `json:"draft"`
	Description string `json:"description"`
	UserName string `json:"userName"`
	UserId string `json:"userId"`
	LikeCount int `json:"likeCount"`
	CommentCount int `json:"commentCount"`
	Publish int `json:"publish"`
}

type Comment struct {
	Id string `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	UserName string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	UserId string `json:"userId"`
	Content string `json:"content"`
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	limitQuery := r.URL.Query()["limit"]
	userQuery := r.URL.Query()["user"]
	publishQuery := r.URL.Query()["publish"]
	var limit int
	var user string
	var publish string

	if len(limitQuery) == 0 || limitQuery[0] == "" {
		limit = -1
	} else {
		limit, _ = strconv.Atoi(limitQuery[0])
	}

	if len(userQuery) == 0 || userQuery[0] == "" {
		user = "%"
	} else {
		user = userQuery[0]
	}

	if len(publishQuery) == 0 || publishQuery[0] == "" {
		publish = "%"
	} else {
		publish = publishQuery[0]
	}

	rows, err := db.Table("articles").Select("articles.title, articles.id, articles.updated_at, articles.avatar, articles.categories, articles.content, articles.description, articles.slug, articles.publish, users.name, users.id").Joins("inner join users on articles.user_id = users.id").Where("user_id LIKE ? AND publish LIKE ?", user, publish).Limit(limit).Order("updated_at desc").Rows()
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		articles := []Response{}

		for rows.Next() {
			var title string
			var id string
			var slug string
			var updatedAt string
			var avatar string
			var categories string
			var content string
			var description string
			var userName string
			var userId string
			var likeCount int
			var commentCount int
			var publish int

			rows.Scan(&title, &id, &updatedAt, &avatar, &categories, &content, &description, &slug, &publish, &userName, &userId)
			db.Model(&models.ArticleLike{}).Where("article_id = ?", id).Count(&likeCount)
			db.Model(&models.ArticleComment{}).Where("article_id = ?", id).Count(&commentCount)
		
			articles = append(articles, Response{
				title,
				id,
				slug,
				updatedAt,
				avatar,
				categories,
				content,
				"",
				description,
				userName,
				userId,
				likeCount,
				commentCount,
				publish,
			})
		}

		utils.JSONResponse(w, 200, utils.JSON {
			"articles": articles,
		})
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	id := r.URL.Query()["id"][0]
	limitNewestQuery := r.URL.Query()["newest"]
	var limitNewest int

	if len(limitNewestQuery) == 0 || limitNewestQuery[0] == "" {
		limitNewest = -1
	} else {
		limitNewest, _ = strconv.Atoi(limitNewestQuery[0])
	}

	if db.Table("articles").Where("id = ?", id).First(&models.Article{}).RecordNotFound() {
		utils.JSONResponse(w, 404, utils.JSON {
			"status": "Không tìm thấy bài viết",
		})
		return
	}

	type NewestArticle struct {
		Title string `json:"title"`
		Id string `json:"id"`
		Slug string `json:"slug"`
		Avatar string `json:"avatar"`
		UserName string `json:"userName"`
	}

	newestArticleChan := make(chan NewestArticle, 5)

	go func() {
		rows, err := db.Limit(limitNewest).Table("articles").Select("articles.title, articles.id, articles.slug, articles.avatar, users.name").Joins("inner join users on articles.user_id = users.id").Where("articles.id <> ? AND articles.publish = ?", id, 1).Rows()
		if err != nil {
		} else {
			for rows.Next() {
				var title string
				var id string
				var slug string
				var avatar string
				var userName string

				rows.Scan(&title, &id, &slug, &avatar, &userName)
				newestArticleChan <- NewestArticle{
					title,
					id,
					slug,
					avatar,
					userName,
				}
			}
			close(newestArticleChan)
		}
	}() 

	var article Response
	rows, err := db.Table("articles").Select("articles.title, articles.id, articles.updated_at, articles.avatar, articles.categories, articles.content, articles.draft, articles.description, articles.slug, articles.publish, users.name, users.id").Joins("inner join users on articles.user_id = users.id").Where("articles.id = ?", id).Rows()
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		rows.Next()
		var title string
		var id string
		var slug string
		var updatedAt string
		var avatar string
		var categories string
		var content string
		var draft string
		var description string
		var userName string
		var userId string
		var likeCount int
		var publish int
		rows.Scan(&title, &id, &updatedAt, &avatar, &categories, &content, &draft, &description, &slug, &publish, &userName, &userId)
		db.Model(&models.ArticleLike{}).Where("article_id = ?", id).Count(&likeCount)
		article = Response{
			title,
			id,
			slug,
			updatedAt,
			avatar,
			categories,
			content,
			draft,
			description,
			userName,
			userId,
			likeCount,
			0,
			publish,
		}
		
		newestArticles := []NewestArticle {}
		for article := range newestArticleChan {
			newestArticles = append(newestArticles, article)
    }

		utils.JSONResponse(w, 200, utils.JSON {
			"article": article,
			"newestArticles": newestArticles,
		})
	}
}

func removeArticle(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	id := r.FormValue("id")

	removingArticle := models.Article{}
	if err := db.Where("id = ? AND user_id = ?", id, userId).First(&removingArticle).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	if err := db.Unscoped().Delete(&removingArticle).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã xóa bài viết",
		})
	}
}

func saveDraft(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	user := models.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	description := r.FormValue("description")
	categories := r.FormValue("categories")
	avatar := r.FormValue("avatar")
	slug := r.FormValue("slug")

	article := models.Article{
		Title: title,
		Draft: content,
		Description: description,
		Categories: categories,
		Avatar: avatar,
		UserID: user.ID,
		Slug: slug,
	}

	if _, err := govalidator.ValidateStruct(article); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if r.Method == "POST" {
			if err := db.Create(&article).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Thông tin không hợp lệ",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Đã lưu nháp bài viết",
					"id": article.ID,
				})
			}
		} else if r.Method == "PUT" {
			id := r.FormValue("id")
			updatingArticle := models.Article{}
			if err := db.Where("id = ? AND user_id = ?", id, user.ID).First(&updatingArticle).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Thông tin không hợp lệ",
				})
				return
			}
			updatingArticle.Title = article.Title
			updatingArticle.Draft = article.Draft
			updatingArticle.Description = article.Description
			updatingArticle.Categories = article.Categories
			updatingArticle.Avatar = article.Avatar

			if err := db.Save(&updatingArticle).Error; err != nil {
				utils.JSONResponse(w, 500, utils.JSON {
					"status": "Đã xảy ra sự cố, vui lòng thử lại.",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Đã lưu nháp bài viết",
				})
			}
		}
	}
}

func publish(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	user := models.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	description := r.FormValue("description")
	categories := r.FormValue("categories")
	avatar := r.FormValue("avatar")
	slug := r.FormValue("slug")

	article := models.Article{
		Title: title,
		Content: content,
		Description: description,
		Categories: categories,
		Avatar: avatar,
		UserID: user.ID,
		Slug: slug,
	}

	if _, err := govalidator.ValidateStruct(article); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if r.Method == "POST" {
			if err := db.Create(&article).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Thông tin không hợp lệ",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Bài đăng đang chờ xét duyệt",
					"id": article.ID,
				})
			}
		} else if r.Method == "PUT" {
			id := r.FormValue("id")
			updatingArticle := models.Article{}
			if err := db.Where("id = ? AND user_id = ?", id, user.ID).First(&updatingArticle).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Thông tin không hợp lệ",
				})
				return
			}
			updatingArticle.Title = article.Title
			updatingArticle.Content = article.Content
			updatingArticle.Draft = ""
			updatingArticle.Description = article.Description
			updatingArticle.Categories = article.Categories
			updatingArticle.Avatar = article.Avatar

			if err := db.Save(&updatingArticle).Error; err != nil {
				utils.JSONResponse(w, 500, utils.JSON {
					"status": "Đã xảy ra sự cố, vui lòng thử lại.",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Đã cập nhật bài viết",
				})
			}
		}
	}
}

func getComments(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	articleId := r.URL.Query()["id"][0]
	comments := []Comment{}
	rows, err := db.Table("article_comments").Select("article_comments.id, article_comments.updated_at, article_comments.content, users.name, users.avatar, users.id").Joins("inner join articles on article_comments.article_id = articles.id").Joins("inner join users on article_comments.user_id = users.id").Where("article_comments.article_id = ?", articleId).Order("updated_at desc").Rows()
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		for rows.Next() {
			var id string
			var updatedAt string
			var content string
			var userName string
			var userAvatar string
			var userId string

			rows.Scan(&id, &updatedAt, &content, &userName, &userAvatar, &userId)
			comments = append(comments, Comment{
				id,
				updatedAt,
				userName,
				userAvatar,
				userId,
				content,
			})
		}

		utils.JSONResponse(w, 200, utils.JSON {
			"comments": comments,
		})
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	user := models.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	articleId, _ := strconv.Atoi(r.FormValue("articleId"))
	content := r.FormValue("content")

	comment := models.ArticleComment{
		UserID: user.ID,
		ArticleID: uint(articleId),
		Content: content,
	}

	if _, err := govalidator.ValidateStruct(comment); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if err := db.Create(&comment).Error; err != nil {
			utils.JSONResponse(w, 400, utils.JSON {
				"status": "Thông tin không hợp lệ",
			})
		} else {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã đăng bình luận",
				"comment": utils.JSON {
					"id": comment.ID,
					"updatedAt": comment.UpdatedAt,
					"userName": user.Name,
					"userAvatar": user.Avatar,
					"userId": user.ID,
					"content": comment.Content,
				},
			})
		}
	}
}

func removeComment(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	commentId := r.FormValue("commentId")

	db := models.GetDb()
	var err error
	if userId == os.Getenv("ADMIN_ID") {
		err = db.Unscoped().Where("id = ?", commentId).Delete(&models.ArticleComment{}).Error
	} else {
		err = db.Unscoped().Where("id = ? AND user_id = ?", commentId, userId).Delete(&models.ArticleComment{}).Error
	}
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã xóa bình luận",
		})
	}
}

func like(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	user := models.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	articleId, _ := strconv.Atoi(r.FormValue("articleId"))

	like := models.ArticleLike{
		UserID: user.ID,
		ArticleID: uint(articleId),
	}

	if _, err := govalidator.ValidateStruct(like); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if err := db.Create(&like).Error; err != nil {
			utils.JSONResponse(w, 400, utils.JSON {
				"status": "Thông tin không hợp lệ",
			})
		} else {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã bày tỏ cảm xúc",
			})
		}
	}
}

func unlike(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	articleId := r.FormValue("articleId")

	db := models.GetDb()
	
	err := db.Unscoped().Where("article_id = ? AND user_id = ?", articleId, userId).Delete(&models.ArticleLike{}).Error
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã bỏ bày tỏ cảm xúc",
		})
	}
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" || userId != os.Getenv("ADMIN_ID") {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()

	id := r.FormValue("id")
	status, _ := strconv.Atoi(r.FormValue("status"))
		
	updatingArticle := models.Article{}
	if err := db.Where("id = ?", id).First(&updatingArticle).Error; err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
		return
	}
	
	updatingArticle.Publish = status

	if err := db.Save(&updatingArticle).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		if status == 1 {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã xuất bản bài viết",
			})
		} else {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã thu hồi bài viết",
			})
		}
	}
}