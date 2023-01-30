### A Fork of Golang TLS Library with Kernel TLS Implementation

This Repo is to add the
**[Kernel TLS](https://docs.kernel.org/networking/tls.html)** (KTLS)
support on top of the standard Golang TLS library.

The library is **100% percent compatible** with the original Golang TLS library.

The library enables the
kernel TLS after the TLS handshake is completed. In other words, the handshake
is implemented by golang itself, whereas the symmetric encryption is offloaded
to kernel.
#### Benefits
With KTLS, one can offload the encryption to the kernel or even the NIC,
depending on the kernel version and the hardware. This can significantly
reduce the CPU utility and improve the throughput of the TLS connection.

#### How to use
An example of the echo server is available at
https://github.com/secure-for-ai/goktls_examples.

- Enable the Kernel TLS on Linux in your terminal.
    ```bash
    sudo modprobe tls
    ```

- Add the package into `go.mod`
    ```go
    require github.com/secure-for-ai/goktls v1.20.0-rc3.1
    ```
    The release version tracks the Golang main release with minor updates.

- Drop in replacement

    Replace
    ```go
    import "crypto/tls"
    ```

    With
    ```go
    import tls "github.com/secure-for-ai/goktls"
    ```

- Run your code with KTLS
    ```bash
    GOKTLS=1 go run main.go
    ```

- Print out the debug information
    ```bash
    GOKTLS=1 go run -tag=debug main.go
    ```
    If KTLS is enabled, you should be able to see the following in kernel 5.15.
    ```
    2023/01/27 16:17:21 kTLS Enabled Status: true
    2023/01/27 16:17:21 Kernel Version: 5.15.0-58-lowlatency
    2023/01/27 16:17:21 ======Supported Features======
    2023/01/27 16:17:21 kTLS TX: true
    2023/01/27 16:17:21 kTLS RX: true
    2023/01/27 16:17:21 kTLS TLS 1.3 TX: true
    2023/01/27 16:17:21 kTLS TLS 1.3 RX: false
    2023/01/27 16:17:21 kTLS TX ZeroCopy: false
    2023/01/27 16:17:21 kTLS RX Expected No Pad: false
    2023/01/27 16:17:21 =========CipherSuites=========
    2023/01/27 16:17:21 kTLS AES-GCM-128: true
    2023/01/27 16:17:21 kTLS AES-GCM-256: true
    2023/01/27 16:17:21 kTLS CHACHA20POLY1305: true
    ```

    Once received the message, you will see
    ```
    2023/01/27 16:20:59 try to enable kernel tls AES_128_GCM for tls 1.2
    2023/01/27 16:20:59
    key: e3104f668dfa699b3bbc49b431a2f8dd
    iv: 2cea73b9
    seq: 0000000000000001
    2023/01/27 16:20:59 kTLS: TLS_TX enabled
    2023/01/27 16:20:59
    key: 299c717c1fa903bbe56abd30567fe09e
    iv: a8b78628
    seq: 0000000000000001
    2023/01/27 16:20:59 kTLS: TLS_RX enabled
    2023/01/27 16:20:59 kTLS: recvmsg, type: 23, payload len: 6
    2023/01/27 16:20:59 kTLS: recvmsg, type: 21, payload len: 2
    ```

#### Features
- KTLS 1.2 TX & RX
- KTLS 1.3 TX & RX
- zerocopy and no pad for TLS 1.3
- ciphersuites: AES-GCM-128, AES-GCM-256, CHACHA20POLY1305


#### TODO
1. KTLS 1.3 RX disabled on kernel < 5.19 as it causes weird package lost
2. zero copy and no pad have not been tested yet. zero copy is enabled
    on kernel >= 5.19, and no pad is enabled on kernel >= 6.0

#### Implementation

The implementation was based on [@jim3m](https://github.com/jim3ma/go/tree/dev.ktls.1.16.3)'s
implementation, which was based on [@FiloSottile](https://github.com/FiloSottile/go/commit/dbed9972d9947eb0001e9f5b639e0df05acec8bd)'s implementation.
