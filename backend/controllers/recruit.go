package controllers

import (
	"appengine"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"backend/errors"
	"backend/handlers"
	"backend/models"
	"backend/services"

	"github.com/codegangsta/martini"
)

func init() {
	m.Get("/recruits", getRecruits)
	m.Get("/recruit/:id", getRecruit)
	m.Post("/recruit", createRecruit)
}

func getRecruits(r handlers.Respond, req *http.Request) {
	recruits, err := services.GetRecruits(appengine.NewContext(req), req.FormValue("include"))
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, recruits)
}

func getRecruit(params martini.Params, r handlers.Respond, req *http.Request) {
	r.Valid(200, fmt.Sprintf("Get Recruit: %s", params["id"]))
}

func createRecruit(r handlers.Respond, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.Error(errors.New(err, "Unable to read body bytes", 400))
		return
	}

	recruit := &models.Recruit{}
	err = json.Unmarshal(bodyBytes, recruit)
	if err != nil {
		r.Error(errors.New(err, "Unable to unmarshal donor into object", 500))
		return
	}

	recruitKey, serverErr := services.CreateRecruit(appengine.NewContext(req), recruit)
	if serverErr != nil {
		r.Error(serverErr)
		return
	}

	r.Valid(200, recruitKey)
}
