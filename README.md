# openssl-channel-proxy

As native go does not support tls widely like openssl, I write a proxy for tcp or tls foward to openssl tls.
So you can connect to some special tls server by the proxy which the native go can not.
Currently only support tcp, tls will come soon.

## System required

### Ubuntu
```
apt-get install -y openssl libssl-dev
```

### Centos
```
yum install -y openssl openssl-devel
```

## Install
```
$ go get -u github.com/sanguohot/openssl-channel-proxy/cmd/openssl-channel-proxy
$ openssl-channel-proxy -h
Usage of openssl-channel-proxy:
  -l string
        local address (default ":8000")
  -r string
        remote address (default "10.6.250.54:8822")
  -r-ca string
        A PEM eoncoded ca's certificate file. (default "/opt/conf/ca.crt")
  -r-cert string
        A PEM eoncoded certificate file. (default "/opt/conf/sdk.crt")
  -r-key string
        A PEM encoded private key file. (default "/opt/conf/sdk.key")
```

## Simply test.


### Proxy side

```
$ openssl-channel-proxy -r 10.6.250.53:8822 -r-ca /root/key/ca-cert.pem -r-cert /root/key/client-cert.pem -r-key /root/key/client-key.pem
2018/08/15 10:32:26 New connection from: 192.168.5.98:60600
2018/08/15 10:32:26 Proxy connection closed (192.168.5.98:60600 -> 10.6.250.53:8822) sent: 12 B received: 0 B
```


### Client side

Write a simple tcp client and run it
```
package main
import (
	"net"
	"log"
	"fmt"
)
func main() {
	raddr, err := net.ResolveTCPAddr("tcp", "10.6.250.52:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 0xffff)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			b := buf[:n]
			fmt.Print(string(b))
		}
	}()
	conn.Write([]byte("hello world!"))
}
```


### Server side

```
$ openssl s_server -CAfile /root/key/ca-cert.pem -key /root/key/server-key.pem -cert /root/key/server-cert.pem -accept 8822
Using default temp DH parameters
ACCEPT
-----BEGIN SSL SESSION PARAMETERS-----
MGQCAQECAgMDBALAMAQABDBeoiwAsSHel7hV2rkHlF/RlrGm3STdJSUzgJ6/OZKW
lQdKGLtp28ShvcYJ6FIjNAOhBgIEW3OQuqIEAgIBLKQGBAQBAAAApg0ECzEwLjYu
MjUwLjUz
-----END SSL SESSION PARAMETERS-----
Shared ciphers:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DH-DSS-AES256-GCM-SHA384:DHE-DSS-AES256-GCM-SHA384:DH-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA256:DH-RSA-AES256-SHA256:DH-DSS-AES256-SHA256:DHE-RSA-AES256-SHA:DHE-DSS-AES256-SHA:DH-RSA-AES256-SHA:DH-DSS-AES256-SHA:DHE-RSA-CAMELLIA256-SHA:DHE-DSS-CAMELLIA256-SHA:DH-RSA-CAMELLIA256-SHA:DH-DSS-CAMELLIA256-SHA:ECDH-RSA-AES256-GCM-SHA384:ECDH-ECDSA-AES256-GCM-SHA384:ECDH-RSA-AES256-SHA384:ECDH-ECDSA-AES256-SHA384:ECDH-RSA-AES256-SHA:ECDH-ECDSA-AES256-SHA:AES256-GCM-SHA384:AES256-SHA256:AES256-SHA:CAMELLIA256-SHA:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:DH-DSS-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:DH-RSA-AES128-GCM-SHA256:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES128-SHA256:DHE-DSS-AES128-SHA256:DH-RSA-AES128-SHA256:DH-DSS-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA:DH-RSA-AES128-SHA:DH-DSS-AES128-SHA:DHE-RSA-SEED-SHA:DHE-DSS-SEED-SHA:DH-RSA-SEED-SHA:DH-DSS-SEED-SHA:DHE-RSA-CAMELLIA128-SHA:DHE-DSS-CAMELLIA128-SHA:DH-RSA-CAMELLIA128-SHA:DH-DSS-CAMELLIA128-SHA:ECDH-RSA-AES128-GCM-SHA256:ECDH-ECDSA-AES128-GCM-SHA256:ECDH-RSA-AES128-SHA256:ECDH-ECDSA-AES128-SHA256:ECDH-RSA-AES128-SHA:ECDH-ECDSA-AES128-SHA:AES128-GCM-SHA256:AES128-SHA256:AES128-SHA:SEED-SHA:CAMELLIA128-SHA:ECDHE-RSA-DES-CBC3-SHA:ECDHE-ECDSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:EDH-DSS-DES-CBC3-SHA:DH-RSA-DES-CBC3-SHA:DH-DSS-DES-CBC3-SHA:ECDH-RSA-DES-CBC3-SHA:ECDH-ECDSA-DES-CBC3-SHA:DES-CBC3-SHA:IDEA-CBC-SHA:ECDHE-RSA-RC4-SHA:ECDHE-ECDSA-RC4-SHA:ECDH-RSA-RC4-SHA:ECDH-ECDSA-RC4-SHA:RC4-SHA:RC4-MD5
Signature Algorithms: RSA+SHA512:DSA+SHA512:ECDSA+SHA512:RSA+SHA384:DSA+SHA384:ECDSA+SHA384:RSA+SHA256:DSA+SHA256:ECDSA+SHA256:RSA+SHA224:DSA+SHA224:ECDSA+SHA224:RSA+SHA1:DSA+SHA1:ECDSA+SHA1
Shared Signature Algorithms: RSA+SHA512:DSA+SHA512:ECDSA+SHA512:RSA+SHA384:DSA+SHA384:ECDSA+SHA384:RSA+SHA256:DSA+SHA256:ECDSA+SHA256:RSA+SHA224:DSA+SHA224:ECDSA+SHA224:RSA+SHA1:DSA+SHA1:ECDSA+SHA1
Supported Elliptic Curve Point Formats: uncompressed:ansiX962_compressed_prime:ansiX962_compressed_char2
Supported Elliptic Curves: P-256:P-521:P-384:secp256k1
Shared Elliptic curves: P-256:P-521:P-384:secp256k1
CIPHER is ECDHE-RSA-AES256-GCM-SHA384
Secure Renegotiation IS supported
hello world!DONE
shutting down SSL
CONNECTION CLOSED
ACCEPT
```