package queries

type WalletQueries struct {
	GetWalletByUserId GetWalletByUserIdHandler
}

func NewWalletQueries(getWalletByUserId GetWalletByUserIdHandler) *WalletQueries {
	return &WalletQueries{GetWalletByUserId: getWalletByUserId}
}
