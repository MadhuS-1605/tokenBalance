package models


// BalanceResponse data structure
type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

// Error data structure
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}