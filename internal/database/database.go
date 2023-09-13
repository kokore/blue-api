package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	Host              string        `envconfig:"DB_HOST" default:"localhost:27017"`
	Name              string        `envconfig:"DB_NAME"`
	Username          string        `envconfig:"DB_USERNAME"`
	Password          string        `envconfig:"DB_PASSWORD"`
	Timeout           time.Duration `envconfig:"DB_TIMEOUT" default:"30s"`
	ConnectionTimeout time.Duration `envconfig:"DB_CONNECTION_TIMEOUT" default:"30s"`
}

func (dbConfig DBConfig) createClientOptions() *options.ClientOptions {
	hosts := []string{dbConfig.Host}
	opts := options.Client().
		SetHosts(hosts).
		SetConnectTimeout(dbConfig.ConnectionTimeout)

	if dbConfig.Username != "" || dbConfig.Password != "" {
		opts.SetAuth(options.Credential{
			AuthSource: dbConfig.Name,
			Username:   dbConfig.Username,
			Password:   dbConfig.Password,
		})
	}

	return opts
}

const (
	disconnectTimeout = 20 * time.Second
)

type Connection interface {
	Client() *mongo.Client
	Database() *mongo.Database

	IsConnected() bool
	Stop()
}

type connection struct {
	client    *mongo.Client
	dbConfig  *DBConfig
	connected bool
}

func (c *connection) init() error {
	opts := c.dbConfig.createClientOptions()

	ctx, cancelFn := context.WithTimeout(context.Background(), c.dbConfig.ConnectionTimeout)
	defer cancelFn()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	c.client = client
	c.connected = true
	return nil
}

func (c connection) Client() *mongo.Client {
	return c.client
}

func (c connection) IsConnected() bool {
	return c.connected
}

func (c connection) Database() *mongo.Database {
	return c.client.Database(c.dbConfig.Name)
}

func (c *connection) Stop() {
	if !c.connected {
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), disconnectTimeout)
	defer cancelFn()
	_ = c.client.Disconnect(ctx)
	c.connected = false
}

func NewConnection(dbConfig *DBConfig) (Connection, error) {
	c := &connection{
		dbConfig: dbConfig,
	}
	if err := c.init(); err != nil {
		return nil, err
	}
	return c, nil
}
