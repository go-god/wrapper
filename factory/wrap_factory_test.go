package factory

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	wrap := New(WgWrapper)
	wrap.Wrap(func() {
		log.Println(1111)
	})

	wrap.Wait()
}
