package keeper

import (
	"safrochain/x/safrochain/types"
)

var _ types.QueryServer = Keeper{}
