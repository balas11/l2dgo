package catalog

import (
	"encoding/json"
)

type Category struct {
	id   int
	name string
}

func NewCategory(name string) Category {
	return Category{name: name}
}

func (c *Category) ID() int {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) SetName(name string) {
	c.name = name
}

func (c Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   c.id,
		Name: c.name,
	})
}

// func (c *Category) UnmarshalJSON(data []byte) error {
// 	log.Println("UnmarshalJSON")
// 	var cc struct {
// 		ID   int    `json:"id"`
// 		Name string `json:"name"`
// 	}
// 	if err := json.Unmarshal(data, &cc); err != nil {
// 		return err
// 	}
// 	c.id = cc.ID
// 	c.name = cc.Name
// 	return nil
// }

type Product struct {
	id       int
	name     string
	price    float64
	category *Category
}

func NewProduct(name string, price float64, categoryID int) Product {
	return Product{
		name:     name,
		price:    price,
		category: &Category{id: categoryID},
	}
}

func (p *Product) ID() int {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

func (p *Product) Category() *Category {
	return p.category
}

func (p *Product) SetCategory(category *Category) {
	p.category = category
}

func (p *Product) SetCategoryID(categoryID int) {
	p.category.id = categoryID
}

// func (p Product) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		ID       int     `json:"id"`
// 		Name     string  `json:"name"`
// 		Price    float64 `json:"price"`
// 		Category struct {
// 			ID   int    `json:"id"`
// 			Name string `json:"name"`
// 		} `json:"category"`
// 	}{
// 		ID:    p.id,
// 		Name:  p.name,
// 		Price: p.price,
// 		Category: struct {
// 			ID   int    `json:"id"`
// 			Name string `json:"name"`
// 		}{ID: p.category.id, Name: p.category.name},
// 	})
// }

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		Price    float64  `json:"price"`
		Category Category `json:"category"`
	}{
		ID:       p.id,
		Name:     p.name,
		Price:    p.price,
		Category: *p.category,
	})
}
