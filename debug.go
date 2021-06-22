package fivetran

var debug struct {
	debug     bool
	debugAuth bool
}

func Debug(b bool) {
	debug.debug = b
}

func DebugAuth(b bool) {
	debug.debugAuth = b
}
