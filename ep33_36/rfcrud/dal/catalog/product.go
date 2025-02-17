package catalog

import "rfcrud/dal"

func CreateProduct(product Product) (Product, error) {

	category, err := GetCategory(product.category.id)
	if err != nil {
		return Product{}, err
	}

	sql := `insert into products(name, price, category_id) 
		values($1, $2, $3) returning id`

	row := dal.Db.QueryRow(sql, product.name, product.price, product.category.id)
	err = row.Scan(&product.id)
	if err != nil {
		return Product{}, err
	}
	product.category = &category
	return product, nil
}

func ListProducts() ([]Product, error) {
	sql := `select products.id, products.name, products.price, 
		categories.id, categories.name from products
		join categories on products.category_id = categories.id`

	rows, err := dal.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		var category Category
		if err := rows.Scan(&product.id, &product.name, &product.price,
			&category.id, &category.name); err != nil {
			return nil, err
		}
		product.category = &category
		products = append(products, product)
	}
	return products, nil
}

func GetProduct(id int) (Product, error) {
	sql := `select products.id, products.name, products.price, 
		categories.id, categories.name from products
		join categories on products.category_id = categories.id
		where products.id = $1`

	row := dal.Db.QueryRow(sql, id)
	var product Product
	var category Category
	if err := row.Scan(&product.id, &product.name, &product.price,
		&category.id, &category.name); err != nil {
		return Product{}, err
	}
	product.category = &category
	return product, nil
}

func UpdateProduct(product Product) (Product, error) {
	//Check if category exists
	category, err := GetCategory(product.category.id)
	if err != nil {
		return Product{}, err
	}
	product.category = &category
	sql := `update products set name = $1, price = $2, category_id = $3 where id = $4`

	_, err = dal.Db.Exec(sql, product.name, product.price, product.category.id, product.id)

	if err != nil {
		return Product{}, err
	}

	UpdateProduct, err := GetProduct(product.id)
	if err != nil {
		return Product{}, err
	}
	return UpdateProduct, nil
}

func DeleteProduct(id int) error {
	sql := `delete from products where id = $1`
	_, err := dal.Db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func ListProductsByCategoryID(categoryID int) ([]Product, error) {
	sql := `select products.id, products.name, products.price, 
		categories.id, categories.name from products
		join categories on products.category_id = categories.id
		where products.category_id = $1`

	rows, err := dal.Db.Query(sql, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		var category Category
		if err := rows.Scan(&product.id, &product.name, &product.price,
			&category.id, &category.name); err != nil {
			return nil, err
		}
		product.category = &category
		products = append(products, product)
	}
	return products, nil
}
