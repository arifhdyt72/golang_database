package repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"log"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()

	comment := entity.Comments{
		Email:   "test.repository@test.com",
		Comment: "Comment repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(result)
}

func TestCommentFindByID(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, int32(35))
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()

	result, err := CommentRepository.FindAll(ctx)
	if err != nil {
		log.Panic(err.Error())
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
