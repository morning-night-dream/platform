package test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DBDocker struct {
	pid string
	DSN string
	*client.Client
}

func NewDBDocker(t *testing.T) *DBDocker {
	t.Helper()

	cli, err := client.NewClientWithOpts()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	if _, err = cli.ImagePull(ctx, os.Getenv("POSTGRES_IMAGE"), types.ImagePullOptions{}); err != nil {
		t.Fatal(err)
	}

	cfg := &container.Config{
		Env: []string{
			`TZ=UTC`,
			`LANG=ja_JP.UTF-8`,
			`POSTGRES_DB=postgres`,
			`POSTGRES_USER=postgres`,
			`POSTGRES_PASSWORD=postgres`,
			`POSTGRES_INITDB_ARGS="--encoding=UTF-8"`,
			`POSTGRES_HOST_AUTH_METHOD=trust`,
		},
		Image:        os.Getenv("POSTGRES_IMAGE"),
		ExposedPorts: nat.PortSet{"5432/tcp": {}},
	}

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	addr := l.Addr().String()
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		t.Fatal(err)
	}

	hcfg := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, cfg, hcfg, nil, &v1.Platform{}, port)
	if err != nil {
		t.Fatal(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		t.Fatal()
	}

	time.Sleep(10 * time.Second)

	return &DBDocker{
		DSN:    fmt.Sprintf("postgres://postgres:postgres@localhost:%s/postgres?sslmode=disable", port),
		pid:    resp.ID,
		Client: cli,
	}
}

func (d *DBDocker) TearDown(t *testing.T) {
	t.Helper()

	ctx := context.Background()

	timeout := time.Second
	if err := d.ContainerStop(ctx, d.pid, &timeout); err != nil {
		t.Error(err)
	}

	if err := d.ContainerRemove(ctx, d.pid, types.ContainerRemoveOptions{}); err != nil {
		t.Error(err)
	}
}
