package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
)

type repositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &repositoryImpl{DB: db}
}

func (r *repositoryImpl) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {
	query := "INSERT INTO comments(email,comment)VALUES(?,?)"
	result, err := r.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.ID = int32(id)
	return comment, nil
}

func (r *repositoryImpl) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	comment := entity.Comments{}
	query := "SELECT * FROM comments WHERE id = ? LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Data Not Found")
	}
}

func (r *repositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	var comments []entity.Comments
	query := "SELECT * FROM comments"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return comments, err
	}

	defer rows.Close()
	for rows.Next() {
		comment := entity.Comments{}
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
