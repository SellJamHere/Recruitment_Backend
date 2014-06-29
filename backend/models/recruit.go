package models

import (
	"time"
)

type Recruit struct {
	Id             string    `datastore:"id" json:"id,omitempty"`
	Name           string    `datastore:"name" json:"name,omitempty"`
	Email          string    `datastore:"email" json:"email,omitempty"`
	Year           int       `datastore:"year,noindex" json:"year,omitempty"`
	Major          string    `datastore:"major,noindex" json:"major,omitempty"` //Possibly make this a dropdown with enumerated majors
	DateRegistered time.Time `datastore:"date_registered,noindex" json:"date_registered,omitempty"`
	// interest_level  int
}

func (r *Recruit) DatastoreKind() string {
	return "recruit"
}
