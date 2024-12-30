package commands

type ProductCommands struct {
	CreateProduct CreateProductCmdHandler
}

func NewProductCommands(createProduct CreateProductCmdHandler) *ProductCommands {
	return &ProductCommands{CreateProduct: createProduct}
}
