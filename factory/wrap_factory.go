package factory

import (
	"github.com/go-god/wrapper"
	"github.com/go-god/wrapper/chanwrap"
	"github.com/go-god/wrapper/waitgroup"
)

type constructor func() wrapper.Wrapper

const (
	// WgWrapper waitGroup wrapper
	WgWrapper = "wg"
	// ChWrapper chan wrapper
	ChWrapper = "ch"
)

var wrapperMap = map[string]constructor{
	WgWrapper: waitgroup.New,
	ChWrapper: nil,
}

// New create wrapper interface
func New(name string, c ...int) wrapper.Wrapper {
	if name == ChWrapper && len(c) > 0 && c[0] > 0 {
		return chanwrap.New(c[0])
	}

	if w, ok := wrapperMap[name]; ok {
		return w()
	}

	panic("wrapper type not exists")
}

// Register register gdi.Injector
func Register(name string, c constructor) {
	_, ok := wrapperMap[name]
	if ok {
		panic("registered injector already exists")
	}

	wrapperMap[name] = c
}
