//go:build !debug
package tls

const Dev = false

func Debugln(a ...any) {}

func Debugf(format string, a ...any) {}
