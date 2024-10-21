package repository

import (
	"database/sql"
	"fmt"
	"server/model"
)

type DbAdapter[T any] struct {
	db      *sql.DB
	factory model.Factory[T]
}

func CreateDbAdapter[T any](db *sql.DB, factory model.Factory[T]) *DbAdapter[T] {
	return &DbAdapter[T]{
		db:      db,
		factory: factory,
	}
}

func (a *DbAdapter[T]) ExecuteForAll(query string) ([]T, error) {
	return a.Execute(query, nil)
}

func (a *DbAdapter[T]) Execute(query string, args []interface{}) ([]T, error) {
	rows, err := a.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not get result: %v", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var result []T
	for rows.Next() {
		o, fields := a.factory.CreateTemplate()
		err := rows.Scan(fields...)
		if err != nil {
			return nil, fmt.Errorf("could not scan model: %v", err)
		}
		result = append(result, *o)
	}
	return result, nil
}

func (a *DbAdapter[T]) ExecuteSingle(query string, args []interface{}) (T, error) {
	o, fields := a.factory.CreateTemplate()
	err := a.db.QueryRow(query, args...).Scan(fields...)
	if err != nil {
		return *o, fmt.Errorf("could not scan model: %v", err)
	}
	return *o, nil
}

func (a *DbAdapter[T]) ExecuteUpdatingField(query string, args []interface{}, field interface{}) error {
	return a.ExecuteWithFields(query, args, []interface{}{field})
}

func (a *DbAdapter[T]) ExecuteWithFields(query string, args []interface{}, fields []interface{}) error {
	err := a.db.QueryRow(query, args...).Scan(fields...)
	if err != nil {
		return fmt.Errorf("could not scan model: %v", err)
	}
	return nil
}
