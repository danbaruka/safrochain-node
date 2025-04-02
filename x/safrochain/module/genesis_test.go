package safrochain_test

import (
	"testing"

	keepertest "safrochain/testutil/keeper"
	"safrochain/testutil/nullify"
	safrochain "safrochain/x/safrochain/module"
	"safrochain/x/safrochain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SafrochainKeeper(t)
	safrochain.InitGenesis(ctx, k, genesisState)
	got := safrochain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
