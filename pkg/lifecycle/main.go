package lifecycle

var releaseFuncs []func()
var initFuncs []func()

func RegisterInit(f func()) {
	initFuncs = append(initFuncs, f)
}

func RegisterRelease(f func()) {
	releaseFuncs = append(releaseFuncs, f)
}

func Release() {
	for _, f := range releaseFuncs {
		defer recover()
		f()
	}
}

func Init() {
	for _, f := range initFuncs {
		f()
	}
}
