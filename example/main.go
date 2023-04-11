package main

import (
	"log"

	"github.com/go-god/wrapper"
)

func main() {
	chWrap := wrapper.New(wrapper.ChWrapper, wrapper.WithBufCap(2))
	chWrap.Wrap(func() {
		log.Println("chan wrapper: 1111")
	})
	chWrap.Wrap(func() {
		log.Println("chan wrapper: 2222")
	})

	chWrap.Wait()

	// factory.WgWrapper No need to pass second parameter.
	wg := wrapper.New(wrapper.WgWrapper)
	wg.Wrap(func() {
		log.Println("wg wrapper:1111")
	})
	wg.Wrap(func() {
		log.Println("wg wrapper:2222")
	})

	wg.WrapWithRecover(func() {
		log.Println("wg wrapper:3333")
		panic("mock panic:abc")
	})

	wg.Wait()
}

// output:
// 2022/01/09 22:47:48 chan wrapper: 2222
// 2022/01/09 22:47:48 chan wrapper: 1111
// 2022/01/09 22:47:48 wg wrapper:3333
// 2022/01/09 22:47:48 wg wrapper:1111
// 2022/01/09 22:47:48 wrapper exec recover:mock panic:abc
// 2022/01/09 22:47:48 wg wrapper:2222
