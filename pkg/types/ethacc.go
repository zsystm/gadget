package types

import (
	"github.com/cometbft/cometbft/libs/json"
	"github.com/ethereum/go-ethereum/common"
)

type EthAcc struct {
	EthPrivKey string         `json:"privateKey"`
	EthAddr    common.Address `json:"address"`
}

// implement json marshaler for EthAcc
func (e *EthAcc) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EthPrivKey string `json:"privateKey"`
		EthAddr    string `json:"address"`
	}{
		EthPrivKey: e.EthPrivKey,
		EthAddr:    e.EthAddr.Hex(),
	})
}
