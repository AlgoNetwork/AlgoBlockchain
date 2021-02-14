package keeper

import (
	"github.com/AlgoNetwork/AlgoBlockchain/x/AlgoBlockchain/types"
)

var _ types.QueryServer = Keeper{}
