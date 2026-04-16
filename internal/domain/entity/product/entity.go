package product

import "errors"

type Product struct {
	sku   string
	nama  string
	harga int
	stock int
}

func (p *Product) ReduceStock(qty int) error {
	if qty <= 0 {
		return errors.New("qty tidak boleh negatif")
	}

	if qty > p.stock {
		return errors.New("stock tidak cukup")
	}

	p.stock -= qty
	return nil
}

func (p *Product) UpdatePrice(price int) error {
	if price <= 0 {
		return errors.New("harga tidak boleh negatif")
	}

	p.stock = price
	return nil
}