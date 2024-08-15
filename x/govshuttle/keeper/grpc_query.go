package keeper

import (
	"github.com/TucanaProtocol/Canto/v8/x/govshuttle/types"
)

var _ types.QueryServer = Keeper{}
