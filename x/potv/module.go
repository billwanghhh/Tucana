package potv

import (
	"fmt"

	"github.com/TucanaProtocol/Tucana/v8/x/potv/keeper"
	"github.com/TucanaProtocol/Tucana/v8/x/potv/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

// AppModule represents the potv module type
type AppModule struct {
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper keeper.Keeper) AppModule {
	return AppModule{keeper: keeper}
}

func (am AppModule) IsOnePerModuleType() {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----potv module-------------------IsOnePerModuleType")
}

func (am AppModule) IsAppModule() {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----potv module-------------------IsAppModule")
}

// Name returns the potv module's name
func (am AppModule) Name() string {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----potv module-------------------Name")
	return types.ModuleName
}

func (am AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----potv module-------------------RegisterLegacyAminoCodec")
}

func (am AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----potv module-------------------RegisterInterfaces")
}

// RegisterGRPCGatewayRoutes registers gRPC Gateway routes for the potv module.
func (am AppModule) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("-----potv module-------------------RegisterGRPCGatewayRoutes")
}

// EndBlock is a method that will be run after transactions are processed in a block.
func (am AppModule) EndBlock(ctx context.Context) error {
	//TODO implement me
	//panic("implement me")
	return am.keeper.EndBlocker(ctx)
}

/*

// RegisterInvariants registers the potv module's invariants
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}*/

/*// Route returns the potv module's message routing key
func (am AppModule) Route() string {
	return types.RouterKey
}

// NewHandler returns a handler for "potv" type messages and an uninitialized Keeper.
func (am AppModule) NewHandler() sdk.PostHandler {
	return nil // Since we don't have any messages, we return nil
}

// QuerierRoute returns the potv module's querier route name
func (am AppModule) QuerierRoute() string {
	return types.QuerierRoute
}*/

// NewQuerierHandler returns the potv module's querier.
/*func (am AppModule) NewQuerierHandler() Querier {
	return nil // Placeholder for actual querier
}*/

// InitGenesis performs genesis initialization for the potv module. It sets initial state.
/*func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []*codec.GenesisValidator {
	// No need to handle genesis data for now
	return nil
}*/

/*// ExportGenesis returns the exported genesis state as raw JSON bytes for the potv module.
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return nil // No need to export genesis data for now
}

// GetTxCmd returns the root tx command for the potv module.
func (am AppModule) GetTxCmd() *cobra.Command {
	return nil // No tx commands for now
}

// GetQueryCmd returns the root query command for the potv module.
func (am AppModule) GetQueryCmd() *cobra.Command {
	// Placeholder for future query commands
	return nil
}*/

/*// RegisterRESTRoutes registers REST routes for the potv module.
func (am AppModule) RegisterRESTRoutes(clientCtx sdk.ClientContext, rtr *mux.Router) {
	// Placeholder for future REST routes
}



// GetQueryService returns the potv module's QueryService
func (am AppModule) GetQueryService() types.QueryService {
	return nil // Placeholder for future QueryService
}

// LegacyQuerierHandler returns a legacy querier handler.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil // Placeholder for legacy querier
}*/
