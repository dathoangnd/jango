package routers

import (
	"net/http"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

func likeKanji(w http.ResponseWriter, r *http.Request) {
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

	kanjiOrder := r.FormValue("kanjiOrder")

	like := models.KanjiLike{
		UserID: user.ID,
		KanjiOrder: kanjiOrder,
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
				"status": "Đã thêm vào chữ Hán yêu thích",
			})
		}
	}
}

func unlikeKanji(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	kanjiOrder := r.FormValue("kanjiOrder")

	db := models.GetDb()
	
	err := db.Unscoped().Where("kanji_order = ? AND user_id = ?", kanjiOrder, userId).Delete(&models.KanjiLike{}).Error
	
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