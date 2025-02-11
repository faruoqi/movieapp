package memory

import (
	"context"
	"github.com/faruoqi/movieapp/metadata/internal/repository"
	"github.com/faruoqi/movieapp/metadata/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
