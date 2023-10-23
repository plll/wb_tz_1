package server

import (
	"context"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	"github.com/plll/wb_tz_1/internal/store"
	"net/http"
)

type Server struct {
	ctx    context.Context
	db     *pgx.Conn
	server *http.Server
	cache  *lru.Cache
	sc     *stan.Conn
	repos  *store.Repositories
}

func NewServer(
	ctx context.Context,
	server *http.Server,
	db *pgx.Conn,
	sc *stan.Conn,
	cache *lru.Cache,
	repos *store.Repositories,
) *Server {
	return &Server{
		ctx:    ctx,
		db:     db,
		server: server,
		cache:  cache,
		sc:     sc,
		repos:  repos,
	}
}

func (s *Server) Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", s.ordersHandler)

	go s.createSubscriber()

	s.server.Handler = mux

	s.prepareCache()
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Print("Error in Run process")
		}
	}()
}

func (s *Server) Shutdown() {
	s.server.Shutdown(s.ctx)
}
