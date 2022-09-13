package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
}

func (Database) CreateSingleConnection() {
	fmt.Println("Create single connection")
	time.Sleep(2 * time.Second)
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating a DB conection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB alredy created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}

	wg.Wait()
}
