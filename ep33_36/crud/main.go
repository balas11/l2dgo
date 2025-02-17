package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Category struct {
	id   int
	name string
}

var db *sql.DB

func main() {
	var err error
	dbURI := "postgres://postgres:postgres@localhost:5432/store?sslmode=disable"

	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Step 1 Connect and check
	if err = db.Ping(); err != nil {
		log.Println(err)
		return
	}
	log.Println("connected to database")

	// Step2 - Query all
	categories, err := ListCategories()
	if err != nil {
		log.Println(err)
		return
	}
	for _, cat := range categories {
		fmt.Printf("id: %d, name: %s\n", cat.id, cat.name)
	}

	//Step3 - Query one
	// category, err := GetCategory(2)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("id: %d, name: %s\n", category.id, category.name)

	//step4 create a category
	// category := Category{
	// 	name: "Mobile Accessories",
	// }
	// category, err = CreateCategory(category)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("id: %d, name: %s\n", category.id, category.name)

	//Step5 - Update
	// category, err := GetCategory(4)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// category.name = "Electronic Gadgets"
	// resultCategory, err := UpdateCategory(category)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("id: %d, name: %s\n", resultCategory.id, resultCategory.name)

	//Step6 - Delete
	// err = DeleteCategory(4)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("deleted category")
}

func CreateCategory(category Category) (Category, error) {
	sql := "insert into categories(name) values($1) returning id"
	row := db.QueryRow(sql, category.name)

	if err := row.Scan(&category.id); err != nil {
		return Category{}, err
	}
	fmt.Println(category)
	return category, nil
}

func ListCategories() ([]Category, error) {
	sql := "select id, name from categories"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.id, &category.name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategory(id int) (Category, error) {
	sql := "select id, name from categories where id = $1 LIMIT 1"
	row := db.QueryRow(sql, id)
	var category Category
	if err := row.Scan(&category.id, &category.name); err != nil {
		return Category{}, err
	}
	return category, nil
}

func UpdateCategory(category Category) (Category, error) {
	sql := "update categories set name = $1 where id = $2"
	_, err := db.Exec(sql, category.name, category.id)
	if err != nil {
		return Category{}, err
	}

	updatedCategory, err := GetCategory(category.id)
	if err != nil {
		return Category{}, err
	}
	return updatedCategory, nil
}

func DeleteCategory(id int) error {
	sql := "delete from categories where id = $1"
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
