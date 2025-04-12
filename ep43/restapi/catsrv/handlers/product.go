package handlers

import (
	"catsrv/validators"
	"encoding/json"
	"errors"
	"net/http"
	"rfcrud/dal"
	"rfcrud/dal/catalog"
	"strconv"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := catalog.ListProducts()
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
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product, err := catalog.GetProduct(id)
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
	json.NewEncoder(w).Encode(product)
}

func ListProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products, err := catalog.ListProductsByCategoryID(id)
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
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = catalog.DeleteProduct(id)
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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	data, err := decodeProduct(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	prodMap, valid := validators.ValidateStruct(data)
	if !valid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(prodMap)
		return
	}

	prod := catalog.NewProduct(prodMap["name"].(string),
		prodMap["price"].(float64),
		int(prodMap["category_id"].(float64)))

	prodCreated, err := catalog.CreateProduct(prod)
	if err != nil {
		if err == dal.ErrNoRows {
			nferr := errors.New("category id not found")
			http.Error(w, nferr.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prodCreated)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	prod, err := catalog.GetProduct(id)
	if err != nil {
		if err == dal.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := decodeProduct(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	prodMap, valid := validators.ValidateStruct(data)
	if !valid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(prodMap)
		return
	}
	productFromMap(&prod, prodMap)
	prodUpdated, err := catalog.UpdateProduct(prod)
	if err != nil {
		if err == dal.ErrNoRows {
			nferr := errors.New("category id not found")
			http.Error(w, nferr.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prodUpdated)
}

func decodeProduct(r *http.Request) (any, error) {
	type Product struct {
		Name       string  `json:"name" validate:"required,min=6,max=75"`
		Price      float64 `json:"price" validate:"required"`
		CategoryID int     `json:"category_id" validate:"required,gte=1,fkcategoryid"`
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func productFromMap(prod *catalog.Product, prodMap map[string]any) {
	prod.SetName(prodMap["name"].(string))
	prod.SetPrice(prodMap["price"].(float64))
	prod.SetCategoryID(int(prodMap["category_id"].(float64)))
}
