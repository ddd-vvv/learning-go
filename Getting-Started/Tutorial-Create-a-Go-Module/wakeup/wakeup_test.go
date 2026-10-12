package wakeup

import (
	"regexp"
	"testing"
)

func TestWakeupName(t *testing.T) {
	name := "Neo"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := WakeUp("Neo")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`WakeUp("Neo") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestWakeUpEmpty(t *testing.T) {
	msg, err := WakeUp("")
	if msg != "" || err == nil {
		t.Errorf(`WakeUp("") = %q, %v, want "", error`, msg, err)
	}
}
