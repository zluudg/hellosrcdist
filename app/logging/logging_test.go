package logging

import (
	"testing"
)

func TestLoggingCreateNoError(t *testing.T) {
    var tests = []struct {
        name     string
        indata   int
        expected error
    }{
        {"1", 1, nil},
        {"2", 2, nil},
        {"3", 3, nil},
        {"4", 4, nil},
        {"5", 5, nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, got := Create(tt.indata)
            if got != tt.expected {
                t.Fatalf("got %q, expected %q", got, tt.expected)
            }
        })
    }
}

func TestLoggingFormat(t *testing.T) {
    var tests = []struct {
        name           string
        indata_fmtstr  string
        indata_varargs []any
        expected       string
    }{
        {"NO_VARARGS",  "hello", []any{},           "hello"},
        {"NIL_VARARGS", "hello", nil,               "hello"},
        {"ONE_VARARGS", "hello %s", []any{"world"}, "hello world"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := format(tt.indata_fmtstr, tt.indata_varargs)
            if got != tt.expected {
                t.Fatalf("got %s, expected %s", got, tt.expected)
            }
        })
    }
}

func TestLoggingDebugNoPanic(t *testing.T) {
    var tests = []struct {
        name     string
        indata   int
        expected error
    }{
        {"1", 1, nil},
        {"2", 2, nil},
        {"3", 3, nil},
        {"4", 4, nil},
        {"5", 5, nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l, _ := Create(tt.indata)
            l.Debug("nothing")
        })
    }
}

func TestLoggingInfoNoPanic(t *testing.T) {
    var tests = []struct {
        name     string
        indata   int
        expected error
    }{
        {"1", 1, nil},
        {"2", 2, nil},
        {"3", 3, nil},
        {"4", 4, nil},
        {"5", 5, nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l, _ := Create(tt.indata)
            l.Info("nothing")
        })
    }
}

func TestLoggingWarningNoPanic(t *testing.T) {
    var tests = []struct {
        name     string
        indata   int
        expected error
    }{
        {"1", 1, nil},
        {"2", 2, nil},
        {"3", 3, nil},
        {"4", 4, nil},
        {"5", 5, nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l, _ := Create(tt.indata)
            l.Warning("nothing")
        })
    }
}

func TestLoggingErrorNoPanic(t *testing.T) {
    var tests = []struct {
        name     string
        indata   int
        expected error
    }{
        {"1", 1, nil},
        {"2", 2, nil},
        {"3", 3, nil},
        {"4", 4, nil},
        {"5", 5, nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l, _ := Create(tt.indata)
            l.Error("nothing")
        })
    }
}
