package mem

import (
	"context"
	"errors"
	"sync"

	"github.com/rgynn/subscription-api/pkg/subscription"
)

// Repository for in memory subscriptions
type Repository struct {
	subscriptions map[string]*subscription.Model
	sync.Mutex
}

func NewRepository() (subscription.Repository, error) {
	return &Repository{
		subscriptions: map[string]*subscription.Model{},
	}, nil
}

func (repo *Repository) List(ctx context.Context) ([]*subscription.Model, error) {

	repo.Lock()
	defer repo.Unlock()

	result := make([]*subscription.Model, len(repo.subscriptions))

	for _, sub := range repo.subscriptions {
		result = append(result, sub)
	}

	return result, nil
}

func (repo *Repository) Get(ctx context.Context, msisdn *string) (*subscription.Model, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	repo.Lock()
	defer repo.Unlock()

	sub, ok := repo.subscriptions[*msisdn]
	if !ok {
		return nil, subscription.ErrNotFound
	}

	return sub, nil
}

func (repo *Repository) Create(ctx context.Context, m *subscription.Model) (*subscription.Model, error) {

	if m == nil {
		return nil, errors.New("no m *subscription.Model provided")
	}

	repo.Lock()
	defer repo.Unlock()

	if sub, ok := repo.subscriptions[*m.MSISDN]; ok && sub.IsActive() {
		return nil, subscription.ErrAlreadyExists
	}

	repo.subscriptions[*m.MSISDN] = m

	return m, nil
}

func (repo *Repository) Update(ctx context.Context, m *subscription.Model) (*subscription.Model, error) {

	if m == nil {
		return nil, errors.New("no subscription provided")
	}

	repo.Lock()
	defer repo.Unlock()

	sub, ok := repo.subscriptions[*m.MSISDN]
	if !ok {
		return nil, subscription.ErrNotFound
	}

	// check if trying to change activate_at and status isn't pending
	if !sub.ActivateAt.Equal(*m.ActivateAt) && sub.Status != &subscription.StatusPending {
		return nil, errors.New("subscription needs to be pending to update activate_at")
	}

	sub.ActivateAt = m.ActivateAt
	sub.Type = m.Type

	if err := sub.UpdateStatus(nil); err != nil {
		return nil, err
	}

	repo.subscriptions[*m.MSISDN] = sub

	return sub, nil
}

func (repo *Repository) TogglePaused(ctx context.Context, msisdn *string) (*subscription.Model, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	repo.Lock()
	defer repo.Unlock()

	sub, ok := repo.subscriptions[*msisdn]
	if !ok {
		return nil, subscription.ErrNotFound
	}

	if sub.Status == &subscription.StatusPaused {
		if err := sub.UpdateStatus(nil); err != nil {
			return nil, err
		}
	} else {
		if err := sub.UpdateStatus(&subscription.StatusPaused); err != nil {
			return nil, err
		}
	}

	repo.subscriptions[*msisdn] = sub

	return sub, nil
}

func (repo *Repository) Cancel(ctx context.Context, msisdn *string) (*subscription.Model, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	repo.Lock()
	defer repo.Unlock()

	sub, ok := repo.subscriptions[*msisdn]
	if !ok {
		return nil, subscription.ErrNotFound
	}

	if err := sub.UpdateStatus(&subscription.StatusCancelled); err != nil {
		return nil, err
	}

	repo.subscriptions[*msisdn] = sub

	return sub, nil
}
