package main

import (
	"fmt"
	"log"
	"rfcrud/dal"
	"rfcrud/dal/catalog"
)

func main() {

	dbURI := "postgres://postgres:postgres@localhost:5432/store?sslmode=disable"

	err := dal.OpenDB(dbURI)
	if err != nil {
		panic(err)
	}
	defer dal.CloseDB()
	if err := dal.PingTest(); err != nil {
		log.Println(err)
	}

	// ListProducts()
	// GetProduct()
	// CreateProduct()
	// UpdateProduct()
	// DeleteProduct()
	// ListProducts()
	// ListProductsByCategoryID()

	// create category
	// category := catalog.NewCategory("Electronics")
	// insertedCategory,err = catalog.CreateCategory(category)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("id: %d, name: %s\n", insertedCategory.ID(), insertedCategory.Name())

	// List
	// categories, err := catalog.ListCategories()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for _, category := range categories {
	// 	fmt.Printf("id: %d, name: %s\n", category.ID(), category.Name())
	// }

	// Get
	category, err := catalog.GetCategory(4)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("id: %d, name: %s\n", category.ID(), category.Name())

	// Get & Update
	// category.SetName("Electronic Gadgets")
	// updatedCategory, err := catalog.UpdateCategory(category)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("id: %d, name: %s\n", updatedCategory.ID(), updatedCategory.Name())

	// Delete
	// if err := catalog.DeleteCategory(12); err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("deleted")
}

func ListProducts() {
	products, err := catalog.ListProducts()
	if err != nil {
		log.Println(err)
		return
	}
	for _, product := range products {
		fmt.Printf("id: %d, name: %s, price: %.2f, category_id: %d, category_name: %s\n",
			product.ID(), product.Name(), product.Price(), product.Category().ID(), product.Category().Name())
	}
}

func ListProductsByCategoryID() {
	products, err := catalog.ListProductsByCategoryID(1)
	if err != nil {
		log.Println(err)
		return
	}
	for _, product := range products {
		fmt.Printf("id: %d, name: %s, price: %.2f, category_id: %d, category_name: %s\n",
			product.ID(), product.Name(), product.Price(), product.Category().ID(), product.Category().Name())
	}
}

func GetProduct() {
	product, err := catalog.GetProduct(2)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("id: %d, name: %s, price: %.2f, category_id: %d, category_name: %s\n",
		product.ID(), product.Name(), product.Price(), product.Category().ID(), product.Category().Name())
}

func CreateProduct() {

	product := catalog.NewProduct("Milton Flask", 549.0, 2)
	insertedProduct, err := catalog.CreateProduct(product)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("id: %d, name: %s, price: %.2f, category_id: %d, category_name: %s\n",
		insertedProduct.ID(), insertedProduct.Name(), insertedProduct.Price(), insertedProduct.Category().ID(), insertedProduct.Category().Name())
}

func UpdateProduct() {

	product, err := catalog.GetProduct(5)
	if err != nil {
		log.Println(err)
		return
	}

	product.SetName("Butterfly Mixer/Grinder")

	updatedProduct, err := catalog.UpdateProduct(product)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("id: %d, name: %s, price: %.2f, category_id: %d, category_name: %s\n",
		updatedProduct.ID(), updatedProduct.Name(), updatedProduct.Price(), updatedProduct.Category().ID(), updatedProduct.Category().Name())
}

func DeleteProduct() {
	err := catalog.DeleteProduct(8)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("deleted")
}
