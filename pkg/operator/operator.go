package operator

import (
	"context"
	"errors"
)

// ErrNotFound returned if a subscription not found for the provided msisdn
var ErrNotFound = errors.New("subscription not found for the provided msisdn")

type Repository interface {
	Get(ctx context.Context, msisdn *string) (*string, error)
}
