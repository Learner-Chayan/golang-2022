package controllers

import (
	"encoding/json"
	"github.com/rest-go-example/pkg/models"
	"io/ioutil"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	person := &models.Person{}
	if requestBodyInByte, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(requestBodyInByte, person); err != nil {
			return
		}

		p := person.CreatePerson()
		res, _ := json.Marshal(p)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	return

}
