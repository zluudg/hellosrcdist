package interfaces

import (
	"testing"
)


func TestInterfaceNullStorageInitializeNoError(t *testing.T) {
    stor := NullStorage{}

    err := stor.Initialize()
    if err != nil {
        t.Fatalf("Unexpeced error: '%s'", err)
    }
}
