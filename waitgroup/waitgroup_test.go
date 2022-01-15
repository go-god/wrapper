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
	wg := New(wrapper.WithRecover(mockRecovery))
	wg.Wrap(func() {
		log.Println("1111")
	})

	num := 10 * 10000
	for i := 0; i < num; i++ {
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
$ go test -v
2022/01/15 18:15:50 current index: 99993
2022/01/15 18:15:50 current index: 99994
2022/01/15 18:15:50 exec recover:mock panic test
2022/01/15 18:15:50 current index: 99991
2022/01/15 18:15:50 current index: 99995
2022/01/15 18:15:50 current index: 99999
2022/01/15 18:15:50 current index: 99997
2022/01/15 18:15:50 current index: 99998
--- PASS: TestWrapper (4.06s)
PASS
*/
