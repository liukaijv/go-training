package common

import (
	"testing"
	"time"
)

func TestDispatcher(t *testing.T) {

	d := NewDispatcher()

	d.AddFunc("aa", nil, func() {
		time.Sleep(1*time.Second)
	})

	d.Run("aa","1")
}
