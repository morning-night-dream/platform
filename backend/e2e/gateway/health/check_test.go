//go:build e2e
// +build e2e

package health_test

import (
	"context"
	"testing"

	"github.com/morning-night-dream/platform/e2e/helper"
)

func TestGatewayE2EHealthCheck(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("ヘルスチェックが成功する", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClient(t, url)

		_, err := client.Client.V1Health(context.Background())
		if err != nil {
			t.Fatalf("failed to health check: %s", err)
		}
	})
}
