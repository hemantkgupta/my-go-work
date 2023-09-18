package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Create the certificate key for client:
openssl genrsa -out client.key 2048

Create the certificate-signing request (CSR) for the client:
openssl req -new -key client.key -out client.csr

Generate the certificate for the client using the TSL CSR and key along with the CA Root key:
openssl x509 -req -in client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out client.crt -days 500 -sha256

Run below command to test
curl https://localhost:4443 --insecure --cert client.crt --key client.key

*/

func httpRequestHandler2(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello,World!\n"))
}
func main() {
	var (
		CACertFilePath = "./rootCA.crt"
		CertFilePath   = "./localhost.crt"
		KeyFilePath    = "./localhost.key"
	)
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}
	// Configure the server to trust TLS client cert issued by your CA.
	certPool := x509.NewCertPool()
	if caCertPEM, err := ioutil.ReadFile(CACertFilePath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}
	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		Certificates: []tls.Certificate{serverTLSCert},
	}
	server := http.Server{
		Addr:      ":4443",
		Handler:   http.HandlerFunc(httpRequestHandler2),
		TLSConfig: tlsConfig,
	}
	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
