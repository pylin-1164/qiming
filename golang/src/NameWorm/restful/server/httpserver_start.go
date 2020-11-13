package server

import (
	config "NameWorm/common"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
	"time"
)

type StartRestfulApi struct {

}
/**
 * 启动restful api服务
 */
func (s StartRestfulApi) StartApi(){
	StartRestfulServer()
	restful.DefaultContainer.Add(WS)
	fmt.Printf("start restful api listening on localhost:%d \n",config.API_SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%d",config.API_SERVER_PORT), nil)
}

/**
 *  启动restful api服务
 */
func (s StartRestfulApi) StartSSLapi(){
	var certFile = "/data/cert/server/server.crt"
	var keyFile = "/data/cert/server/server_key.pem"

	StartRestfulServer()
	wsContainer := restful.DefaultContainer
	wsContainer.Add(WS)

	sslconfig := &tls.Config{}
	s.inittls(sslconfig, certFile, keyFile)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d",config.API_SERVER_PORT),
		Handler:      wsContainer,
		TLSConfig:    sslconfig,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	fmt.Printf("start tls restful api listening on localhost:8099 \n")
	srv.ListenAndServeTLS(certFile, keyFile)
}


func (s StartRestfulApi) inittls(cfg *tls.Config, certFile, keyFile string) {
	crt, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic("init TLS Restful service fail")
	}
	cfg.Time = time.Now
	cfg.Rand = rand.Reader
	cfg.Certificates = []tls.Certificate{crt}
}
