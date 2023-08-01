package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Implements StakingHooks interface
var _ types.BankHooks = BaseSendKeeper{}

// TrackBeforeSend executes the TrackBeforeSend hook if registered.
func (k BaseSendKeeper) TrackBeforeSend(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error {
	if k.hooks != nil {
		err := k.hooks.TrackBeforeSend(ctx, from, to, amount)
		if err != nil {
			return err
		}
	}
	return nil
}

// BlockBeforeSend executes the BlockBeforeSend hook if registered.
func (k BaseSendKeeper) BlockBeforeSend(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error {
	if k.hooks != nil {
		return k.hooks.BlockBeforeSend(ctx, from, to, amount)
	}
	return nil
}
