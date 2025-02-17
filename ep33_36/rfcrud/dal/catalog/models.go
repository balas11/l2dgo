package catalog

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
