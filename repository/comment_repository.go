package repository

import (
	"context"
	"golang_database/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, data entity.Comments) (entity.Comments, error)
	FindById(ctx context.Context, id int32) (entity.Comments, error)
	FindAll(ctx context.Context) ([]entity.Comments, error)
}
