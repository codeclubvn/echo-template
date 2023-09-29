package utils

import (
	"echo_template/infra"
	"echo_template/presenter/request"
	"gorm.io/gorm"
)

func ErrNoRows(err error) bool {
	return err == gorm.ErrRecordNotFound
}

func MustHaveDb(db *infra.Database) {
	if db == nil {
		panic("Database engine is null")
	}
}

type QueryPaginationBuilder[E any] struct {
	db *infra.Database
}

func QueryPagination[E any](db *gorm.DB, o request.PageOptions, data *[]*E) error {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	offset := (o.Page - 1) * o.Limit

	if err := db.Debug().Offset(int(offset)).Limit(int(o.Limit)).Find(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *QueryPaginationBuilder[E]) Count(total *int64) *QueryPaginationBuilder[E] {
	q.db.Count(total)
	return q
}

func (q *QueryPaginationBuilder[E]) Error() error {
	return q.db.Error
}
