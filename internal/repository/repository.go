package repository

import (
	"fmt"
	"github.com/vokinneberg/go-url-shortener-ddd/domain"
)

type MemRepository struct {
	data map[string]*domain.URL
}

func NewMemRepository() *MemRepository {
	r := MemRepository{data: make(map[string]*domain.URL, 0)}
	return &r
}

func (r *MemRepository) Find(id string) (*domain.URL, error) {
	urlFound, ok := r.data[id]
	if ok {
		return urlFound, nil
	}
	return nil, fmt.Errorf(`Not found`)
}

func (r *MemRepository) Save(url *domain.URL) error {
	r.data[url.ID] = url
	return nil
}
