package app

import (
	"context"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	"github.com/plll/wb_tz_1/internal/datastruct"
	"github.com/plll/wb_tz_1/internal/server"
	"github.com/plll/wb_tz_1/internal/store"
	"net/http"
)

func Run() {
	localserver := &http.Server{
		Addr: "localhost:8181",
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:Lepuhe63@localhost/wb_tz_1")
	if err != nil {
	}
	cache, err := lru.New[string, datastruct.Order](10)
	if err != nil {
	}
	sc, err := stan.Connect("test-cluster", "12", stan.NatsURL("127.0.0.1:4223"))
	if err != nil {
	}
	repositories := store.NewRepositories(conn)
	s := server.NewServer(ctx, localserver, conn, &sc, cache, repositories)
	s.Init()
	s.Run(ctx)
}
