package main

import (
	"github.com/involvestecnologia/mole/internal/collectors"
	"github.com/involvestecnologia/mole/internal/handlers"
	"github.com/involvestecnologia/mole/internal/mongodb"
	"github.com/involvestecnologia/mole/internal/storage"
	"github.com/involvestecnologia/mole/pkg/config"
	"github.com/involvestecnologia/mole/pkg/logger"
	"github.com/labstack/echo"
)

func main() {

	conf := config.Load()

	storageCollector := collectors.NewStorageCollector()
	elasticsearch := storage.Elasticsearch(conf.Elasticsearch, storageCollector)

	log := logger.New(conf)
	mongo := mongodb.New(conf.Mongo)

	oplogCollector := collectors.NewOplogCollector()
	oplogReplication := mongodb.NewOplogReplication(mongo, elasticsearch, oplogCollector, log)
	go oplogReplication.Start()

	router := echo.New()
	handlers.NewPrometheusHandler(router)
	handlers.NewHealthHandler(router)

	log.Fatal(router.Start(":8080"))
}
