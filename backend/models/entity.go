package models

type Entity interface {
	DatastoreKind() string
}
