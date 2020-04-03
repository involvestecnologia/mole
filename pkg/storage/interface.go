package storage

import (
	"time"

	"github.com/involvestecnologia/mole/models"
)

//Storage - Responsible for storing the records
type Storage interface {

	//Add - Adds a record to the batch to be stored
	Add(oplog *models.Oplog) error

	//StartTime - Searches for the date of the last stored record
	StartTime() (time.Time, error)
}
