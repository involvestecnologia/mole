package main

import (
	"github.com/involvestecnologia/mole/pkg/config"
	"github.com/involvestecnologia/mole/pkg/mongodb"
	"github.com/involvestecnologia/mole/pkg/storage"
	notifierModels "github.com/involvestecnologia/notify/pkg/models"
	"github.com/involvestecnologia/notify/pkg/notifiers"
)

func main() {

	conf := config.Load()

	notifier := notifiers.MM(conf.Notifier.URL, notifierModels.Options{
		DefaultSender:       conf.AppName,
		DefaultDestinations: []string{conf.Notifier.Channel},
	})

	mongo := mongodb.New(conf.Mongo)
	storage := storage.Elasticsearch(conf.Elasticsearch)

	oplogReplication := mongodb.NewOplogReplication(mongo, storage)
	if err := oplogReplication.Start(); err != nil {
		_ = notifier.Notify("", nil, err.Error(), "Oplog replication process failed")
	}
}
