package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/TucanaProtocol/Tucana/v8/x/mybeginblocker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper defines the module's keeper
type Keeper struct {
	//storeKey sdk.StoreKey
	storeKey storetypes.StoreKey // 用于访问状态数据库的键
	//cdc      codec.Codec         // 用于编解码的编解码器
}

// NewKeeper creates a new Keeper instance
func NewKeeper(storeKey storetypes.StoreKey) Keeper {
	fmt.Printf("Mybeginblock Keeper----NewKeeper------------  ")
	return Keeper{storeKey: storeKey}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
