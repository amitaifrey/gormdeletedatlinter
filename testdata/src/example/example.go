package example

import (
	"time"

	"gorm.io/gorm"
)

// Good: *gorm.DeletedAt
type ModelCorrect struct {
	ID        string
	DeletedAt *gorm.DeletedAt
}

// Bad: gorm.DeletedAt (non-pointer)
type ModelNonPointer struct {
	ID        string
	DeletedAt gorm.DeletedAt // want `DeletedAt field should be \*gorm\.DeletedAt`
}

// Bad: time.Time
type ModelTime struct {
	ID        string
	DeletedAt time.Time // want `DeletedAt field should be \*gorm\.DeletedAt`
}

// Bad: *time.Time
type ModelTimePtr struct {
	ID        string
	DeletedAt *time.Time // want `DeletedAt field should be \*gorm\.DeletedAt`
}

// Bad: int64
type ModelInt struct {
	ID        string
	DeletedAt int64 // want `DeletedAt field should be \*gorm\.DeletedAt`
}

// Good: no DeletedAt field at all
type ModelNoDeletedAt struct {
	ID   string
	Name string
}

// Good: *gorm.DeletedAt with gorm tags
type ModelWithTags struct {
	ID        string
	DeletedAt *gorm.DeletedAt `gorm:"type:timestamptz"`
}
