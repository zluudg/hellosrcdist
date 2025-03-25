package storage

import (
	"testing"

    "github.com/zluudg/hellosrcdist/app/interfaces"
)

func TestStorageCreateNoErrorEmptyConf(t *testing.T) {
    _, err := Create(Config{})

    if err != nil {
        t.Fatalf("Unexpeced error during creation: '%s'", err)
    }
}

func TestStorageCreateNoErrorConfWithLogger(t *testing.T) {
    conf := Config {
        Log: interfaces.NullLogger{},
    }
    _, err := Create(conf)

    if err != nil {
        t.Fatalf("Unexpeced error during creation: '%s'", err)
    }
}

func TestStorageInitializeNoErrorEmptyConf(t *testing.T) {
    stor, err := Create(Config{})

    if err != nil {
        t.Fatalf("Unexpeced error during creation: '%s'", err)
    }

    err = stor.Initialize()
    if err != nil {
        t.Fatalf("Unexpeced error during initialization: '%s'", err)
    }
}

func TestStorageShutdownNoPanicEmptyConf(t *testing.T) {
    stor, err := Create(Config{})

    if err != nil {
        t.Fatalf("Unexpeced error during creation: '%s'", err)
    }

    stor.Shutdown()
}
