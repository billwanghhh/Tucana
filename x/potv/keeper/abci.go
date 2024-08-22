package keeper

//1. todo: build the endblocker
//  demo: inflation
import (
	"context"
	"github.com/TucanaProtocol/Tucana/v8/x/potv/types"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlockerHandler is the ABCI BeginBlock handler for the potv module
func (k Keeper) BeginBlockerHandler(ctx context.Context) {

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.BeginBlocker(sdkCtx)

	logger := k.Logger(sdkCtx)
	logger.Info("-------------------------------bill------------")
}
