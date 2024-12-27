package queries

type WalletQueries struct {
	GetWalletById GetWalletByIdHandler
}

func NewWalletQueries(getWalletById GetWalletByIdHandler) *WalletQueries {
	return &WalletQueries{GetWalletById: getWalletById}
}
