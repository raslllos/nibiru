package binding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	oraclekeeper "github.com/NibiruChain/nibiru/x/oracle/keeper"
	perpv2keeper "github.com/NibiruChain/nibiru/x/perp/v2/keeper"
	"github.com/NibiruChain/nibiru/x/sudo"
)

func RegisterWasmOptions(
	perpv2 perpv2keeper.Keeper,
	sudoKeeper sudo.Keeper,
	oracleKeeper oraclekeeper.Keeper,
) []wasm.Option {
	wasmQueryPlugin := NewQueryPlugin(perpv2)
	wasmQueryOption := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})

	wasmExecuteOption := wasmkeeper.WithMessageHandlerDecorator(
		CustomExecuteMsgHandler(perpv2, sudoKeeper, oracleKeeper),
	)

	return []wasm.Option{wasmQueryOption, wasmExecuteOption}
}
