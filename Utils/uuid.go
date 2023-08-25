package Utils

import (
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

func Setuuid() string {
	id := uuid.New()
	return id.String()
}

func Randint() string {
	Number := strconv.Itoa(rand.Intn(999999-100000) + 100000)
	return Number
}
