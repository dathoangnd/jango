package routers

import (
	"os"
	"net/http"
	"strconv"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

type ThreadResponse struct {
	Id string `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	Categories string `json:"categories"`
	Content string `json:"content"`
	UserName string `json:"userName"`
	UserId string `json:"userId"`
	UserAvatar string `json:"userAvatar"`
	Vote int `json:"vote"`
	ReplyCount int `json:"replyCount"`
}

type ThreadReply struct {
	Id string `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	UserName string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	UserId string `json:"userId"`
	Content string `json:"content"`
	Vote int `json:"vote"`
}

func getThreads(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	limitQuery := r.URL.Query()["limit"]
	userQuery := r.URL.Query()["user"]
	var limit int
	var user string

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

	rows, err := db.Table("threads").Select("threads.id, threads.created_at, threads.categories, threads.content, threads.vote, users.name, users.id, users.avatar").Joins("inner join users on threads.user_id = users.id").Where("user_id LIKE ?", user).Limit(limit).Order("created_at desc").Rows()
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		threads := []ThreadResponse{}

		for rows.Next() {
			var id string
			var updatedAt string
			var categories string
			var content string
			var userName string
			var userId string
			var userAvatar string
			var vote int
			var replyCount int

			rows.Scan(&id, &updatedAt, &categories, &content, &vote, &userName, &userId, &userAvatar)
			db.Model(&models.ThreadReply{}).Where("thread_id = ?", id).Count(&replyCount)
		
			threads = append(threads, ThreadResponse{
				id,
				updatedAt,
				categories,
				content,
				userName,
				userId,
				userAvatar,
				vote,
				replyCount,
			})
		}

		utils.JSONResponse(w, 200, utils.JSON {
			"threads": threads,
		})
	}
}

func getThread(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	id := r.URL.Query()["id"][0]
	limitNewestQuery := r.URL.Query()["newest"]
	var limitNewest int

	if len(limitNewestQuery) == 0 || limitNewestQuery[0] == "" {
		limitNewest = -1
	} else {
		limitNewest, _ = strconv.Atoi(limitNewestQuery[0])
	}

	if db.Table("threads").Where("id = ?", id).First(&models.Thread{}).RecordNotFound() {
		utils.JSONResponse(w, 404, utils.JSON {
			"status": "Không tìm thấy chủ đề",
		})
		return
	}

	type NewestThread struct {
		Id string `json:"id"`
		Content string `json:"content"`
		UserName string `json:"userName"`
	}

	newestThreadChan := make(chan NewestThread, 5)

	go func() {
		rows, err := db.Limit(limitNewest).Table("threads").Select("threads.id, threads.content, users.name").Joins("inner join users on threads.user_id = users.id").Where("threads.id <> ?", id).Rows()
		if err != nil {
		} else {
			for rows.Next() {
				var id string
				var content string
				var userName string

				rows.Scan(&id, &content, &userName)
				newestThreadChan <- NewestThread{
					id,
					content,
					userName,
				}
			}
			close(newestThreadChan)
		}
	}() 

	var thread ThreadResponse
	rows, err := db.Table("threads").Select("threads.created_at, threads.categories, threads.content, threads.vote, users.name, users.id, users.avatar").Joins("inner join users on threads.user_id = users.id").Where("threads.id = ?", id).Rows()
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		rows.Next()
		var updatedAt string
		var categories string
		var content string
		var userName string
		var userId string
		var userAvatar string
		var vote int
		var replyCount int
		rows.Scan(&updatedAt, &categories, &content, &vote, &userName, &userId, &userAvatar)
		db.Model(&models.ThreadReply{}).Where("thread_id = ?", id).Count(&replyCount)
		thread = ThreadResponse{
			id,
			updatedAt,
			categories,
			content,
			userName,
			userId,
			userAvatar,
			vote,
			replyCount,
		}
		
		newestThreads := []NewestThread {}
		for thread := range newestThreadChan {
			newestThreads = append(newestThreads, thread)
    }

		utils.JSONResponse(w, 200, utils.JSON {
			"thread": thread,
			"newestThreads": newestThreads,
		})
	}
}

func removeThread(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	id := r.FormValue("threadId")
	var err error
	if userId == os.Getenv("ADMIN_ID") {
		err = db.Unscoped().Where("id = ?", id).Delete(&models.Thread{}).Error
	} else {
		err = db.Unscoped().Where("id = ? AND user_id = ?", id, userId).Delete(&models.Thread{}).Error
	}

	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã xóa chủ đề",
		})
	}
}

