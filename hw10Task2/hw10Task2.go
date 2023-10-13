//  example how to use
//  curl -X POST localhost:8080/translate -H 'Content-Type: application/json' -d '{"text":"Hello World","sourceLang":"en", "targetLang":"uk"}'

package main

import (
	"encoding/json"
	"net/http"

	"github.com/bregydoc/gtranslate"
)

type TranslateRequest struct {
	Text       string
	SourceLang string
	TargetLang string
}

type TranslateResponse struct {
	TranslatedText string
}

func main() {
	http.HandleFunc("/translate", translateHandler)
	http.ListenAndServe(":8080", nil)
}

func translateHandler(w http.ResponseWriter, r *http.Request) {
	var request TranslateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "request error", http.StatusBadRequest)
		return
	}

	// https://github.com/bregydoc/gtranslate

	translated, err := gtranslate.TranslateWithParams(
		request.Text, gtranslate.TranslationParams{
			From: request.SourceLang,
			To:   request.TargetLang,
		},
	)
	if err != nil {
		http.Error(w, "translate error", http.StatusInternalServerError)
		return
	}

	response := TranslateResponse{TranslatedText: translated}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
