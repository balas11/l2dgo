package handlers

import (
	"catsrv/validators"
	"encoding/json"
	"net/http"
	"rfcrud/dal"
	"rfcrud/dal/catalog"
	"strconv"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := catalog.ListCategories()
	if err != nil {
		if err == dal.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	category, err := catalog.GetCategory(id)
	if err != nil {
		if err == dal.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = catalog.DeleteCategory(id)
	if err != nil {
		if err == dal.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func Createcategory(w http.ResponseWriter, r *http.Request) {
	data, err := decodeCategory(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	catMap, valid := validators.ValidateStruct(data)
	if !valid {
		//write to response with status unprocessable entity
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(catMap)
		return
	}
	cat := catalog.NewCategory(catMap["name"].(string))
	catCreated, err := catalog.CreateCategory(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(catCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cat, err := catalog.GetCategory(id)
	if err != nil {
		if err == dal.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := decodeCategory(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	catMap, valid := validators.ValidateStruct(data)
	if !valid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(catMap)
		return
	}
	categoryFromMap(&cat, catMap)
	catUpdated, err := catalog.UpdateCategory(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(catUpdated)
}

func decodeCategory(r *http.Request) (any, error) {
	type Category struct {
		Name string `json:"name" validate:"required,min=6,max=75"`
	}
	var category Category

	err := json.NewDecoder(r.Body).Decode(&category)
	return category, err
}

func categoryFromMap(cat *catalog.Category, catMap map[string]any) {
	cat.SetName(catMap["name"].(string))
}
