package product

import "blue-api/internal/repository/product"

func ToCreatedProductResponseDTO(product *CreateProductRequest) *CreateProductResponseDTO {
	return &CreateProductResponseDTO{
		Name:         product.Name,
		Price:        product.Price,
		CurrentStock: product.CurrentStock,
		Image:        product.Image,
	}
}

func ToProductsResponseDTO(products []product.Product) *ProductsResponseDTO {
	return &ProductsResponseDTO{
		Products: products,
	}
}
