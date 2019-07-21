package person

import (
	"encoding/json"
	"net/http"
	config "rest-api/rest-api/config"
	models "rest-api/rest-api/models"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		personModel := models.PersonModel{
			Db: db,
		}
		persons, err2 := personModel.FindAll()

		if err != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, persons)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
