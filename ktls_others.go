//go:build !linux
// +build !linux

package tls

import (
	"net"
)

const kTLSOverhead = 0

func (c *Conn) enableKernelTLS(cipherSuiteID uint16, inKey, outKey, inIV, outIV []byte) error {
	return nil
}

func ktlsSendCtrlMessage(c *net.TCPConn, typ recordType, b []byte) (int, error) {
	panic("not implement")
}

func ktlsReadDataFromRecord(c *net.TCPConn, b []byte) (int, error) {
	panic("not implement")
}

func ktlsReadRecord(c *net.TCPConn, b []byte) (recordType, int, error) {
	panic("not implement")
}
