package app

import (
    "context"
	"testing"
)

func TestAppVersionCmdNoError(t *testing.T) {
	var application App
    application.Ctx = context.Background()

    err := application.Run(CmdVersion)
    if err != nil {
        t.Fatalf("unexpected error: '%s'", err)
    }
}
