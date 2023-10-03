package httputils

var debug struct {
	enable     bool
	authEnable bool
}

// Debug sets debug
func Debug(b bool) {
	debug.enable = b
}

// DebugAuth sets authentication debug
func DebugAuth(b bool) {
	debug.authEnable = b
}
