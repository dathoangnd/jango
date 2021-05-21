package routers

import (
	"net/http"
	"strconv"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

type Voca struct {
	ID uint `json:"id"`
	Text string `json:"text"`
	Read string `json:"read"`
	Meaning string `json:"meaning"`
	ExampleText string `json:"exampleText"`
	ExampleRead string `json:"exampleRead"`
	ExampleMeaning string `json:"exampleMeaning"`
}

func getVocas(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	courseQuery := r.URL.Query()["course"]
	lessonQuery := r.URL.Query()["lesson"]

	var course int
	var lesson int

	if len(courseQuery) == 0 || courseQuery[0] == "" {
		course = 0
	} else {
		course, _ = strconv.Atoi(courseQuery[0])
	}

	if len(lessonQuery) == 0 || lessonQuery[0] == "" {
		lesson = 0
	} else {
		lesson, _ = strconv.Atoi(lessonQuery[0])
	}

	var vocas []Voca
	var err error

	if course == 0 {
		err = db.Table("vocas").Scan(&vocas).Error
	} else if lesson == 0 {
		err = db.Table("vocas").Where("lesson BETWEEN ? AND ?", (course - 1) * 10 + 1, course * 10).Scan(&vocas).Error
	} else {
		err = db.Table("vocas").Where("lesson = ?", (course - 1) * 10 + lesson).Scan(&vocas).Error
	}
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"vocas": vocas,
		})
	}
}

func getLikedVocas(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	db := models.GetDb()
	courseQuery := r.URL.Query()["course"]

	var course int

	if len(courseQuery) == 0 || courseQuery[0] == "" {
		course = 0
	} else {
		course, _ = strconv.Atoi(courseQuery[0])
	}

	var vocas []Voca
	var err error

	if course == 0 {
		err = db.Table("vocas").Joins("inner join voca_likes on vocas.id = voca_likes.voca_order").Where("voca_likes.user_id = ?", userId).Scan(&vocas).Error
	} else {
		err = db.Table("vocas").Joins("inner join voca_likes on vocas.id = voca_likes.voca_order").Where("voca_likes.user_id = ? AND vocas.lesson BETWEEN ? AND ?", userId, (course - 1) * 10 + 1, course * 10).Scan(&vocas).Error
	}
	
	if err != nil {
		utils.JSONResponse(w, 500, utils.JSON {
			"status": "Đã xảy ra sự cố, vui lòng thử lại.",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"vocas": vocas,
		})
	}
}

func likeVoca(w http.ResponseWriter, r *http.Request) {
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

	vocaOrder, _ := strconv.Atoi(r.FormValue("vocaOrder"))

	like := models.VocaLike{
		UserID: user.ID,
		VocaOrder: uint(vocaOrder),
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
				"status": "Đã thêm vào từ yêu thích",
			})
		}
	}
}

func unlikeVoca(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	vocaOrder := r.FormValue("vocaOrder")

	db := models.GetDb()
	
	err := db.Unscoped().Where("voca_order = ? AND user_id = ?", vocaOrder, userId).Delete(&models.VocaLike{}).Error
	
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