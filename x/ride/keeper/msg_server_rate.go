package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

func (k msgServer) Rate(goCtx context.Context, msg *types.MsgRate) (*types.MsgRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.RideId)
	if !found {
		return &types.MsgRateResponse{}, errors.Wrapf(types.ErrRideNotFound,
			"ride not found from input parameter %s", msg.RideId)
	}
	err := ValidateRate(msg, storedRide)
	if err != nil {
		return &types.MsgRateResponse{}, err
	}

	rating, found := k.Keeper.GetRatingStruct(ctx, msg.Ratee)
	if !found {
		rating = types.RatingStruct{
			Index:  msg.Ratee,
			Rating: 0, // New member starts out with a rating of 0 to be built up.
		}
	}

	rating.Rating = types.ComputePseudoAverage(rating.Rating, msg.Rating)
	k.Keeper.SetRatingStruct(ctx, rating)

	return &types.MsgRateResponse{Success: true}, nil
}

func ValidateRate(msg *types.MsgRate, storedRide types.StoredRide) error {

	validAddressesToRateDriver := msg.Creator == storedRide.Passenger && msg.Ratee == storedRide.Driver
	validAddressesToRatePassenger := msg.Creator == storedRide.Driver && msg.Ratee == storedRide.Passenger

	if !validAddressesToRateDriver && !validAddressesToRatePassenger {
		return errors.Wrapf(types.ErrInvalidAddressCombo,
			"passenger for ride is stored as %s and driver is %s", storedRide.Passenger, storedRide.Driver)
	}
	if msg.Rating > 10 { // Implicit assumption: uint cant be negative.
		return types.ErrInvalidRatingValue
	}
	return nil
}
