package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/TucanaProtocol/Tucana/v8/x/potv/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// todo 2. build keeper
// demo: //  demo: inflation

// Keeper maintains the potv state
type Keeper struct {
	//cdc      codec.Codec         // The wire codec for binary encoding/decoding.
	storeKey storetypes.StoreKey // key to access store from Context
}

// NewKeeper creates a new potv Keeper instance
func NewKeeper(storeKey storetypes.StoreKey) Keeper {
	return Keeper{storeKey: storeKey}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
