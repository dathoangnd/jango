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
func uploadPdf(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	if err != nil {
		utils.JSONResponse(w, 500)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "application/pdf":
		savePdfFile(w, file, handle)
	default:
		utils.JSONResponse(w, 400)
	}
}

func savePdfFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		utils.JSONResponse(w, 500)
		return
	}

	pdfFolder := fmt.Sprintf("%sdocuments/", UPLOAD_DIR[1:])
	if _, err := os.Stat(pdfFolder); os.IsNotExist(err) {
    os.Mkdir(pdfFolder, 0777)
	}
	pdfLocation := fmt.Sprintf("%s%d_%s", pdfFolder, time.Now().Unix(), handle.Filename)

	err = ioutil.WriteFile(pdfLocation, data, 0666)
	if err != nil {
		utils.JSONResponse(w, 500)
		return
	}
	utils.JSONResponse(w, 200, utils.JSON {
		"location": pdfLocation,
	})
}