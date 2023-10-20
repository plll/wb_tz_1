package app

import (
	"context"
	"encoding/json"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	"github.com/plll/wb_tz_1/internal/store"
	"net/http"
	"time"
)

func Run() { //pathToConfig string) {
	server := &http.Server{
		Addr: "localhost:8181",
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost/wb_tz_1")
	if err != nil {
	}
	cache, err := lru.New[Orders](10)
	if err != nil {
	}
	sc, err := stan.Connect("test-cluster", "12", stan.NatsURL("127.0.0.1:4223"))
	if err != nil {
	}
	repositories := store.NewRepositories(db)
	s := server.NewServer(ctx, server, conn, sc, cache, repositories)
	_, err = go sc.Subscribe("orders",
		func(m *stan.Msg) {
			var o Orders
			err := json.Unmarshal(m.Data, &o)
			if err == nil {
				err = insert_into_db(&o)
				if err != nil {
					fmt.Print(err)
				}
			} else {
				fmt.Print(err)
			}
		},
		stan.AckWait(20*time.Second))
	s.Init()
	s.Run(s.ctx)
}
