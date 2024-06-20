package types_test

import (
	"testing"
	"time"

	"github.com/CudoVentures/cudos-node/testutil/sample"
	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id:          1,
						Sender:      sample.AccAddress(),
						Receiver:    sample.AccAddress(),
						Body:        "sample 1",
						ForwardedId: 0,
						Ts:          uint64(time.Now().Unix()),
					},
					{
						Id:          2,
						Sender:      sample.AccAddress(),
						Receiver:    sample.AccAddress(),
						Body:        "sample 2",
						ForwardedId: 0,
						Ts:          uint64(time.Now().Unix()),
					},
				},
				MessageCount: 2,
				Params:       types.NewParams(32),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated message",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id: 1,
					},
					{
						Id: 1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid message count",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id: 2,
					},
				},
				MessageCount: 1,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
