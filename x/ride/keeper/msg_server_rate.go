package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// Average is serialized to string due to amino codec not being able to handle float serialization.
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
			Rating: fmt.Sprintf("%f", 0.0), // New member starts out with a rating of 0 to be built up.
		}
	}
	currentRating, err := strconv.ParseFloat(rating.Rating, 64)
	if err != nil {
		return &types.MsgRateResponse{}, err
	}
	newRating, err := strconv.ParseFloat(msg.Rating, 64)
	if err != nil {
		return &types.MsgRateResponse{}, err
	}

	rating.Rating = fmt.Sprintf("%f", types.ComputePseudoAverage(currentRating, newRating))
	k.Keeper.SetRatingStruct(ctx, rating)
	return &types.MsgRateResponse{Success: true}, nil
}

func ValidateRate(msg *types.MsgRate, storedRide types.StoredRide) error {

	validAddressesToRateDriver := msg.Creator == storedRide.PassengerAddress && msg.Ratee == storedRide.DriverAddress
	validAddressesToRatePassenger := msg.Creator == storedRide.DriverAddress && msg.Ratee == storedRide.PassengerAddress

	if !validAddressesToRateDriver && !validAddressesToRatePassenger {
		return errors.Wrapf(types.ErrInvalidAddressCombo,
			"passenger for ride is stored as %s and driver is %s. Msg creator is %s and ratee is %s",
			storedRide.PassengerAddress, storedRide.DriverAddress, msg.Creator, msg.Ratee)
	}
	ratingValue, err := strconv.ParseFloat(msg.Rating, 64)
	if err != nil {
		return err
	}
	if ratingValue > 10 || ratingValue < 0 {
		return types.ErrInvalidRatingValue
	}
	return nil
}