func publishThread(w http.ResponseWriter, r *http.Request) {
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

	content := r.FormValue("content")
	categories := r.FormValue("categories")

	thread := models.Thread{
		UserID: user.ID,
		Categories: categories,
		Content: content,
	}

	if _, err := govalidator.ValidateStruct(thread); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {		
		if err := db.Create(&thread).Error; err != nil {
			utils.JSONResponse(w, 400, utils.JSON {
				"status": "Thông tin không hợp lệ",
			})
		} else {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã đăng chủ đề",
				"thread": utils.JSON {
					"id": thread.ID,
					"updatedAt": thread.UpdatedAt,
					"categories": thread.Categories,
					"content": thread.Content,
					"userName": user.Name,
					"userId": user.ID,
					"userAvatar": user.Avatar,
					"vote": 0,
					"replyCount": 0,
				},
			})
		}
	}
}

func getReplies(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	threadId := r.URL.Query()["id"][0]
	replies := []ThreadReply{}
	rows, err := db.Table("thread_replies").Select("thread_replies.id, thread_replies.created_at, thread_replies.content, thread_replies.vote, users.name, users.avatar, users.id").Joins("inner join threads on thread_replies.thread_id = threads.id").Joins("inner join users on thread_replies.user_id = users.id").Where("thread_replies.thread_id = ?", threadId).Order("created_at desc").Rows()
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
			var vote int

			rows.Scan(&id, &updatedAt, &content, &vote, &userName, &userAvatar, &userId)
			replies = append(replies, ThreadReply{
				id,
				updatedAt,
				userName,
				userAvatar,
				userId,
				content,
				vote,
			})
		}

		utils.JSONResponse(w, 200, utils.JSON {
			"replies": replies,
		})
	}
}

func reply(w http.ResponseWriter, r *http.Request) {
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

	threadId, _ := strconv.Atoi(r.FormValue("threadId"))
	content := r.FormValue("content")

	reply := models.ThreadReply{
		UserID: user.ID,
		ThreadId: uint(threadId),
		Content: content,
	}

	if _, err := govalidator.ValidateStruct(reply); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if err := db.Create(&reply).Error; err != nil {
			utils.JSONResponse(w, 400, utils.JSON {
				"status": "Thông tin không hợp lệ",
			})
		} else {
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã đăng trả lời",
				"reply": utils.JSON {
					"id": reply.ID,
					"updatedAt": reply.UpdatedAt,
					"userName": user.Name,
					"userAvatar": user.Avatar,
					"userId": user.ID,
					"content": reply.Content,
					"vote": 0,
				},
			})
		}
	}
}

func removeReply(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	replyId := r.FormValue("replyId")

	db := models.GetDb()
	var err error
	if userId == os.Getenv("ADMIN_ID") {
		err = db.Unscoped().Where("id = ?", replyId).Delete(&models.ThreadReply{}).Error
	} else {
		err = db.Unscoped().Where("id = ? AND user_id = ?", replyId, userId).Delete(&models.ThreadReply{}).Error
	}
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã xóa trả lời",
		})
	}
}

func upvoteThread(w http.ResponseWriter, r *http.Request) {
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

	threadId, _ := strconv.Atoi(r.FormValue("threadId"))

	updatingThread := models.Thread{}

	if err := db.Where("id = ?", threadId).First(&updatingThread).Error; err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
		return
	}

	updatingThread.Vote++

	if err := db.Save(&updatingThread).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã bình chọn",
		})
	}
}

func downvoteThread(w http.ResponseWriter, r *http.Request) {
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

	threadId, _ := strconv.Atoi(r.FormValue("threadId"))

	updatingThread := models.Thread{}

	if err := db.Where("id = ?", threadId).First(&updatingThread).Error; err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
		return
	}

	updatingThread.Vote--

	if err := db.Save(&updatingThread).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã bình chọn",
		})
	}
}

func upvoteReply(w http.ResponseWriter, r *http.Request) {
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

	replyId, _ := strconv.Atoi(r.FormValue("replyId"))

	updatingReply := models.ThreadReply{}

	if err := db.Where("id = ?", replyId).First(&updatingReply).Error; err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
		return
	}

	updatingReply.Vote++

	if err := db.Save(&updatingReply).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã bình chọn",
		})
	}
}

func downvoteReply(w http.ResponseWriter, r *http.Request) {
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

	replyId, _ := strconv.Atoi(r.FormValue("replyId"))

	updatingReply := models.ThreadReply{}

	if err := db.Where("id = ?", replyId).First(&updatingReply).Error; err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
		return
	}

	updatingReply.Vote--

	if err := db.Save(&updatingReply).Error; err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã bình chọn",
		})
	}
}