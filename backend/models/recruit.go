package models

import (
	"time"
)

type Recruit struct {
	EncodedKey     string    `datastore:"-" json:"key,omitempty"`
	Name           string    `datastore:"name" json:"name,omitempty"`
	Email          string    `datastore:"email" json:"email,omitempty"`
	Year           int       `datastore:"year" json:"year,omitempty"`
	Major          string    `datastore:"major" json:"major,omitempty"` //Possibly make this a dropdown with enumerated majors
	IsMale         bool      `datastore:"male" json:"male,omitempty"`
	DateRegistered time.Time `datastore:"date_registered" json:"date_registered,omitempty"`
	// interest_level  int
}

func (r *Recruit) DatastoreKind() string {
	return "recruit"
}
