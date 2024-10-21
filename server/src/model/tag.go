package model

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagFactory struct{}

func CreateTagFactory() TagFactory {
	return TagFactory{}
}

func (TagFactory) CreateTemplate() (*Tag, []interface{}) {
	var tag Tag
	return &tag, []interface{}{&tag.ID, &tag.Name}
}
