package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

/*
Create the root key:
openssl genrsa -des3 -out rootCA.key 4096
Create and self-sign the root certificate:
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.crt

Create the certificate key:
openssl genrsa -out localhost.key 2048
Create the certificate-signing request (CSR):
openssl req -new -key localhost.key -out localhost.csr
Generate the certificate using the TSL CSR and key along with the CA Root key:
openssl x509 -req -in localhost.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out localhost.crt -days 500 -sha256



*/

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Running HTTPS Server!!")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8443),
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
				// Always get latest localhost.crt and localhost.key
				// ex: keeping certificates file somewhere in global location
				// where created certificates updated and this closure function can refer that
				cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
				if err != nil {
					return nil, err
				}
				return &cert, nil
			},
		},
	}

	// run server on port "8443"
	log.Fatal(srv.ListenAndServeTLS("localhost.crt", "localhost.key"))
}
