package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/controller/util"
	. "server/model"
	. "server/repository"
)

type TagController struct {
	repository *TagRepository
}

func CreateTagController(repository *TagRepository) *TagController {
	return &TagController{
		repository: repository,
	}
}

func (c *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	var newTag Tag
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tag, err := c.repository.Insert(newTag)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	util.SetHeaders(w)
	json.NewEncoder(w).Encode(tag)
}

func (c *TagController) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := c.repository.GetAll()
	if err != nil {
		log.Printf("Internal Server Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SetHeaders(w)
	json.NewEncoder(w).Encode(tags)
}
