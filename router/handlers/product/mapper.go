package product

func ToCreatedProductResponseDTO(product *CreateProductRequest) *CreateProductResponseDTO {
	return &CreateProductResponseDTO{
		Name:         product.Name,
		Price:        product.Price,
		CurrentStock: product.CurrentStock,
		Image:        product.Image,
	}
}
