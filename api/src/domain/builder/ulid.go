package builder

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// NewULID cria um id
func NewULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
