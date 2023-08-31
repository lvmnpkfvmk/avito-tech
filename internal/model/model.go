package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Segments *Segments `gorm:"many2many:user_tags;"`
}

type Segment struct {
	Name string `gorm:"primarykey", json:"name"`
}

type Segments []Segment

func (is *Segments) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Error casting value to bytes")
	}

	var segments []Segment
	err := json.Unmarshal(bytes, &segments)
	if err != nil {
		return err
	}

	*is = segments
	return nil
}

func (is Segments) Value() (driver.Value, error) {
	return json.Marshal(is)
}
