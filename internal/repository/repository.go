package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/lvmnpkfvmk/avito-tech/config"
	"github.com/lvmnpkfvmk/avito-tech/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNotFound = errors.New("User not found")
)

type ISegmentRepository interface {
	CreateSegment(segment *model.Segment) error
	DeleteSegment(segment *model.Segment) error
	GetAllSegments() (*model.Segments, error)
	CreateUser(user *model.User) error
	GetUser(ID uint) (*model.User, error)
	GetAllUsers() (*[]model.User, error)
	UpdateUser(user *model.User) error
}

type Repository struct {
	db    *gorm.DB
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
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

	repo := &Repository{db}

	return repo, nil
}

func (sr *Repository) GetSegmentByID(orderUID string) (*model.User, error) {
	return nil, nil
}

func (sr *Repository) CreateSegment(segment *model.Segment) error {
	result := sr.db.FirstOrCreate(&model.Segment{}, segment)
	if result.Error != nil {
		return fmt.Errorf("Error creating segment: %v", result.Error)
	}
	return nil
}

func (sr *Repository) DeleteSegment(segment *model.Segment) error {
	result := sr.db.Model(&model.Segment{}).Where("name = ?", segment.Name).Delete(segment)
	if result.Error != nil {
		return fmt.Errorf("Error deleting segment: %v", result.Error)
	}
	return nil
}

func (sr *Repository) GetAllSegments() (*model.Segments, error) {
	var segments model.Segments
	result := sr.db.Model(&model.Segment{}).Find(&segments)
	if result.Error != nil {
		return nil, fmt.Errorf("Error getting segments: %v", result.Error)
	}
	return &segments, nil
}

func (sr *Repository) GetUser(ID uint) (*model.User, error) {
	user := model.User{}
	result := sr.db.First(&user, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("Error getting user: %v", result.Error)
	}
	return &user, nil
}
func (sr *Repository) GetAllUsers() (*[]model.User, error) {
	var users []model.User
	result := sr.db.Model(&model.User{}).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Error getting users: %v", result.Error)
	}
	return &users, nil
}

func (sr *Repository) CreateUser(user *model.User) error {
	result := sr.db.Model(&model.User{}).Create(user)
	if result.Error != nil {
		return fmt.Errorf("Error creating user: %v", result.Error)
	}
	return nil
}

func (sr *Repository) UpdateUser(user *model.User) error {
	result := sr.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("Error updating user: %v", result.Error)
	}
	return nil
}