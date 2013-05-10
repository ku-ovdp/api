// Package dummy implements dummy storage for API entities
package dummy

import (
	"github.com/ku-ovdp/api/persistence"
)

type dummyBackend struct{}

func init() {
	persistence.Register("dummy", dummyBackend{})
}
