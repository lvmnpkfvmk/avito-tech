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

type ISegmentRepository interface {
	CreateSegment(segment *model.Segment) error
	DeleteSegment(segment *model.Segment) error
}

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
		return nil, fmt.Errorf("Error AutoMigrate User: %v", err)
	}

	err = db.AutoMigrate(&model.Segment{})
	if err != nil {
		return nil, fmt.Errorf("Error AutoMigrate Segment: %v", err)
	}

	repo := &SegmentRepository{db}

	return repo, nil
}

func (sr *SegmentRepository) GetSegmentByID(orderUID string) (*model.User, error) {
	return nil, nil
}

func (sr *SegmentRepository) CreateSegment(segment *model.Segment) error {
	result := sr.db.Model(&model.Segment{}).Create(segment)
	if result.Error != nil {
		return fmt.Errorf("Error creating segment: %v", result.Error)
	}
	return nil
}

func (sr *SegmentRepository) DeleteSegment(segment *model.Segment) error {
	result := sr.db.Model(&model.Segment{}).Where("name = ?", segment.Name).Delete(segment)
	if result.Error != nil {
		return fmt.Errorf("Error deleting segment: %v", result.Error)
	}
	return nil
}