package services

import (
	// "fmt"
	"time"

	"appengine"
	"appengine/datastore"

	"backend/errors"
	"backend/models"
)

func GetRecruits(context appengine.Context, updatedAt string) ([]*models.Recruit, *errors.ServerError) {
	recruit := &models.Recruit{}
	var query *datastore.Query

	updatedTime, timeParseErr := time.Parse(time.RFC3339, updatedAt)

	if timeParseErr == nil {
		context.Infof("updatedAt: %v", updatedTime)
		query = datastore.NewQuery(recruit.DatastoreKind()).Filter("updated_at >", updatedTime).Order("updated_at")
	} else {
		context.Infof("updatedAt not sent. timeParseErr: %v", timeParseErr)
		query = datastore.NewQuery(recruit.DatastoreKind())
	}

	context.Infof("query: %v", query)

	var recruits []*models.Recruit
	keys, err := query.GetAll(context, &recruits)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve recruits", 500)
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		recruit := recruits[i]
		recruit.EncodedKey = key.Encode()
	}

	return recruits, nil
}

func GetRecruit(context appengine.Context, key *datastore.Key) (*models.Recruit, *errors.ServerError) {
	recruit := &models.Recruit{}

	err := datastore.Get(context, key, recruit)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve recruit", 500)
	}

	return recruit, nil
}

func GetRecruitByEmail(context appengine.Context, email string) ([]*models.Recruit, *errors.ServerError) {
	recruit := &models.Recruit{}
	query := datastore.NewQuery(recruit.DatastoreKind()).Filter("email = ", email)

	var recruits []*models.Recruit
	keys, err := query.GetAll(context, &recruits)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve recruits", 500)
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		recruit := recruits[i]
		recruit.EncodedKey = key.Encode()
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

func UpdateRecruit(context appengine.Context, recruit *models.Recruit) *errors.ServerError {
	//Decodes recruit's string key into datastore.Key
	key, keyErr := datastore.DecodeKey(recruit.EncodedKey)
	if keyErr != nil {
		return errors.New(keyErr, "Invalid Key", 400)
	}

	err := datastore.RunInTransaction(context, func(context appengine.Context) error {
		storedRecruit := &models.Recruit{}
		getErr := datastore.Get(context, key, storedRecruit)
		if getErr != nil {
			return getErr
		}

		storedRecruit.UpdateFieldsWithRecruit(recruit)

		_, putErr := datastore.Put(context, key, storedRecruit)
		return putErr
	}, nil)
	if err != nil {
		return errors.New(err, "Unable to update recruit", 500)
	}

	return nil
}

func DeleteRecruit(context appengine.Context, key *datastore.Key) *errors.ServerError {
	err := datastore.Delete(context, key)
	if err != nil {
		return errors.New(err, "Unable to delete recruit", 500)
	}

	return nil
}
