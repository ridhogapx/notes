package lib

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hash := HashPassword("myPass")

	fmt.Println("Password:", hash)
	fmt.Println("Length:", len(hash))
}
