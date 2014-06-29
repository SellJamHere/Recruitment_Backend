package services

import (
	// "fmt"

	"appengine"
	"appengine/datastore"

	"backend/errors"
	"backend/models"
)

func GetRecruits(context appengine.Context, includes string) ([]models.Recruit, *errors.ServerError) {
	query := datastore.NewQuery("recruit")

	var recruits []models.Recruit
	keys, err := query.GetAll(context, &recruits)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve recruits", 500)
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		recruit := recruits[i]
		recruit.EncodedKey = key.Encode()
		recruits[i] = recruit
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

func CreateRecruit(context appengine.Context, recruit *models.Recruit) (*datastore.Key, *errors.ServerError) {
	key := datastore.NewIncompleteKey(context, recruit.DatastoreKind(), nil)

	key, err := datastore.Put(context, key, recruit)
	if err != nil {
		return nil, errors.New(err, "Unable to create recruit", 500)
	}

	return key, nil
}

func UpdateRecruit(context appengine.Context, recruit *models.Recruit) *errors.ServerError {
	context.Infof("%v", recruit)
	key, keyErr := datastore.DecodeKey(recruit.EncodedKey)
	if keyErr != nil {
		return errors.New(keyErr, "Invalid Key", 400)
	}

	context.Infof("%v", key.String())

	err := datastore.RunInTransaction(context, func(context appengine.Context) error {
		getErr := datastore.Get(context, key, recruit)
		if getErr != nil {
			return getErr
		}

		_, putErr := datastore.Put(context, key, recruit)
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
