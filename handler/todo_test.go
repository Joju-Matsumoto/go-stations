package handler

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	t.Skip()
	s := ""
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(i)
}
