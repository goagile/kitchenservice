package ticket

import (
	"errors"
)

//
// Errrors
//
var (
	AcceptedFromPreparedIsNotValid = errors.New("AcceptedFromPreparedIsNotValid")
	AcceptedFromReadyToPickUpIsNotValid = errors.New("AcceptedFromReadyToPickUpIsNotValid")

	PrepareFromCreatedIsNotValid = errors.New("PrepareFromCreatedIsNotValid")
	PrepareFromReadyToPickUpIsNotValid = errors.New("PrepareFromReadyToPickUpIsNotValid")

	ReadyToPickUpFromCreatedIsNotValid = errors.New("ReadyToPickUpFromCreatedIsNotValid")
	ReadyToPickUpFromAcceptedIsNotValid = errors.New("ReadyToPickUpFromAcceptedIsNotValid")
		
	CancelFromAcceptedIsNotValid = errors.New("CancelFromAcceptedIsNotValid")
	CancelFromPreparedIsNotValid = errors.New("CancelFromPreparedIsNotValid")
)
