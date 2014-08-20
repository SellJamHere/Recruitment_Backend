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
	IsMale         GAEBool   `datastore:"male" json:"male,omitempty"`
	DateRegistered time.Time `datastore:"date_registered" json:"date_registered,omitempty"`
	// interest_level  int
}

func (r *Recruit) DatastoreKind() string {
	return "Recruit"
}

func (r *Recruit) UpdateFieldsWithRecruit(newFields *Recruit) {
	if newFields.Name != "" {
		r.Name = newFields.Name
	}
	if newFields.Email != "" {
		r.Email = newFields.Email
	}
	if newFields.Year != 0 {
		r.Year = newFields.Year
	}
	if newFields.Major != "" {
		r.Major = newFields.Major
	}
	if r.IsMale.Bool() != newFields.IsMale.Bool() {
		r.IsMale = newFields.IsMale
	}
}
