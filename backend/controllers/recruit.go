package controllers

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"

	"backend/errors"
	"backend/handlers"
	"backend/models"
	"backend/services"

	"github.com/codegangsta/martini"
)

func init() {
	m.Get("/recruits", getRecruits)
	m.Get("/recruits/:key", getRecruit)
	m.Post("/recruits", createRecruit)
	// m.Post("/recruits", createMultipleRecruits)
	m.Patch("/recruits", updateRecruit)
	m.Delete("/recruits/:key", deleteRecruit)
}

func getRecruits(r handlers.Respond, req *http.Request) {
	recruits, err := services.GetRecruits(appengine.NewContext(req), req.FormValue("updated_at"))
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, recruits)
}

func getRecruit(params martini.Params, r handlers.Respond, req *http.Request) {
	key, err := datastore.DecodeKey(params["key"])
	if err != nil {
		r.Error(errors.New(err, "Invalid key", 400))
		return
	}

	recruit, datastoreErr := services.GetRecruit(appengine.NewContext(req), key)
	if err != nil {
		r.Error(datastoreErr)
		return
	}

	r.Valid(200, recruit)
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
	recruit.UpdatedAt = time.Now()

	//check if email already used
	recruits, emailErr := services.GetRecruitByEmail(appengine.NewContext(req), recruit.Email)
	if emailErr != nil {
		r.Error(emailErr)
		return
	}

	appengine.NewContext(req).Infof("%v", recruits)

	if len(recruits) > 0 {
		r.Error(errors.New(nil, "Email already registered", 499))
		return
	}

	recruitKey, serverErr := services.CreateRecruit(appengine.NewContext(req), recruit)
	if serverErr != nil {
		r.Error(serverErr)
		return
	}

	recruit.EncodedKey = recruitKey.Encode()

	r.Valid(200, recruit)
}

func createMultipleRecruits(r handlers.Respond, req *http.Request) {
	r.Valid(200, "createMultipleRecruits")
}

func updateRecruit(r handlers.Respond, req *http.Request) {
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
	recruit.UpdatedAt = time.Now()

	updateErr := services.UpdateRecruit(appengine.NewContext(req), recruit)
	if updateErr != nil {
		r.Error(updateErr)
		return
	}

	r.Valid(200, "")
}

func deleteRecruit(params martini.Params, r handlers.Respond, req *http.Request) {
	key, keyErr := datastore.DecodeKey(params["key"])
	if keyErr != nil {
		r.Error(errors.New(keyErr, "Invalid key", 400))
		return
	}

	err := services.DeleteRecruit(appengine.NewContext(req), key)
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, "")
}
