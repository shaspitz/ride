package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// RequestRide message handler obtains the next ride counter from app chain state,
// creates and stores a new ride object using this counter, increments that counter
// for future rides, and returns a response.
func (k msgServer) RequestRide(goCtx context.Context, msg *types.MsgRequestRide) (*types.MsgRequestRideResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.MutualStake < msg.DistanceTip+msg.HourlyPay {
		return nil, errors.Wrapf(types.ErrRideParameters, "mutual stake is below distance tip + 1 hour pay")
	}

	ctx.GasMeter().ConsumeGas(types.RequestRideGas, "Request Ride")

	nextRide, found := k.Keeper.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}

	newRideIndex := strconv.FormatUint(nextRide.IdValue, 10)

	storedRide := types.StoredRide{
		Index: newRideIndex,
		// NOTE: Start location would be validated/stored if a proof-of-location system were implemented.
		Destination: msg.Destination,
		// NOTE: Driver is set in the accept Tx.
		PassengerAddress: msg.Creator,
		MutualStake:      msg.MutualStake,
		PayPerHour:       msg.HourlyPay,
		DistanceTip:      msg.DistanceTip,
		// To be updated below.
		BeforeId: types.NoFifoIdKey,
		AfterId:  types.NoFifoIdKey,
	}

	// Validate assigned passenger address can be parsed.
	_, err := storedRide.GetPassengerSdkAddress()
	if err != nil {
		return nil, err
	}

	err = k.Keeper.CollectPassengerStake(ctx, &storedRide)
	if err != nil {
		return nil, err
	}

	// Ride is stored this block, send it to the FIFO tail.
	k.Keeper.SendToFifoTail(ctx, &storedRide, &nextRide)

	// Store ride via keeper.
	k.Keeper.SetStoredRide(ctx, storedRide)

	// Increment counter for next stored ride.
	nextRide.IdValue++
	k.Keeper.SetNextRide(ctx, nextRide)

	// Emit appropriate event with starting location of ride.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.RequestRideEventKey),
			sdk.NewAttribute(types.RequestRideEventPassenger, storedRide.PassengerAddress),
			sdk.NewAttribute(types.RequestRideEventIndex, storedRide.Index),
			sdk.NewAttribute(types.RequestRideStartLocation, msg.StartLocation),
		),
	)

	return &types.MsgRequestRideResponse{
		IdValue: newRideIndex,
	}, nil
}

// Handles an acceptance transaction, for a driver accepting a requested ride.
func (k msgServer) Accept(goCtx context.Context, msg *types.MsgAccept) (*types.MsgAcceptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.IdValue)
	if !found {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrRideNotFound, "ride not found %s", msg.IdValue)
	}

	nextRide, found := k.Keeper.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}

	if storedRide.HasAssignedDriver() {
		return &types.MsgAcceptResponse{Success: false}, errors.Wrapf(types.ErrAssignedDriver, "driver has already been assigned for ride with Id: %s", msg.IdValue)
	}

	if msg.Creator == storedRide.PassengerAddress {
		return &types.MsgAcceptResponse{Success: false}, errors.Wrapf(types.ErrCannotDriveYourself,
			"%s is the passenger of this ride, and cannot be the driver", storedRide.PassengerAddress)
	}

	ctx.GasMeter().ConsumeGas(types.AcceptRideGas, "Accept Ride")

	// Assign driver to ride.
	storedRide.DriverAddress = msg.Creator

	// Validate assigned driver address, could be in unit test to save gas.
	_, err := storedRide.GetDriverSdkAddress()
	if err != nil {
		return &types.MsgAcceptResponse{Success: false},
			errors.Wrapf(types.ErrInvalidDriver, "invalid driver address %s", msg.IdValue)
	}

	// Store acceptance time in default format.
	storedRide.AcceptanceTime = types.TimeToString(ctx.BlockTime())

	err = k.Keeper.CollectDriverStake(ctx, &storedRide)
	if err != nil {
		return nil, err
	}

	// Ride is mutated this block, send it to the FIFO tail then update store accordingly.
	k.Keeper.SendToFifoTail(ctx, &storedRide, &nextRide)
	k.Keeper.SetStoredRide(ctx, storedRide)
	k.Keeper.SetNextRide(ctx, nextRide)

	// Emit appropriate event for ride acceptance.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AcceptRideEventKey),
			sdk.NewAttribute(types.AcceptRideEventDriver, storedRide.DriverAddress),

			sdk.NewAttribute(types.AcceptRideEventIdValue, storedRide.Index),
		),
	)

	return &types.MsgAcceptResponse{Success: true}, nil
}

// Handles a finish transaction.
func (k msgServer) Finish(goCtx context.Context, msg *types.MsgFinish) (*types.MsgFinishResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedRide, found := k.Keeper.GetStoredRide(ctx, msg.IdValue)
	if !found {
		return &types.MsgFinishResponse{Success: false},
			errors.Wrapf(types.ErrRideNotFound, "ride not found with Id: %s", msg.IdValue)
	}

	nextRide, found := k.Keeper.GetNextRide(ctx)
	if !found {
		panic("NextRide not found")
	}

	if msg.Creator != storedRide.PassengerAddress && msg.Creator != storedRide.DriverAddress {
		return &types.MsgFinishResponse{Success: false}, errors.Wrapf(types.ErrIrrelevantRide, "%s is not associated with ride %s",
			msg.Creator, msg.IdValue)
	}

	if storedRide.IsFinished() {
		return &types.MsgFinishResponse{Success: false}, errors.Wrapf(types.ErrAlreadyFinishedRide, "ride %s already finished", msg.IdValue)
	}

	ctx.GasMeter().ConsumeGas(types.FinishRideGas, "Finish Ride")

	// No driver accepted yet in this (valid) case. Return funds and erase game.
	if !storedRide.HasAssignedDriver() {
		k.Keeper.MustRefundStakes(ctx, &storedRide)
		k.Keeper.RemoveFromFifo(ctx, &storedRide, &nextRide)
		k.Keeper.RemoveStoredRide(ctx, msg.IdValue)
		k.Keeper.SetNextRide(ctx, nextRide)
		return &types.MsgFinishResponse{Success: true}, nil
	}

	storedRide.FinishTime = types.TimeToString(ctx.BlockTime())
	storedRide.FinishLocation = msg.Location

	// Ride is mutated this block, send it to the FIFO tail then update store accordingly.
	k.Keeper.SendToFifoTail(ctx, &storedRide, &nextRide)
	k.Keeper.SetStoredRide(ctx, storedRide)
	k.Keeper.SetNextRide(ctx, nextRide)

	// Emit appropriate event for ride acceptance.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.FinishRideEventKey),
		),
	)

	return &types.MsgFinishResponse{Success: true}, nil
}

// Handles a rate transaction.
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
