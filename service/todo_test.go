package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/TechBowl-japan/go-stations/db"
)

func TestTODO(t *testing.T) {
	t.Skip()
	db, err := db.NewDB("test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	service := NewTODOService(db)

	todo, err := service.CreateTODO(context.Background(), "sub1", "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(todo)
}
