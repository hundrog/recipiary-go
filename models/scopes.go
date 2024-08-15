package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Query struct {
	Page      int  `form:"page"`
	Limitless bool `form:"limitless"`
}

const PageSize = 10

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var query Query
		if err := c.ShouldBindQuery(&query); err != nil {
			return db.Offset(0).Limit(PageSize)
		}

		if query.Limitless {
			return db
		}

		offset := (query.Page - 1) * PageSize
		return db.Offset(offset).Limit(PageSize)
	}
}
