// Author(s): Carl Saldanha

package model

import (
	"time"
	"github.com/gofrs/uuid"
		  "github.com/jinzhu/gorm"
		    _ "github.com/jinzhu/gorm/dialects/sqlite"
)
// Build defines information about server build
type Task struct {
	ID          uuid.UUID `json:"id" gorm:"primarykey;type:uuid;column:id"`
	Title        string    `json:"title"`
	Description        string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}


func (t *Task) Insert(db *gorm.DB) error{
	if dbc := db.Create(&t); dbc.Error != nil {
	    // Create failed, do something e.g. return, panic etc.
	        return dbc.Error
	}
        return nil
}

func GetTasks(db *gorm.DB) ([]*Task, error){
	var t []*Task
	if dbc := db.Find(&t); dbc.Error != nil {
	    // Create failed, do something e.g. return, panic etc.
	        return nil, dbc.Error
	}
        return t, nil
}

