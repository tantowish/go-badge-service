package service

import (
	"badge-service/pkg/db"
	"badge-service/pkg/nsqpublisher"
	"badge-service/pkg/redis"
	proto "badge-service/protobuf"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BadgeService struct {
	proto.UnimplementedBadgeServiceServer
	RedisClient  redis.RedisClient
	DBClient     db.DatabaseClient
	NsqPublisher *nsqpublisher.Publisher
}

func (bs *BadgeService) GetBadge(ctx context.Context, req *proto.GetBadgeRequest) (*proto.Badge, error) {
	// Check redis for the same badge
	res, err := bs.RedisClient.Get(ctx, "badge-"+strconv.Itoa(int(req.Id)))
	if err == nil {
		var badge proto.Badge
		if err := json.Unmarshal([]byte(res), &badge); err != nil {
			return nil, fmt.Errorf("failed to unmarshal badge data from Redis: %v", err)
		}
		return &badge, nil
	}

	// Fetch from DB
	badge, err := bs.DBClient.GetBadge(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "badge not found")
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// Cache to redis
	if err := bs.RedisClient.SetBadge(ctx, "badge-"+strconv.Itoa(int(req.Id)), badge, 30*time.Minute); err != nil {
		log.Printf("Error caching badge in Redis: %v", err)
	}

	return badge, nil
}

func (bs *BadgeService) CreateBadge(ctx context.Context, req *proto.CreateBadgeRequest) (*proto.Badge, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "badge name cannot be empty")
	}

	badge, err := bs.DBClient.CreateBadge(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// Cache to redis
	if err := bs.RedisClient.SetBadge(ctx, "badge-"+strconv.Itoa(int(badge.Id)), badge, 30*time.Minute); err != nil {
		log.Printf("Error caching badge in Redis: %v", err)
	}

	return badge, nil
}

func (bs *BadgeService) UpdateBadge(ctx context.Context, req *proto.UpdateBadgeRequest) (*proto.Badge, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "badge name cannot be empty")
	}

	badge, err := bs.DBClient.UpdateBadge(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "badge not found")
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// Update the badge in redis
	if err := bs.RedisClient.SetBadge(ctx, "badge-"+strconv.Itoa(int(badge.Id)), badge, 30*time.Minute); err != nil {
		log.Printf("Error caching badge in Redis: %v", err)
	}

	return badge, nil
}

func (bs *BadgeService) DeleteBadge(ctx context.Context, req *proto.DeleteBadgeRequest) (*proto.DeleteBadgeResponse, error) {
	_, err := bs.DBClient.DeleteBadge(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &proto.DeleteBadgeResponse{Deleted: false}, status.Errorf(codes.NotFound, "badge not found")
		}
		return &proto.DeleteBadgeResponse{Deleted: false}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// Delete from redis
	if err := bs.RedisClient.Delete(ctx, "badge-"+strconv.Itoa(int(req.Id))); err != nil {
		log.Printf("Error deleting badge in Redis: %v", err)
	}

	return &proto.DeleteBadgeResponse{Deleted: true}, nil
}

func (bs *BadgeService) GetShopBadge(ctx context.Context, req *proto.GetShopBadgeRequest) (*proto.ShopBadge, error) {
	// Check redis for the same shopBadge
	res, err := bs.RedisClient.Get(ctx, "shop-"+strconv.FormatInt(req.ShopID, 10)+"-badge")
	if err == nil {
		var shopBadge proto.ShopBadge
		if err := json.Unmarshal([]byte(res), &shopBadge); err != nil {
			return nil, fmt.Errorf("failed to unmarshal shop badge data from Redis: %v", err)
		}
		return &shopBadge, nil
	}

	shopBadge, err := bs.DBClient.GetMerchantBadge(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	err = bs.RedisClient.SetMerchantBadge(ctx, "shop-"+strconv.FormatInt(req.ShopID, 10)+"-badge", shopBadge, 30*time.Minute)
	if err != nil {
		log.Printf("Error caching shop badge in Redis: %v", err)
	}

	return shopBadge, nil
}

func (bs *BadgeService) InvokeNsq(ctx context.Context, req *proto.GetShopBadgeRequest) (*proto.InvokeNsqResponse, error) {
	shopBadge, err := bs.GetShopBadge(ctx, req)
	if err != nil {
		return nil, err
	}

	shopIDMessage, err := json.Marshal(map[string]int64{
		"shop_id": req.ShopID,
	})
	if err != nil {
		return nil, err
	}

	err = bs.NsqPublisher.PublishMessage(string(shopIDMessage))
	if err != nil {
		return nil, err
	}

	return &proto.InvokeNsqResponse{
		ShopBadge: shopBadge,
		Message:   "success publishing message",
	}, nil
}