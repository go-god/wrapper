package factory

import (
	"github.com/go-god/wrapper"
	"github.com/go-god/wrapper/chanwrap"
	"github.com/go-god/wrapper/waitgroup"
)

type constructor func(opts ...wrapper.Option) wrapper.Wrapper

// WrapType wrap type
type WrapType int

const (
	// WgWrapper waitGroup wrapper
	WgWrapper WrapType = iota
	// ChWrapper chan wrapper
	ChWrapper
)

var wrapperMap = map[WrapType]constructor{
	WgWrapper: waitgroup.New,
	ChWrapper: chanwrap.New,
}

// New create wrapper interface
func New(wrapType WrapType, opts ...wrapper.Option) wrapper.Wrapper {
	if w, ok := wrapperMap[wrapType]; ok {
		return w(opts...)
	}

	panic("wrapper type not exists")
}

// Register register wrapper
func Register(wrapType WrapType, c constructor) {
	_, ok := wrapperMap[wrapType]
	if ok {
		panic("registered injector already exists")
	}

	wrapperMap[wrapType] = c
}
