package integration

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type (
	PostgresContainer struct {
		testcontainers.Container
		Config PostgresContainerConfig
	}

	PostgresContainerOption func(c *PostgresContainerConfig)

	PostgresContainerConfig struct {
		ImageTag   string
		User       string
		Password   string
		Host       string
		MappedPort string
		Database   string
		SSLMode    string
	}
)

func (c PostgresContainer) GetConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		c.Config.User,
		c.Config.Password,
		c.Config.Host,
		c.Config.MappedPort,
		c.Config.Database,
		c.Config.SSLMode,
	)
}

func NewPostgresContainer(ctx context.Context, opts ...PostgresContainerOption) (*PostgresContainer, error) {
	const (
		psqlImage = "postgres"
		psqlPort  = "5432"
	)

	config := PostgresContainerConfig{
		ImageTag: "15.2-alpine",
		User:     "test_user",
		Password: "test_password",
		Database: "test_db",
	}

	for _, opt := range opts {
		opt(&config)
	}

	containerPort := psqlPort + "/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Env: map[string]string{
				"POSTGRES_USER":     config.User,
				"POSTGRES_PASSWORD": config.Password,
				"POSTGRES_DB":       config.Database,
			},
			ExposedPorts: []string{
				containerPort,
			},
			Image: fmt.Sprintf("%s:%s", psqlImage, config.ImageTag),
			WaitingFor: wait.ForExec([]string{"pg_isready", "-d", config.Database, "-U", config.User}).
				WithPollInterval(1 * time.Second).
				WithExitCodeMatcher(func(exitCode int) bool {
					return exitCode == 0
				}),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "generic container")
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "container host")
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(containerPort))
	if err != nil {
		return nil, errors.WithMessagef(err, "getting mapped port for %s: %v", containerPort, err)
	}

	config.MappedPort = mappedPort.Port()
	config.Host = host

	return &PostgresContainer{
		Container: container,
		Config:    config,
	}, nil
}
