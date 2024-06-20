package keeper

import (
	"github.com/CudoVentures/cudos-node/x/mail/types"
)

var _ types.QueryServer = Keeper{}
