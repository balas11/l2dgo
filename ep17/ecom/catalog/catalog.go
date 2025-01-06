package catalog

type pricing struct {
	cost   float32
	markup float32
}

func (p *pricing) profit() float32 {
	return p.cost * p.markup / 100.0
}

type Product struct {
	Name    string
	pricing *pricing
}

func NewProduct(name string, cost float32, markup float32) *Product {
	return &Product{
		Name:    name,
		pricing: &pricing{cost, markup},
	}
}

func (product *Product) SellingPrice() float32 {
	return product.pricing.cost + product.pricing.profit()
}

func (product *Product) Profit() float32 {
	return product.pricing.profit()
}
