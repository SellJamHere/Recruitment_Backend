package services

import (
	// "fmt"
	"appengine"
	"appengine/datastore"

	"backend/errors"
	"backend/models"
)

func GetRecruits(context appengine.Context, includes string) ([]models.Recruit, *errors.ServerError) {
	query := datastore.NewQuery("Recruit")

	var recruits []models.Recruit
	_, err := query.GetAll(context, &recruits)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve recruits", 500)
	}

	return recruits, nil
}

func CreateRecruit(context appengine.Context, recruit *models.Recruit) (*datastore.Key, *errors.ServerError) {
	key := datastore.NewIncompleteKey(context, recruit.DatastoreKind(), nil)

	key, err := datastore.Put(context, key, recruit)
	if err != nil {
		return nil, errors.New(err, "Unable to create recruit", 500)
	}

	return key, nil
}
