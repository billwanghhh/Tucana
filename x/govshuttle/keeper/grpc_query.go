package keeper

import (
	"github.com/TucanaProtocol/Tucana/v8/x/govshuttle/types"
)

var _ types.QueryServer = Keeper{}
