package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/lvmnpkfvmk/avito-tech/config"
	"github.com/lvmnpkfvmk/avito-tech/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("User not found")
)

type SegmentRepository struct {
	db    *gorm.DB
}

func NewSegmentRepository(ctx context.Context, cfg *config.Config) (*SegmentRepository, error) {
	db, err := gorm.Open(postgres.Open(cfg.PgURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error opening gorm: %v", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, fmt.Errorf("Error AutoMigrate: %v", err)
	}

	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Error retrieving orders: %v", err)
	}

	return &SegmentRepository{db}, nil
}

func (sr *SegmentRepository) GetSegmentByID(orderUID string) (*model.User, error) {
	return nil, nil
}

func (sr *SegmentRepository) CreateOrder(order *model.User) error {
	return nil
}