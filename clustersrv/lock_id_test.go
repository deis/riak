package clustersrv

import (
	"testing"
)

func TestConcurrentEquals(t *testing.T) {
	lid := NewLockID()
	origID := "abc"
	lid.id = origID
	for i := 0; i < 10; i++ {
		go func(i int) {
			if !lid.Equals(lid.id) {
				t.Errorf("goroutine %d reported lock ID != %s", i, origID)
			}
		}(i)
	}
}

func TestGenerate(t *testing.T) {
	lid := NewLockID()
	origID := "abc"
	lid.id = origID
	if !lid.Equals(lid.id) {
		t.Errorf("lock ID wasn't equal %s before generate", origID)
	}
	newID := lid.Generate()
	if newID == origID {
		t.Errorf("new lock ID value %s was equal old %s", newID, origID)
	}
	if lid.Equals(origID) {
		t.Errorf("lock ID was equal %s after generate", origID)
	}

}
