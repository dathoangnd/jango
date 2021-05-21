package routers

import (
	"net/http"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

func likeGrammar(w http.ResponseWriter, r *http.Request) {
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

	grammarOrder := r.FormValue("grammarOrder")

	like := models.GrammarLike{
		UserID: user.ID,
		GrammarOrder: grammarOrder,
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
				"status": "Đã thêm vào ngữ pháp yêu thích",
			})
		}
	}
}

func unlikeGrammar(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	grammarOrder := r.FormValue("grammarOrder")

	db := models.GetDb()
	
	err := db.Unscoped().Where("grammar_order = ? AND user_id = ?", grammarOrder, userId).Delete(&models.GrammarLike{}).Error
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"status": "Đã bỏ yêu thích",
		})
	}
}