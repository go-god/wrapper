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

	for i := 0; i < 10*10000; i++ {
		// The method of copying is used here to avoid the i
		// in the wrap func being the same variable
		index := i
		wg.Wrap(func() {
			log.Printf("current index: %d\n", index)
		})
	}

	wg.WrapWithRecover(func() {
		panic("mock panic test")
	})

	wg.Wait()
}

/*
2022/01/15 17:58:32 current index: 99998
2022/01/15 17:58:32 current index: 99999
2022/01/15 17:58:32 exec recover:mock panic test
--- PASS: TestWrapper (4.66s)
PASS
*/
