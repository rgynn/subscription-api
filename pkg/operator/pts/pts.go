package pts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rgynn/subscription-api/pkg/operator"
)

// PTSResponse from api
type PTSResponse struct {
	D struct {
		Type   string `json:"__type"`
		Name   string `json:"Name"`
		Number string `json:"Number"`
	} `json:"d"`
}

type Repository struct {
	timeout time.Duration
	url     string
	client  *http.Client
}

// NewRepsitory for operator using pts api
func NewRepository(timeout time.Duration, url string) (operator.Repository, error) {
	return &Repository{
		timeout: timeout,
		url:     url,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (repo *Repository) Get(ctx context.Context, msisdn *string) (*string, error) {

	if msisdn == nil {
		return nil, errors.New("no msisdn provided")
	}

	url := fmt.Sprintf("%s?number=%s", repo.url, *msisdn)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to pts: %s", err)
	}

	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := repo.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call pts: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("expected status code 200 OK from PTS API, got: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response PTSResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body from pts response: %w", err)
	}

	if response.D.Name == "Operatör saknas" {
		return nil, operator.ErrNotFound
	}

	return &response.D.Name, nil
}
