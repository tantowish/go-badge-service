package db

import (
	proto "badge-service/protobuf"
	"context"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DatabaseClient interface {
	GetBadge(ctx context.Context, badgeRequest *proto.GetBadgeRequest) (*proto.Badge, error)
	CreateBadge(ctx context.Context, badgeRequest *proto.CreateBadgeRequest) (*proto.Badge, error)
	UpdateBadge(ctx context.Context, badgeRequest *proto.UpdateBadgeRequest) (*proto.Badge, error)
	DeleteBadge(ctx context.Context, badgeRequest *proto.DeleteBadgeRequest) (bool, error)
	GetMerchantBadge(ctx context.Context, shopRequest *proto.GetShopBadgeRequest) (*proto.ShopBadge, error)
}

type databaseClient struct {
	db *sql.DB
}

func NewDatabaseClient(migrationPath string, dsn string) (DatabaseClient, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging to PostgreSQL: %w", err)
	}
	fmt.Println("Database connected successfully!")

	m, err := migrate.New(
		"file://"+migrationPath,
		dsn)
	if err != nil {
		return nil, fmt.Errorf("error initializing migration: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("error applying migrations: %w", err)
	}
	fmt.Println("Migrations applied successfully!")

	return &databaseClient{db: db}, nil
}

// GetBadge implements DatabaseClient.
func (d *databaseClient) GetBadge(ctx context.Context, badgeRequest *proto.GetBadgeRequest) (*proto.Badge, error) {
	badge := proto.Badge{}
	err := d.db.QueryRowContext(ctx, "SELECT id, name FROM badges WHERE id = $1", badgeRequest.Id).Scan(&badge.Id, &badge.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("badge with ID %d not found: %w", badgeRequest.Id, err)
		}
		return nil, fmt.Errorf("error querying badge: %w", err)
	}

	return &badge, nil
}

// CreateBadge implements DatabaseClient.
func (d *databaseClient) CreateBadge(ctx context.Context, badgeRequest *proto.CreateBadgeRequest) (*proto.Badge, error) {
	var badgeID int
	err := d.db.QueryRowContext(ctx, "INSERT INTO badges(name) VALUES ($1) RETURNING id", badgeRequest.Name).Scan(&badgeID)
	if err != nil {
		return nil, fmt.Errorf("error creating badge: %w", err)
	}

	return &proto.Badge{
		Id:   int64(badgeID),
		Name: badgeRequest.Name,
	}, nil
}

// UpdateBadge implements DatabaseClient.
func (d *databaseClient) UpdateBadge(ctx context.Context, badgeRequest *proto.UpdateBadgeRequest) (*proto.Badge, error) {
	existingBadge, err := d.GetBadge(ctx, &proto.GetBadgeRequest{Id: badgeRequest.Id})
	if err != nil {
		return nil, err
	}

	_, err = d.db.ExecContext(ctx, "UPDATE badges SET name = $1 WHERE id = $2", badgeRequest.Name, badgeRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("error updating badge: %w", err)
	}

	return &proto.Badge{
		Id:   existingBadge.Id,
		Name: badgeRequest.Name,
	}, nil
}

// DeleteBadge implements DatabaseClient.
func (d *databaseClient) DeleteBadge(ctx context.Context, badgeRequest *proto.DeleteBadgeRequest) (bool, error) {
	_, err := d.GetBadge(ctx, &proto.GetBadgeRequest{Id: badgeRequest.Id})
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}

	result, err := d.db.ExecContext(ctx, "DELETE FROM badges WHERE id = $1", badgeRequest.Id)
	if err != nil {
		return false, fmt.Errorf("error deleting badge: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error deleting badge: %w", err)
	}

	if rowsAffected == 0 {
		return false, err
	}

	return true, nil
}

// GetMerchantBadge implements DatabaseClient.
func (d *databaseClient) GetMerchantBadge(ctx context.Context, shopRequest *proto.GetShopBadgeRequest) (*proto.ShopBadge, error) {
	if shopRequest == nil || shopRequest.ShopID == 0 {
		return nil, fmt.Errorf("invalid shop request or missing shop ID")
	}

	var shop proto.ShopBadge
	if shop.Badge == nil {
		shop.Badge = &proto.Badge{}
	}

	query := `
		SELECT m.id, m.name, m.badge_id, b.id, b.name
		FROM shops m
		JOIN badges b ON m.badge_id = b.id
		WHERE m.id = $1
	`

	err := d.db.QueryRowContext(ctx, query, shopRequest.ShopID).Scan(&shop.Id, &shop.Name, &shop.BadgeID, &shop.Badge.Id, &shop.Badge.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("shop with ID %d not found: %w", shopRequest.ShopID, err)
		}
		return nil, fmt.Errorf("error querying shop badge: %w", err)
	}

	return &shop, nil
}
