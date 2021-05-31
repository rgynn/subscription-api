package subscription

import (
	"context"
	"errors"
	"time"
)

// ErrAlreadyExists returned if a subscription already exists for the provided msisdn
var ErrAlreadyExists = errors.New("subscription already exists for the provided msisdn")

// ErrNotFound returned if a subscription not found for the provided msisdn
var ErrNotFound = errors.New("subscription not found for the provided msisdn")

// ErrNotValid returned if provided subscription not valid
var ErrNotValid = errors.New("provided subscription not valid")

var (
	// StatusPending for subscription
	StatusPending = "pending"
	// StatusActivated for subscription
	StatusActivated = "activated"
	// StatusPaused for subscription
	StatusPaused = "paused"
	// StatusCancelled for subscription
	StatusCancelled = "cancelled"
)

// Repository interface for subscription
type Repository interface {
	List(ctx context.Context) ([]*Model, error)
	Get(ctx context.Context, msisdn *string) (*Model, error)
	Create(ctx context.Context, m *Model) (*Model, error)
	Update(ctx context.Context, m *Model) (*Model, error)
	TogglePaused(ctx context.Context, msisdn *string) (*Model, error)
	Cancel(ctx context.Context, msisdn *string) (*Model, error)
}

// Model of a subscription
type Model struct {
	MSISDN     *string    `json:"msisdn"`
	ActivateAt *time.Time `json:"activate_at"`
	Type       *string    `json:"type"`
	Status     *string    `json:"status"`
	Operator   *string    `json:"operator,omitempty"`
}

// ValidForSave?
func (m *Model) ValidForSave() error {

	if m == nil {
		return errors.New("cannot validate a nil subscription")
	}

	if m.MSISDN == nil {
		return errors.New("no msisdn provided")
	}

	if m.ActivateAt == nil {
		return errors.New("no activate_at provided")
	}

	if m.Type == nil {
		return errors.New("no type provided")
	} else {
		switch *m.Type {
		case "PBX", "CELL":
			break
		default:
			return errors.New("type needs to be either PBX or CELL")
		}
	}

	if m.Status != nil {
		return errors.New("cannot provided status when creating new subscription")
	}

	if m.Operator != nil {
		return errors.New("field operator is read only")
	}

	return nil
}

// ValidForSave?
func (m *Model) ValidForUpdate() error {

	if m == nil {
		return errors.New("cannot validate a nil subscription")
	}

	if m.MSISDN == nil {
		return errors.New("no msisdn provided")
	}

	if m.ActivateAt == nil {
		return errors.New("no activate_at provided")
	}

	if m.Type == nil {
		return errors.New("no type provided")
	} else {
		switch *m.Type {
		case "PBX", "CELL":
			break
		default:
			return errors.New("type needs to be either PBX or CELL")
		}
	}

	if m.Status != nil {
		return errors.New("cannot provided status when updating subscription")
	}

	if m.Operator != nil {
		return errors.New("field operator is read only")
	}

	return nil
}

func (m *Model) IsActive() bool {
	return m != nil && m.Status != nil && *m.Status == "activated"
}

func (m *Model) UpdateStatus(status *string) error {

	if m == nil {
		return errors.New("cannot update status of nil subscription")
	}

	if m.ActivateAt == nil {
		return errors.New("no activate_at provided")
	}

	if status != nil {
		m.Status = status
		return nil
	}

	now := time.Now().UTC()

	if m.ActivateAt.After(now) {
		m.Status = &StatusPending
		return nil
	}

	if m.ActivateAt.Before(now) {
		m.Status = &StatusActivated
	}

	return nil
}
