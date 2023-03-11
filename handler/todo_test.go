package handler

import (
	"fmt"
	"strconv"
	"strings"
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

func TestRepeat(t *testing.T) {
	ids := []int{1, 2, 3, 4, 5}
	s := fmt.Sprintf(strings.Repeat("%s, ", len(ids)), ids)
	fmt.Println(s)
}
