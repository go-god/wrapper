package waitgroup

import (
	"log"
	"testing"

	"github.com/go-god/wrapper"
)

func mockRecovery() {
	if err := recover(); err != nil {
		log.Printf("exec recover:%v\n", err)
	}
}

func TestWrapper(t *testing.T) {
	var wg = New(wrapper.WithRecover(mockRecovery))
	wg.Wrap(func() {
		log.Println("1111")
	})
	wg.WrapWithRecover(func() {
		log.Println(2222)
		panic("mock panic test")
	})
	wg.Wait()
}
