package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called at the beginning of each block
func (k Keeper) EndBlocker(ctx context.Context) error {

	// 从ctx中解包出sdkCtx
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// 获取日志记录器
	logger := k.Logger(sdkCtx)

	fmt.Printf("---------------EndBlocker was called at height: %d\n", sdkCtx.BlockHeight())
	logger.Info("-------------------------------myendblocker-----abci--EndBlocker-----")
	// 抛出事件
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		"myendblocker",
		sdk.NewAttribute("height", fmt.Sprintf("%d", sdkCtx.BlockHeight())),
	))

	// 获取所有事件
	events := sdkCtx.EventManager().Events()

	// 定义要查找的事件类型
	eventType := sdk.EventTypeTx
	eventAttributeKey := "myendblocker"

	// 遍历所有事件，查找特定类型的事件
	for _, event := range events {
		if event.Type == eventType {
			for _, attribute := range event.Attributes {
				if attribute.Key == eventAttributeKey {
					// 找到了指定的事件，执行相关逻辑
					logger.Info("-----potv----EndBlocker--------Found the event: %s", attribute.Value)
					break
				}
			}
		}
	}

	//扫描区块以找到合约
	/*for _, event := range events {
		if event.Type == sdk.EventTypeTx && event.Attributes[0].Key == "mycontract" {
			// 找到了合约，执行相关逻辑
			// ...
		}
	}*/

	// 处理费用、验证者状态更新和奖励处理等逻辑
	// ...

	// 返回验证器更新（如果需要）
	//validatorUpdates := []abci.ValidatorUpdate{}

	return nil
}
