package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/TucanaProtocol/Tucana/v8/x/myendblocker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper defines the module's keeper
type Keeper struct {
	//storeKey sdk.StoreKey
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new Keeper instance
func NewKeeper(storeKey storetypes.StoreKey) Keeper {
	return Keeper{storeKey: storeKey}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
