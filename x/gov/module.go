package gov

import (
	"encoding/json"

	cli2 "github.com/KiraCore/sekai/x/gov/client/cli"
	keeper2 "github.com/KiraCore/sekai/x/gov/keeper"
	customgovtypes "github.com/KiraCore/sekai/x/gov/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/grpc"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (b AppModuleBasic) RegisterCodec(amino *codec.LegacyAmino) {
}

func (b AppModuleBasic) Name() string {
	return customgovtypes.ModuleName
}

func (b AppModuleBasic) RegisterInterfaces(registry types2.InterfaceRegistry) {
}

func (b AppModuleBasic) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (b AppModuleBasic) ValidateGenesis(marshaler codec.JSONMarshaler, config client.TxEncodingConfig, message json.RawMessage) error {
	return nil
}

func (b AppModuleBasic) RegisterRESTRoutes(context client.Context, router *mux.Router) {
}

func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (b AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli2.GetCmdQueryPermissions()
}

// AppModule extends the cosmos SDK gov.
type AppModule struct {
	AppModuleBasic
	customGovKeeper keeper2.Keeper
}

func (am AppModule) RegisterInterfaces(registry types2.InterfaceRegistry) {
}

func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONMarshaler,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var genesisState customgovtypes.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	for _, actor := range genesisState.NetworkActors {
		am.customGovKeeper.SaveNetworkActor(ctx, *actor)
	}

	return nil
}

func (am AppModule) ExportGenesis(context sdk.Context, marshaler codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (am AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {}

func (am AppModule) QuerierRoute() string { return "" }

func (am AppModule) LegacyQuerierHandler(marshaler codec.JSONMarshaler) sdk.Querier {
	return nil
}

func (am AppModule) BeginBlock(context sdk.Context, block abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(ctx sdk.Context, block abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

func (am AppModule) Name() string {
	return customgovtypes.ModuleName
}

// Route returns the message routing key for the staking module.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// RegisterQueryService registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterQueryService(server grpc.Server) {
	querier := NewQuerier(am.customGovKeeper)
	customgovtypes.RegisterQueryServer(server, querier)
}

// NewAppModule returns a new Custom Staking module.
func NewAppModule(
	keeper keeper2.Keeper,
) AppModule {
	return AppModule{
		customGovKeeper: keeper,
	}
}