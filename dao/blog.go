package dao

import (
	"github.com/xiaka53/eth-service/public"
)

type Blog struct {
	BlogNumber int `json:"blog_number"`
}

func (b *Blog) TableName() string {
	return "blog"
}

func (b *Blog) Update(number int) error {
	return public.SqlPool.Table(b.TableName()).Where(b).Update("blog_number", number).Error
}

func (b *Blog) First() error {
	return public.SqlPool.First(b).Error
}
