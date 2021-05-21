package utils

import (
	"net/http"
	"encoding/json"
)

type JSON map[string]interface{}

func JSONResponse(w http.ResponseWriter, status int, jsonMap ...JSON) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if len(jsonMap) > 0 {
		parsedJson, _ := json.Marshal(jsonMap[0]) 
		w.Write(parsedJson)
	}
}