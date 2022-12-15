package work

import (
	"crypto/sha256"
	"fmt"
)

func SimpleProcedure(input string) string {
	return input
}

var hasher = sha256.New()

func HashProcedure(input string) string {
	hasher.Reset()

	for i := 0; i < 100000; i++ {
		hasher.Write([]byte(input))
	}

	return fmt.Sprintf("%x", hasher.Sum(nil))
}
