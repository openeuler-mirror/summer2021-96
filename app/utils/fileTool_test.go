package utils

import (
	"fmt"
	"testing"
)

func TestFileRead(t *testing.T) {
	fmt.Println(ReadFile("H:\\redis.conf"))
}