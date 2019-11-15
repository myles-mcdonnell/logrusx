package logrusx

import (
	"os"
	"testing"
)

func Test(t *testing.T) {

	factory := NewLogEntryFactory("test")
	Init(DEBUG, os.Stdout, &JSONFormatter{Indent: true}, factory)

	Error("test").Write()
}
