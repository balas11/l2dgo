package catalog

import (
	"log"
	"rfcrud/dal"
)

func CreateCategory(category Category) (Category, error) {
	sql := "insert into categories(name) values($1) returning id"
	row := dal.Db.QueryRow(sql, category.name)

	if err := row.Scan(&category.id); err != nil {
		return Category{}, err
	}
	return category, nil
}

func ListCategories() ([]Category, error) {
	sql := "select * from categories"
	rows, err := dal.Db.Query(sql)
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
	sql := "select * from categories where id = $1 LIMIT 1"
	row := dal.Db.QueryRow(sql, id)
	var category Category
	if err := row.Scan(&category.id, &category.name); err != nil {
		return Category{}, err
	}
	return category, nil
}

func UpdateCategory(category Category) (Category, error) {
	sql := "update categories set name = $1 where id = $2"
	_, err := dal.Db.Exec(sql, category.name, category.id)
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
	_, err := GetCategory(id)
	if err != nil {
		return err
	}
	pc, err := CountProductsByCategoryID(id)
	if err != nil {
		return err
	}
	if pc > 0 {
		log.Printf("Category %d has %d products. Can't delete it.", id, pc)
		return dal.ErrCatHasProducts
	}
	sql := "delete from categories where id = $1"
	_, err = dal.Db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
