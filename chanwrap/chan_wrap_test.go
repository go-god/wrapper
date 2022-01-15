package chanwrap

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
	num := 10 * 10000
	c := num + 2
	chWrap := New(wrapper.WithBufCap(c), wrapper.WithRecover(mockRecovery))
	chWrap.Wrap(func() {
		log.Println("1111")
	})

	for i := 0; i < num; i++ {
		// The method of copying is used here to avoid the i
		// in the wrap func being the same variable
		index := i
		chWrap.Wrap(func() {
			log.Printf("current index: %d\n", index)
		})
	}

	chWrap.WrapWithRecover(func() {
		log.Println(2222)
		panic("mock panic test")
	})
	chWrap.Wait()
}

/**
$ go test -v
2022/01/15 18:06:42 current index: 99992
2022/01/15 18:06:42 current index: 99996
2022/01/15 18:06:42 current index: 99998
2022/01/15 18:06:42 current index: 99997
2022/01/15 18:06:42 current index: 99999
2022/01/15 18:06:42 2222
2022/01/15 18:06:45 exec recover:mock panic test
--- PASS: TestWrapper (3.70s)
*/
