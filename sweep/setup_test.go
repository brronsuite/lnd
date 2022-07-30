package sweep

import (
	"testing"

	"github.com/brronsuite/lnd/kvdb"
)

func TestMain(m *testing.M) {
	kvdb.RunTests(m)
}
