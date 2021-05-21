package routers

import (
	"os"
	"net/http"
	"strconv"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

type DocumentResponse struct {
	ID uint `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	Title string `json:"title"`
	Avatar string `json:"avatar"`
	UserId uint `json:"userId"`
	UserName string `json:"userName"`
	Categories string `json:"categories"`
	Content string `json:"content"`
	Publish int `json:"publish"`
	Slug string `json:"slug"`
	File string `json:"file"`
	Author string `json:"author"`
	Source string `json:"source"`
	Download uint `json:"download"`
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.Table("documents").Select("documents.id, documents.updated_at, documents.title, documents.avatar, documents.categories, documents.content, documents.publish, documents.slug, documents.file, documents.author, documents.source, documents.download, users.id, users.name").Joins("inner join users on documents.user_id = users.id").Where("user_id LIKE ? AND publish LIKE ?", user, publish).Limit(limit).Order("updated_at desc").Rows()
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		documents := []DocumentResponse{}

		for rows.Next() {
			var id uint
			var updatedAt string
			var title string
			var avatar string
			var categories string
			var content string
			var publish int
			var slug string
			var file string
			var author string
			var source string
			var download uint
			var userId uint
			var userName string

			rows.Scan(&id, &updatedAt, &title, &avatar, &categories, &content, &publish, &slug, &file, &author, &source, &download, &userId, &userName)
			documents = append(documents, DocumentResponse{
				id,
				updatedAt,
				title,
				avatar,
				userId,
				userName,
				categories,
				content,
				publish,
				slug,
				file,
				author,
				source,
				download,
			})
		}

		utils.JSONResponse(w, 200, utils.JSON {
			"documents": documents,
		})
	}
}

func getDocument(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	id := r.URL.Query()["id"][0]
	limitNewestQuery := r.URL.Query()["newest"]
	var limitNewest int

	if len(limitNewestQuery) == 0 || limitNewestQuery[0] == "" {
		limitNewest = -1
	} else {
		limitNewest, _ = strconv.Atoi(limitNewestQuery[0])
	}

	if db.Table("documents").Where("id = ?", id).First(&models.Document{}).RecordNotFound() {
		utils.JSONResponse(w, 404, utils.JSON {
			"status": "Không tìm thấy tài liệu",
		})
		return
	}

	type NewestDocument struct {
		Title string `json:"title"`
		Id uint `json:"id"`
		Slug string `json:"slug"`
		Avatar string `json:"avatar"`
		UserName string `json:"userName"`
	}

	newestDocumentChan := make(chan NewestDocument, 5)

	go func() {
		rows, err := db.Limit(limitNewest).Table("documents").Select("documents.title, documents.id, documents.slug, documents.avatar, users.name").Joins("inner join users on documents.user_id = users.id").Where("documents.id <> ? AND documents.publish = ?", id, 1).Rows()
		if err != nil {
		} else {
			for rows.Next() {
				var title string
				var id uint
				var slug string
				var avatar string
				var userName string

				rows.Scan(&title, &id, &slug, &avatar, &userName)
				newestDocumentChan <- NewestDocument{
					title,
					id,
					slug,
					avatar,
					userName,
				}
			}
			close(newestDocumentChan)
		}
	}()

	var document DocumentResponse
	rows, err := db.Table("documents").Select("documents.id, documents.updated_at, documents.title, documents.avatar, documents.categories, documents.content, documents.publish, documents.slug, documents.file, documents.author, documents.source, documents.download, users.id, users.name").Joins("inner join users on documents.user_id = users.id").Where("documents.id = ?", id).Rows()
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		go func() {	
			var updatingDocument models.Document
			if err := db.Where("id = ?", id).First(&updatingDocument).Error; err != nil {
			} else {
				updatingDocument.Download += 1
				db.Save(&updatingDocument)
			}
		}()

		rows.Next()
		var id uint
		var updatedAt string
		var title string
		var avatar string
		var categories string
		var content string
		var publish int
		var slug string
		var file string
		var author string
		var source string
		var download uint
		var userId uint
		var userName string

		rows.Scan(&id, &updatedAt, &title, &avatar, &categories, &content, &publish, &slug, &file, &author, &source, &download, &userId, &userName)
		document = DocumentResponse{
			id,
			updatedAt,
			title,
			avatar,
			userId,
			userName,
			categories,
			content,
			publish,
			slug,
			file,
			author,
			source,
			download,
		}
		
		newestDocuments := []NewestDocument {}
		for document := range newestDocumentChan {
			newestDocuments = append(newestDocuments, document)
    }

		utils.JSONResponse(w, 200, utils.JSON {
			"document": document,
			"newestDocuments": newestDocuments,
		})
	}
}

func removeDocument(w http.ResponseWriter, r *http.Request) {
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

	removingDocument := models.Document{}
	if err := db.Where("id = ? AND user_id = ?", id, userId).First(&removingDocument).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	if err := db.Unscoped().Delete(&removingDocument).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã xóa tài liệu",
		})
	}
}

func publishDocument(w http.ResponseWriter, r *http.Request) {
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
	avatar := r.FormValue("avatar")
	categories := r.FormValue("categories")
	content := r.FormValue("content")
	slug := r.FormValue("slug")
	file := r.FormValue("file")
	author := r.FormValue("author")
	source := r.FormValue("source")

	document := models.Document{
		Title: title,
		Avatar: avatar,
		UserID: user.ID,
		Categories: categories,
		Content: content,
		Publish: 0,
		Slug: slug,
		File: file,
		Author: author,
		Source: source,
		Download: 0,
	}

	if _, err := govalidator.ValidateStruct(document); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if r.Method == "POST" {
			if err := db.Create(&document).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Thông tin không hợp lệ",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Tài liệu đang chờ xét duyệt",
					"id": document.ID,
				})
			}
		} else if r.Method == "PUT" {
			id := r.FormValue("id")
			updatingDocument := models.Document{}
			if err := db.Where("id = ? AND user_id = ?", id, user.ID).First(&updatingDocument).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Thông tin không hợp lệ",
				})
				return
			}
			updatingDocument.Title = document.Title
			updatingDocument.Avatar = document.Avatar
			updatingDocument.Categories = document.Categories
			updatingDocument.Content = document.Content
			updatingDocument.File = document.File
			updatingDocument.Author = document.Author
			updatingDocument.Source = document.Source

			if err := db.Save(&updatingDocument).Error; err != nil {
				utils.JSONResponse(w, 500, utils.JSON {
					"status": "Đã xảy ra sự cố, vui lòng thử lại.",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Đã cập nhật tài liệu",
				})
			}
		}
	}
}

func changeDocumentStatus(w http.ResponseWriter, r *http.Request) {
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
		
	updatingDocument := models.Document{}
	if err := db.Where("id = ?", id).First(&updatingDocument).Error; err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
		return
	}
	
	updatingDocument.Publish = status

	if err := db.Save(&updatingDocument).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		if status == 1 {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã xuất bản tài liệu",
			})
		} else {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã thu hồi tài liệu",
			})
		}
	}
}