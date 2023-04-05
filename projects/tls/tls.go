package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"

	"github.com/spf13/pflag"
)

var flagCert, flagClient, flagServer bool

func init() {
	pflag.BoolVarP(&flagCert, "cert", "t", false, "Run certificate tests")
	pflag.BoolVarP(&flagClient, "client", "c", false, "Run as client")
	pflag.BoolVarP(&flagServer, "server", "s", false, "Run as server")
	pflag.Parse()
}

const usage = `tls
A little project about TLS
usage:
  --cert	-t
  --client	-c
  --server	-s
`

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.LstdFlags)

	switch {
	case flagCert:
		log.Print("flagCert")
	case flagClient:
		client()
	case flagServer:
		server()
	default:
		fmt.Println(usage)

	}

}

func client() {
	config := &tls.Config{
		ServerName: "learning-go.dev",
	}

	conn, err := tls.Dial("tcp", "localhost:9876", config)
	if err != nil {
		log.Panic(err.Error())
	}
	defer conn.Close()

	err = conn.VerifyHostname("learning-go.dev")
	if err != nil {
		log.Panic("Hostname doesn't match with certificate: " + err.Error())
	}

	_ = conn.ConnectionState().PeerCertificates[0].Issuer
}

func server() {
	_ = &x509.Certificate{}

	cert, err := tls.LoadX509KeyPair("public.crt", "private.key")
	if err != nil {
		log.Panic(err.Error())
	}
	log.Print("Certificate loaded...")

	certs := make([]tls.Certificate, 1)
	certs = append(certs, cert)

	config := &tls.Config{
		Certificates: certs,
	}

	ops := x509.VerifyOptions{
		DNSName: "learning-go.dev",
	}

	_ = ops

	//x509.Ve
	log.Print(cert.Certificate)

	listener, err := tls.Listen("tcp", "localhost:9876", config)
	if err != nil {
		log.Panic(err.Error())
	}
	defer listener.Close()
	log.Print("Listening...")

	conn, err := listener.Accept()
	if err != nil {
		log.Panic(err.Error())
	}

	go handleConnection(conn)

}

func handleConnection(net.Conn) {

}
