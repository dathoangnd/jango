package routers

import (
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"jango/utils"
)

// UploadFile uploads a file to the server
func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	if err != nil {
		utils.JSONResponse(w, 500)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg", "image/png", "	image/bmp", "image/gif":
		saveFile(w, file, handle)
	default:
		utils.JSONResponse(w, 400)
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		utils.JSONResponse(w, 500)
		return
	}

	imageFolder := fmt.Sprintf("%simages/", UPLOAD_DIR[1:])
	if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
    os.Mkdir(imageFolder, 0777)
	}
	imageLocation := fmt.Sprintf("%s%d_%s", imageFolder, time.Now().Unix(), handle.Filename)

	err = ioutil.WriteFile(imageLocation, data, 0666)
	if err != nil {
		utils.JSONResponse(w, 500)
		return
	}
	utils.JSONResponse(w, 200, utils.JSON {
		"location": imageLocation,
	})
}