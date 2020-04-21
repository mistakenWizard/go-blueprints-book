package trace

import (
	"fmt"
	"io"
	"testing"
)

//Tracer Interface describes object capable of tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()

	silentTracer.Trace("something")
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

//Off creates a tracer that will ignore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
