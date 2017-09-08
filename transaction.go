package dilawar

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID          uint64
	Amount      int
	Type        Type
	Description string
	Time        time.Time
}

func (t Transaction) String() string {
	return fmt.Sprintf("%d in %s for (%s) on %s", t.Amount, t.Type, t.Description, t.Time)
}
