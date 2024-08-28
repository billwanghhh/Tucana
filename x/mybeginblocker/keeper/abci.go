package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called at the beginning of each block
func (k Keeper) BeginBlocker(ctx context.Context) error {

	// 从ctx中解包出sdkCtx
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// 获取日志记录器
	logger := k.Logger(sdkCtx)

	fmt.Printf("BeginBlocker was called at height: %d\n", sdkCtx.BlockHeight())
	logger.Info("-------------------------------mybeginblocker-----abci--BeginBlocker-----")
	// 抛出事件
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		"mybeginblocker",
		sdk.NewAttribute("height", fmt.Sprintf("%d", sdkCtx.BlockHeight())),
	))

	return nil
}
