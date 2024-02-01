package colrepo

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConn struct {
	conn *mongo.Client
}

func (cfg *ColConfig) toMongoConnStr() string {
	var connectionString strings.Builder

	connectionType := os.Getenv("MONGO_CONNNECTION_TYPE")
	if connectionType == "SRV" {
		// Add the protocol
		connectionString.WriteString("mongodb+srv://")
	} else {
		connectionString.WriteString("mongodb://")
	}

	// Add username and password if provided
	if cfg.Username != "" && cfg.Password != "" {
		connectionString.WriteString(url.QueryEscape(cfg.Username))
		connectionString.WriteByte(':')
		connectionString.WriteString(url.QueryEscape(cfg.Password))
		connectionString.WriteByte('@')
	}

	// Add host and port
	connectionString.WriteString(cfg.Host)
	if cfg.Port != "" && connectionType != "SRV" {
		connectionString.WriteByte(':')
		connectionString.WriteString(cfg.Port)
	}

	authSource := os.Getenv("MONGO_AUTH_SOURCE")

	if authSource != "" {
		connectionString.WriteByte('/') 
		connectionString.WriteByte('?') // Add auth source
		connectionString.WriteString("authSource=")
		connectionString.WriteString(authSource)
	}

	log.Print(connectionString.String())
	return connectionString.String()
}

func NewMongoConn() (ColDBConn, error) {
	cfg := &ColConfig{}
	if err := cfg.buildConfigFromEnv(); err != nil {
		return nil, fmt.Errorf("could not load config, %w", err)
	}

	connStr := cfg.toMongoConnStr()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not connect to mongodb, %w", err)
	}

	conn := &MongoConn{
		conn: client,
	}

	return conn, nil
}

func (c MongoConn) GetConn() (*mongo.Client, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("no collection database connection found")
	}
	return c.conn, nil
}
