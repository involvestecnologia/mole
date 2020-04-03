package storage

import (
	"time"

	"github.com/involvestecnologia/mole/models"
)

type Storage interface {
	Add(oplog *models.Oplog) error
	StartTime() (time.Time, error)
}
