package main

import (
	"badge-service/pkg/db"
	"badge-service/pkg/nsqpublisher"
	"badge-service/pkg/redis"
	"badge-service/protobuf"
	"badge-service/service"
	"log"
	"net"
	"os"
	"strconv"

	zapGrpc "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error creating Redis client: %v", err)
	}
	redisClient, err := redis.NewRedisClient(redisAddr, redisPassword, redisDB)
	if err != nil {
		log.Fatalf("Error creating Redis client: %v", err)
	}

	dbDSN := os.Getenv("DB_DSN")
	migrationPath := os.Getenv("MIGRATION_PATH")
	dbClient, err := db.NewDatabaseClient(migrationPath,dbDSN)
	if err != nil {
		log.Fatalf("Error creating PostgreSQL client: %v", err)
	}

	nsqAddr := "127.0.0.1:4150"
	topic := "shop_badge_change_data"
	nsqPub, err := nsqpublisher.NewPublisher(nsqAddr, topic)
	if err != nil {
		log.Panic("Failed to initialize NSQ publisher: ", err)
	}
	defer nsqPub.StopPublisher() 

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			zapGrpc.UnaryServerInterceptor(logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
		),
	)

	reflection.Register(server)

	proto.RegisterBadgeServiceServer(server, &service.BadgeService{
		RedisClient: redisClient,
		DBClient: dbClient,
		NsqPublisher: nsqPub,
	})

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server listening on port: %s", port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

