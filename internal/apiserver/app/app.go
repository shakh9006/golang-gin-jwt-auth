package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/shakh9006/golang-gin-jwt-auth/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

var (
	ctx context.Context
)

func InitDB(cfg config.Config) (*mongo.Client, error) {
	mongoConn := options.Client().ApplyURI(cfg.DBUri)
	mongoClient, err := mongo.Connect(ctx, mongoConn)

	if err != nil {
		return nil, err
	}
	fmt.Println(123)
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	fmt.Println("MongoDB connected successfully...")

	return mongoClient, nil
}

func InitRedis(cfg config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisUri,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	err := redisClient.Set(ctx, "test", "Welcome to Goland with MongoDB and Redis", 0).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("Redis connected successfully...")
	return redisClient, nil
}

func Run() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables")
	}

	ctx = context.TODO()

	mongoClient, err := InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to DB: %s", err)
	}

	defer mongoClient.Disconnect(ctx)

	redisClient, err := InitRedis(cfg)
	if err != nil {
		log.Fatalf("Could not connect to Redis: %s", err)
	}

	value, err := redisClient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		log.Fatalf("Error on parsing data from redis: %s", err)
	}

	server := gin.Default()
	router := server.Group("/api")
	router.GET("/healthchecker", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})
	log.Fatal(server.Run(fmt.Sprintf(":%s", cfg.Port)))
}
