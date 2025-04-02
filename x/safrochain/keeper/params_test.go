package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "safrochain/testutil/keeper"
	"safrochain/x/safrochain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SafrochainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
