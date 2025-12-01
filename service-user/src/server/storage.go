package server

import (
	"net/http"

	"github.com/payments/service-user/src"
)

type httpServer struct {
	http.Handler
}

type Server struct {
	Service src.ServiceList
}
