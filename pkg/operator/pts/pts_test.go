package pts

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/rgynn/subscription-api/pkg/config"
)

func TestGet(t *testing.T) {

	cfg, err := config.NewFromEnv("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	repo := &Repository{
		timeout: cfg.ClientTimeout,
		url:     cfg.PTSURL,
		client: &http.Client{
			Timeout: cfg.ClientTimeout,
		},
	}

	ctx := context.Background()
	msisdn := os.Getenv("TEST_MSISDN_NUMBER")
	expectedName := os.Getenv("TEST_OPERATOR_NAME")

	result, err := repo.Get(ctx, &msisdn)
	if err != nil {
		t.Fatal(err)
	}

	if *result != expectedName {
		t.Fatalf("expected operator name to be: %s, got: %s", expectedName, *result)
	}
}
