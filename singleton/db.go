package singleton

import "sync"

var (
  instance *db
)

type db struct {}

func (d db) GetUser(ID string) string {
  return ""
}

var lock = &sync.Mutex{}

var once sync.Once

var onceBody = func() {
  instance = &db{}
}

// GetDBInstance is using singleton pattern
func GetDBInstanceWithOnce() *db {
  once.Do(onceBody)
  return instance
}

func GetDBInstanceWithLock() *db {
  if instance == nil {
    lock.Lock()
    defer lock.Unlock()
    onceBody()
  }
  return instance
}
