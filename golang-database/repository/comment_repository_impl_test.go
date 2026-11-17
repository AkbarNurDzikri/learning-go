package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	
	comment := entity.Comment{
		Email: "repo@gmail.com",
		Comment: "Terserah lu dah",
	}
	result, err := commentRepository.Insert(context.Background(), comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 70)
	
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}