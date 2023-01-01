package service

import (
	"sync"
	"time"

	"github.com/sandrolain/go-utilities/pkg/logutils"
	"github.com/sandrolain/identity/src/api"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/grpc/admingrpc"
	"github.com/sandrolain/identity/src/grpc/clientgrpc"
	"github.com/sandrolain/identity/src/storage/mongostorage"
	"github.com/sandrolain/identity/src/storage/redisstorage"
)

func RunService() {
	logutils.InitLogger()

	cfg, err := config.GetConfiguration()
	if err != nil {
		logutils.Fatalf("cannot load environment configuration: %v", err)
	}

	mongodbStorage, err := mongostorage.CreateMongoDBStorage(cfg.MongoDb.Uri, cfg.MongoDb.Database, time.Duration(cfg.MongoDb.Timeout)*time.Second)
	if err != nil {
		logutils.Fatalf("cannot create MongoDB storage client: %v", err)
	}

	redisStorage, err := redisstorage.CreateRedisStorage(cfg.Redis.Host, cfg.Redis.Password, nil, time.Duration(cfg.Redis.Timeout)*time.Second)
	if err != nil {
		logutils.Fatalf("cannot create Redis storage client: %v", err)
	}

	api := &api.API{
		Config:            cfg,
		VolatileStorage:   redisStorage,
		PersistentStorage: mongodbStorage,
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		err = admingrpc.StartServer(api)
		if err != nil {
			logutils.Fatalf("cannot start admin gRPC server: %v", err)
		}
		wg.Done()
	}()

	go func() {
		err = clientgrpc.StartServer(api)
		if err != nil {
			logutils.Fatalf("cannot start client gRPC server: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
