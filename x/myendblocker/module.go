package myendblocker

import (
	"encoding/json"
	"github.com/TucanaProtocol/Tucana/v8/x/myendblocker/keeper"
	"github.com/TucanaProtocol/Tucana/v8/x/myendblocker/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// AppModule is the AppModule for the mybeginblocker module
type AppModule struct {
	keeper keeper.Keeper
}

func (am AppModule) IsOnePerModuleType() {
	//TODO implement me
	panic("implement me")
}

func (am AppModule) IsAppModule() {
	//TODO implement me
	//panic("implement me")
}

func (am AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	//TODO implement me
	//panic("implement me")
}

func (am AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	//TODO implement me
	//panic("implement me")
}

func (am AppModule) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	//TODO implement me
	//panic("implement me")
}

// NewAppModule creates a new AppModule instance
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{keeper: k}
}

// Name returns the module name
func (AppModule) Name() string {
	return types.ModuleName
}

// Route returns the module's message routing key
func (AppModule) Route() string {
	return ""
}

// QuerierRoute returns the module's query routing key
func (AppModule) QuerierRoute() string {
	return ""
}

// LegacyQuerierHandler returns the module's legacy query handler
//func (AppModule) LegacyQuerierHandler(*codec.Codec) sdk.Querier {
//	return nil
//}

// RegisterInvariants registers the module's invariants
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis initializes the module's genesis state
func (AppModule) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

// ExportGenesis exports the module's genesis state
func (AppModule) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) json.RawMessage {
	return nil
}

// BeginBlock is called at the beginning of each block
//func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
//	am.keeper.BeginBlocker(ctx)
//	return abci.ResponseBeginBlock{}
//
//	return am.keeper.BeginBlocker(ctx)
//}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.Request) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	logger := am.keeper.Logger(sdkCtx)

	logger.Info("-------------------------------myendnblocker---module----EndBlock-----")
	return am.keeper.EndBlocker(ctx)
	//return EndBlocker(ctx, am.keeper)
}
