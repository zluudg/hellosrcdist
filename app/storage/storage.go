package storage

import (
    "github.com/zluudg/hellosrcdist/app/interfaces"
)

type Config struct {
    Log interfaces.Logger
}

type storage struct {
    log interfaces.Logger
}

func Create(conf Config) (*storage, error) {
    stor := storage{
    }

    if conf.Log == nil {
        stor.log = interfaces.NullLogger{}
    } else {
        stor.log = conf.Log
    }

    return &stor, nil
}

func (s storage) Initialize() error {
    s.log.Info("Initializing storage")
    return nil
}

func (s storage) Shutdown() {
    s.log.Info("Shutting down storage")
}
