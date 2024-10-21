package repository

import (
	"fmt"
	. "server/model"
)

type TagRepository struct {
	dbAdapter *DbAdapter[Tag]
}

func CreateTagRepository(dbAdapter *DbAdapter[Tag]) *TagRepository {
	return &TagRepository{
		dbAdapter: dbAdapter,
	}
}

func (r *TagRepository) Insert(tag Tag) (Tag, error) {
	query := "INSERT INTO public.tag (name) VALUES ($1) RETURNING id"
	err := r.dbAdapter.ExecuteWithFields(query, []interface{}{tag.Name}, []interface{}{&tag.ID})
	if err != nil {
		return Tag{}, fmt.Errorf("could not insert tag: %v", err)
	}
	return tag, nil
}

func (r *TagRepository) GetAll() ([]Tag, error) {
	query := "SELECT id, name FROM public.tag"
	return r.dbAdapter.ExecuteForAll(query)
}

func (r *TagRepository) GetByIds(ids []int) ([]Tag, error) {
	query := "SELECT id, name FROM public.tag WHERE id IN ($1)"
	return r.dbAdapter.Execute(query, []interface{}{ids})
}
