package health_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/e2e/helper"
	healthv1 "github.com/morning-night-dream/platform/pkg/connect/proto/health/v1"
)

func TestE2EHealthCheck(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("ヘルスチェックが成功する", func(t *testing.T) {
		t.Parallel()

		client := helper.NewConnectClient(t, http.DefaultClient, url)

		req := &healthv1.CheckRequest{}

		res, err := client.Health.Check(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Errorf("faile to health check: %s", err)
		}

		if !reflect.DeepEqual(res.StatusCode, http.StatusOK) {
			t.Errorf("Articles actual = %v, want %v", res.StatusCode, http.StatusOK)
		}
	})
}
