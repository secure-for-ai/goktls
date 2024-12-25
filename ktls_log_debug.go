//go:build debug

package tls

import (
	"log"
)

const Dev = true

func Debugln(a ...any) {
	log.Println(a...)
}

func Debugf(format string, a ...any) {
	log.Printf(format, a...)
}
