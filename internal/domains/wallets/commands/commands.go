package commands

func NewWalletCommands(createWallet CreateWalletCmdHandler) *WalletCommands {
	return &WalletCommands{CreateWallet: createWallet}
}
