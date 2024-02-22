package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id        uuid.UUID `gorm:"id"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"<-:update"`
	DeletedAt time.Time `gorm:"<-:update"`
}

type ListCategoriesReq struct {
}

type ListCategoriesResp struct {
	Count      int
	Categories []Category
}
