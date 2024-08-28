package mybeginblocker

import (
	"context"
	"fmt"
	"github.com/TucanaProtocol/Tucana/v8/x/mybeginblocker/keeper"
	"github.com/TucanaProtocol/Tucana/v8/x/mybeginblocker/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// AppModule implements the AppModule interface for the mybeginblocker module.
type AppModule struct {
	// TODO implement me
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper keeper.Keeper) AppModule {
	return AppModule{keeper: keeper}
}

func (am AppModule) IsOnePerModuleType() {
	//TODO implement me
	fmt.Println("-----mybeginblock module-------------------IsOnePerModuleType")
}

func (am AppModule) IsAppModule() {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----mybeginblock module-------------------IsAppModule")
}

// Name Returns the name of the module
func (am AppModule) Name() string {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----mybeginblock module-------------------Name")
	return types.ModuleName
}

func (am AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----mybeginblock module-------------------RegisterLegacyAminoCodec")
}

func (am AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----mybeginblock module-------------------RegisterInterfaces")
}

func (am AppModule) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----mybeginblock module-------------------RegisterGRPCGatewayRoutes")
}

// BeginBlock is a method that will be run before transactions are processed in a block.
func (am AppModule) BeginBlock(ctx context.Context) error {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----mybeginblock module-------------------BeginBlock")
	return am.keeper.BeginBlocker(ctx)
	//return nil
}
