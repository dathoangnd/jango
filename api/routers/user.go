package routers

import (
	"net/http"
	"strconv"
	"jango/models"
	"jango/utils"
	"github.com/asaskevich/govalidator"
)

func register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	name := r.FormValue("name")

	newUser := models.User{
		Email: email,
		Password: password,
		Name: name,
	}
	
	if _, err := govalidator.ValidateStruct(newUser); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if hashPassword, err := utils.HashPassword(password); err != nil {
			utils.JSONResponse(w, 500, utils.JSON {
				"status": "Đã xảy ra sự cố, vui lòng thử lại.",
			})
		} else {
			newUser.Password = hashPassword
			db := models.GetDb()
			if err := db.Create(&newUser).Error; err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Email đã được đăng ký",
				})
			} else {
				if token, err := utils.SignToken(newUser); err != nil {
					utils.JSONResponse(w, 500, utils.JSON {
						"status": "Đã xảy ra sự cố, vui lòng thử lại.",
					})
				} else {
					utils.JSONResponse(w, 200, utils.JSON {
						"status": "Đăng ký tài khoản thành công",
						"token": token,
						"likedArticles": "",
						"likedVocas": "",
						"likedGrammars": "",
						"likedKanji": "",
					})
				}
			}
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	const PLACEHOLDER = "Jango"
	user := models.User{
		Email: email,
		Password: password,
		Name: PLACEHOLDER, // Just for validator
	}
	
	if _, err := govalidator.ValidateStruct(user); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		db := models.GetDb()
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			utils.JSONResponse(w, 400, utils.JSON {
				"status": "Tài khoản hoặc mật khẩu không đúng",
			})
		} else {
			if err := utils.ValidatePassword(user.Password, password); err != nil {
				utils.JSONResponse(w, 400, utils.JSON {
					"status": "Tài khoản hoặc mật khẩu không đúng",
				})
			} else {
				if token, err := utils.SignToken(user); err != nil {
					utils.JSONResponse(w, 500, utils.JSON {
						"status": "Đã xảy ra sự cố, vui lòng thử lại.",
					})
				} else {
					vocaChan := make(chan string)
					go func() {
						likedVocas := ""
						rows, _ := db.Table("voca_likes").Select("voca_order").Where("user_id = ?", user.ID).Rows()
						for rows.Next() {
							var likedVoca string
							rows.Scan(&likedVoca)
							likedVocas += (likedVoca + " ")
						}
						vocaChan <- likedVocas
					}()

					grammarChan := make(chan string)
					go func() {
						likedGrammars := ""
						rows, _ := db.Table("grammar_likes").Select("grammar_order").Where("user_id = ?", user.ID).Rows()
						for rows.Next() {
							var likedGrammar string
							rows.Scan(&likedGrammar)
							likedGrammars += (likedGrammar + " ")
						}
						grammarChan <- likedGrammars
					}()

					kanjiChan := make(chan string)
					go func() {
						likedKanjiList := ""
						rows, _ := db.Table("kanji_likes").Select("kanji_order").Where("user_id = ?", user.ID).Rows()
						for rows.Next() {
							var likedKanji string
							rows.Scan(&likedKanji)
							likedKanjiList += (likedKanji + " ")
						}
						kanjiChan <- likedKanjiList
					}()
	
					var likedArticles string
					rows, _ := db.Table("article_likes").Select("article_id").Where("user_id = ?", user.ID).Rows()
					for rows.Next() {
						var likedArticle string
						rows.Scan(&likedArticle)
						likedArticles += (likedArticle + " ")
					}

					likedVocas := <- vocaChan
					likedGrammars := <- grammarChan
					likedKanji := <- kanjiChan

					utils.JSONResponse(w, 200, utils.JSON {
						"status": "Đăng nhập thành công",
						"token": token,
						"likedArticles": likedArticles,
						"likedVocas": likedVocas,
						"likedGrammars": likedGrammars,
						"likedKanji": likedKanji,
					})
				}
			}
		}
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	db := models.GetDb()
	id := r.URL.Query()["id"][0]

	type UserInfo struct {
		Email string `json:"email"`
		Name string `json:"name"`
		Avatar string `json:"avatar"`
		Birthday int64 `json:"birthday"`
		Gender int `json:"gender"`
		City string `json:"city"`
		FacebookId string `json:"facebookId"`
		About string `json:"about"`
	}

	var userInfo UserInfo

	if db.Table("users").Where("id = ?", id).Scan(&userInfo).RecordNotFound() {
		utils.JSONResponse(w, 404, utils.JSON {
			"status": "Không tìm thấy bài viết",
		})
	} else {
		utils.JSONResponse(w, 200, utils.JSON {
			"user": userInfo,
		})
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
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

	name := r.FormValue("name")
	avatar := r.FormValue("avatar")
	birthday, _ := strconv.Atoi(r.FormValue("birthday"))
	gender, _ := strconv.Atoi(r.FormValue("gender"))
	city := r.FormValue("city")
	facebookId := r.FormValue("facebookId")
	about := r.FormValue("about")

	user.Name = name
	user.Avatar = avatar
	user.Birthday = int64(birthday)
	user.Gender = gender
	user.City = city
	user.FacebookId = facebookId
	user.About = about

	if _, err := govalidator.ValidateStruct(user); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if err := db.Save(&user).Error; err != nil {
			utils.JSONResponse(w, 500, utils.JSON {
				"status": "Đã xảy ra sự cố, vui lòng thử lại.",
			})
		} else {
			token, _ := utils.SignToken(user)
			utils.JSONResponse(w, 200, utils.JSON {
				"status": "Đã cập nhật thông tin",
				"token": token,
			})
		}
	}
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	userId := utils.ValidateToken(reqToken)
	if userId == "" {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}
	oldPassword := r.FormValue("oldPassword")
	db := models.GetDb()
	user := models.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Không có quyền truy cập",
		})
		return
	}

	if err := utils.ValidatePassword(user.Password, oldPassword); err != nil {
		utils.JSONResponse(w, 401, utils.JSON {
			"status": "Mật khẩu cũ không đúng",
		})
		return
	}

	password := r.FormValue("password")

	user.Password = password

	if _, err := govalidator.ValidateStruct(user); err != nil {
		utils.JSONResponse(w, 400, utils.JSON {
			"status": "Thông tin không hợp lệ",
		})
	} else {
		if hashPassword, err := utils.HashPassword(password); err != nil {
			utils.JSONResponse(w, 500, utils.JSON {
				"status": "Đã xảy ra sự cố, vui lòng thử lại.",
			})
		} else {
			user.Password = hashPassword
			if err := db.Save(&user).Error; err != nil {
				utils.JSONResponse(w, 500, utils.JSON {
					"status": "Đã xảy ra sự cố, vui lòng thử lại.",
				})
			} else {
				utils.JSONResponse(w, 200, utils.JSON {
					"status": "Đã thay đổi mật khẩu",
				})
			}
		}
	}
}