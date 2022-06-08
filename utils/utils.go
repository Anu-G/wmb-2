package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateId() string {
	rand.Seed(int64(time.Now().Nanosecond()))

	return fmt.Sprintf("TX-%03d", rand.Intn(1000))
}
