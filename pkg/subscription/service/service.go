package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/rgynn/subscription-api/pkg/config"
	"github.com/rgynn/subscription-api/pkg/operator"
	"github.com/rgynn/subscription-api/pkg/operator/pts"
	"github.com/rgynn/subscription-api/pkg/subscription"
	"github.com/rgynn/subscription-api/pkg/subscription/repo/mem"
)

// Service for subscriptions
type Service struct {
	mem       subscription.Repository
	operators operator.Repository
}

func NewServiceFromConfig(cfg *config.Config) (*Service, error) {

	memrepo, err := mem.NewRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to inititalize in memory repository for subscriptions")
	}

	operatorsrepo, err := pts.NewRepository(cfg.ClientTimeout, cfg.PTSURL)
	if err != nil {
		return nil, fmt.Errorf("failed to inititalize in operator repository for subscriptions")
	}

	return &Service{
		mem:       memrepo,
		operators: operatorsrepo,
	}, nil
}

func (svc *Service) List(ctx context.Context) ([]*subscription.Model, error) {

	result, err := svc.mem.List(ctx)
	if err != nil {
		return nil, err
	}

	for i, sub := range result {
		op, err := svc.operators.Get(ctx, sub.MSISDN)
		if err != nil {
			return nil, fmt.Errorf("failed to get operator info for msisdn: %s, error: %s", *sub.MSISDN, err)
		}
		result[i].Operator = op
	}

	return result, nil
}

func (svc *Service) Get(ctx context.Context, msisdn *string) (*subscription.Model, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	result, err := svc.mem.Get(ctx, msisdn)
	if err != nil {
		return nil, err
	}

	op, err := svc.operators.Get(ctx, msisdn)
	if err != nil {
		return nil, fmt.Errorf("failed to get operator info for msisdn: %s, error: %s", *msisdn, err)
	}

	result.Operator = op

	return result, nil
}

func (svc *Service) Create(ctx context.Context, m *subscription.Model) (*subscription.Model, error) {

	if m == nil {
		return nil, errors.New("no subscription provided")
	}

	if err := m.ValidForSave(); err != nil {
		return nil, fmt.Errorf("%s: %w", err.Error(), subscription.ErrNotValid)
	}

	op, err := svc.operators.Get(ctx, m.MSISDN)
	if err != nil {
		return nil, fmt.Errorf("failed to get operator info for msisdn: %s, error: %s", *m.MSISDN, err)
	}

	m.Operator = op

	if err := m.UpdateStatus(nil); err != nil {
		return nil, err
	}

	result, err := svc.mem.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *Service) Update(ctx context.Context, m *subscription.Model) (*subscription.Model, error) {

	if m == nil {
		return nil, errors.New("no subscription provided")
	}

	if err := m.ValidForUpdate(); err != nil {
		return nil, fmt.Errorf("%s: %w", err.Error(), subscription.ErrNotValid)
	}

	sub, err := svc.mem.Update(ctx, m)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (svc *Service) TogglePaused(ctx context.Context, msisdn *string) (*subscription.Model, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	sub, err := svc.mem.TogglePaused(ctx, msisdn)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (svc *Service) Cancel(ctx context.Context, msisdn *string) (*subscription.Model, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	sub, err := svc.mem.Cancel(ctx, msisdn)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
