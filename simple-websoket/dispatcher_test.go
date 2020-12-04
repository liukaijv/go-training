package simple_websoket

import (
	"fmt"
	"testing"
)

func TestDispatcher(t *testing.T) {

	d := NewDispatcher()

	d.AddFunc("aa", func(res string) {
		fmt.Println(res)
	})

	d.Run("aa", "1")
}
