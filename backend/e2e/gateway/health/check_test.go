//go:build e2e
// +build e2e

package health_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/morning-night-dream/platform/e2e/helper"
)

func TestGatewayE2EHealthCheck(t *testing.T) {
	t.Parallel()

	url := helper.GetGatewayEndpoint(t)

	t.Run("ヘルスチェックが成功する", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClient(t, url)

		res, err := client.Client.V1Health(context.Background())
		if err != nil {
			t.Fatalf("failed to health check: %s", err)
		}

		if !reflect.DeepEqual(res.Status, http.StatusOK) {
			t.Errorf("Articles actual = %v, want %v", res.Status, http.StatusOK)
		}
	})
}
