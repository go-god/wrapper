package chanwrap

import (
	"log"
	"testing"
)

func TestWrapper(t *testing.T) {
	var wg = New(2)
	wg.Wrap(func() {
		log.Println("1111")
	})
	wg.WrapWithRecover(func() {
		log.Println(2222)
		panic("mock panic test")
	})
	wg.Wait()

	// time.Sleep(10 * time.Second)
}
