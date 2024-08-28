package keeper

//1. todo: build the endblocker
//  demo: inflation
import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker is the ABCI EndBlocker for the potv module
func (k Keeper) EndBlocker(ctx context.Context) error {

	// Unpack sdkCtx from ctx
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// Get a logger
	logger := k.Logger(sdkCtx)

	fmt.Printf("-----------potv----EndBlocker was called at height: %d\n", sdkCtx.BlockHeight())
	logger.Info("-------------------------------potv-----abci--EndBlocker-----")

	fmt.Printf("-----------potv----chainId: %s, gasMeter:%v, BlockGasMeter:%v \n",
		sdkCtx.ChainID(), sdkCtx.GasMeter(), sdkCtx.BlockGasMeter())

	// EmitEvent emits an event with the given type and attributes.
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		// Event type
		"potv_end",
		// Attribute with key "height" and value of the current block height
		sdk.NewAttribute("height", fmt.Sprintf("%d", sdkCtx.BlockHeight())),
	))

	// 获取所有事件
	events := sdkCtx.EventManager().Events()

	// 定义要查找的事件类型
	eventType := sdk.EventTypeTx
	eventAttributeKey := "height"

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
