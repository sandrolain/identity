package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/sandrolain/go-utilities/pkg/mongoutils"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/grpc/admingrpc"
	"github.com/sandrolain/identity/src/grpc/clientgrpc"
)

func main() {
	cfg, err := config.GetConfiguration()
	if err != nil {
		log.Fatalf("cannot load environment configuration: %v", err)
	}

	fmt.Printf("cfg: %v\n", cfg)

	_, err = mongoutils.NewClient(cfg.MongoDB.URI, cfg.MongoDB.Database, time.Second*time.Duration(cfg.MongoDB.Timeout))
	if err != nil {
		log.Fatalf("cannot create MongoDB client: %v", err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		err = admingrpc.StartServer(cfg.AdminGRPC.Port)
		if err != nil {
			log.Fatalf("cannot start admin gRPC server: %v", err)
		}
		wg.Done()
	}()

	go func() {
		err = clientgrpc.StartServer(cfg.ClientGRPC.Port)
		if err != nil {
			log.Fatalf("cannot start client gRPC server: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
