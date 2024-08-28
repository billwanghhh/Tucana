package mybeginblocker

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type AppModuleBasic123 struct {
}

func (a AppModuleBasic123) IsOnePerModuleType() {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic123) IsAppModule() {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic123) Name() string {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic123) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic123) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}

func (a AppModuleBasic123) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	//TODO implement me
	panic("implement me")
}

//func (a AppModuleBasic123) IsOnePerModuleType() {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a AppModuleBasic123) IsAppModule() {
//	//TODO implement me
//	panic("implement me")
//}
