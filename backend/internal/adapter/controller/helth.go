package controller

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	healthv1 "github.com/morning-night-dream/platform/pkg/proto/health/v1"
)

func (c *Controller) V1Health(w http.ResponseWriter, r *http.Request) {
	req := &healthv1.CheckRequest{}
	res, err := c.client.Health.Check(r.Context(), connect.NewRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res.Msg.String()))
}
