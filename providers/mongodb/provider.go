package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewConfigWith returns a new configuration with specific parameters
func NewConfigWith(connectionUrl string, dbName string, collection string) Config {
	return Config{
		ConnectionUrl: connectionUrl,
		Database:      dbName,
		Collection:    collection,
	}
}

// New returns a new configured mongodb provider
func New(cfg Config) (*Provider, error) {

	clientOptions := options.Client().
		ApplyURI(cfg.ConnectionUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, newErrMongoConnection(err)
	}

	p := &Provider{
		config: cfg,
		db:     client,
	}

	return p, nil
}
