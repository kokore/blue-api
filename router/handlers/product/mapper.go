package product

import "blue-api/internal/repository/product"

func ToCreatedProductResponseDTO(product *ProductRequest) *ProductResponseDTO {
	return &ProductResponseDTO{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Image:    product.Image,
	}
}

func ToUpdateProductResponseDTO(product *ProductRequest) *ProductResponseDTO {
	return &ProductResponseDTO{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Image:    product.Image,
	}
}

func ToProductsResponseDTO(products []product.Product) *ProductsResponseDTO {
	return &ProductsResponseDTO{
		Products: products,
	}
}
