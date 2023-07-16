package token

import (
	"testing"
)

func TestID(t *testing.T) {
	t1 := NewTokenID()
	s := t1.String()
	t2, err := NewTokenIDFromString(s)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	if t1.NUID != t2.NUID {
		t.Errorf("mismatch t1:%s t2:%s", t1.NUID, t2.NUID)
	}
	if t1.Counter != t2.Counter {
		t.Errorf("mismatch t1:%d t2:%d", t1.Counter, t2.Counter)
	}

}
func TestRotate(t *testing.T) {
	t1 := NewTokenID()
	id := t1.NUID
	counter := t1.Counter
	t1.Rotate()

	if id != t1.NUID {
		t.Errorf("mismatch ")
	}
	if counter != t1.Counter-1 {
		t.Errorf("mismatch ")
	}
}
