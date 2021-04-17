package db

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

//Init initializes the database
func Init() error {
	var err error
	db, err = bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//RemoveTask removes task from the database
func RemoveTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

//AddTask adds task in the database
func AddTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		key, _ := b.NextSequence()
		return b.Put(itob(int(key)), []byte(task))
	})
}

//ListTasks lists all tasks in the database
func ListTasks() ([]string, error) {
	var tasks []string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Println(k)
			tasks = append(tasks, string(v))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
