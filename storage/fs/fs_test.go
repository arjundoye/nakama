package fs

import (
	"testing"

	"github.com/nicolasparada/nakama/storage/tests"
)

func TestStore(t *testing.T) {
	tests.RunStoreTests(t, &Store{
		Root: t.TempDir(),
	}, "")
}
