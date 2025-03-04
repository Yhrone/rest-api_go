package db

import "sync"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	memoryDB = make(map[string]User)
	mutex    = sync.RWMutex{}
)

func SaveToMemoryDB(user User) {
	mutex.Lock()
	defer mutex.Unlock()
	memoryDB[user.ID] = user
}

func FetchFromMemoryDB() []User {
	mutex.RLock()
	defer mutex.RUnlock()

	var users []User
	for _, user := range memoryDB {
		users = append(users, user)
	}
	return users
}
