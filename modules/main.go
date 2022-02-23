package modules
import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	 token "tokenBal/build"
)

// GetAddressBalance returns the given address balance =P
func GetAddressBalance(tokenaddress string, address string) (string, error) {
	
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a9126c60ad5346e79ff1061816be1dd4")

	if err != nil {
		return "0", err
	}

	tokenAddress := common.HexToAddress(tokenaddress)
	instance, err1 := token.NewToken(tokenAddress, client)
	if err1 != nil {
		return "0", err1
	}
	account := common.HexToAddress(address)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return "0", err
	}
	

	return balance.String(), nil
}