package product

import domain "mvi/internal/domain/entity/product"

type Product interface {
	// find baranag base SKU
	FindProductWithSKU(sku string) (*domain.Product, error)
}