package model

type Factory[T any] interface {
	CreateTemplate() (*T, []interface{})
}
