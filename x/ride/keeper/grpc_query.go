package keeper

import (
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

var _ types.QueryServer = Keeper{}
