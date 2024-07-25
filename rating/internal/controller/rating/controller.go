package rating

import (
	"context"
	"errors"
	"github.com/faruoqi/movieapp/rating/internal/repository"
	"github.com/faruoqi/movieapp/rating/pkg/model"
)

var ErrNotFound = errors.New("not found")

type ratingRepository interface {
	Get(ctx context.Context, RecordID model.RecordID, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, RecordID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

type Controller struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetAggregatedRating(ctx context.Context, RecordID model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, RecordID, recordType)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}

	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}

	return sum / float64(len(ratings)), nil

}

func (c *Controller) PutRating(ctx context.Context, RecordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return c.repo.Put(ctx, RecordID, recordType, rating)

}
