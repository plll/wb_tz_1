package app

import (
	"context"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"github.com/plll/wb_tz_1/internal/datastruct"
	"github.com/plll/wb_tz_1/internal/server"
	"github.com/plll/wb_tz_1/internal/store"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func Run() {
	localserver := &http.Server{
		Addr: "localhost:8181",
	}
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	conn, err := pgxpool.Connect(ctx, "postgresql://postgres:postgres@localhost/wb_tz_1")
	if err != nil {
		log.Fatal(err)
	}
	cache, err := lru.New[string, datastruct.Order](10)
	if err != nil {
		log.Fatal(err)
	}
	sc, err := stan.Connect("test-cluster", "1", stan.NatsURL("127.0.0.1:4223"))
	if err != nil {
		log.Fatal(err)
	}
	repositories := store.NewRepositories(conn)
	s := server.NewServer(ctx, localserver, conn, sc, cache, repositories)
	s.Init()
	s.Run(ctx)
	<-ctx.Done()
	conn.Close()
	sc.Close()
	localserver.Close()
}
