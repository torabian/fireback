package workspaces

import (
	"sync"
	"time"
)

type ConcurrentUpdater[T any] struct {
	ID    string
	mu    sync.RWMutex
	timer *time.Ticker
	data  T // Holds data of any type

	Done chan bool
}

// Update method now uses a generic manipulation function
func (cu *ConcurrentUpdater[T]) Update(manipulate func(*T)) {
	cu.mu.Lock()
	defer cu.mu.Unlock()
	cu.timer.Reset(time.Second)

	manipulate(&cu.data)
}

// Read method returns the generic data type
func (cu *ConcurrentUpdater[T]) Read() T {
	cu.mu.RLock()
	defer cu.mu.RUnlock()
	return cu.data
}

// General store for objects of any type
type ObjectStore[T any] struct {
	sync.RWMutex
	objects  map[string]*ConcurrentUpdater[T]
	OnCommit func(object *ConcurrentUpdater[T])
}

// Initialize a new ObjectStore
func NewObjectStore[T any]() *ObjectStore[T] {
	return &ObjectStore[T]{objects: make(map[string]*ConcurrentUpdater[T])}
}

func (store *ObjectStore[T]) GetOrCreateObject(objectID string, createFunc func() T) *ConcurrentUpdater[T] {
	store.RLock()
	object, exists := store.objects[objectID]
	store.RUnlock()

	if exists {
		return object
	}

	store.Lock()
	defer store.Unlock()
	commitTimer := time.NewTicker(time.Second)

	object = &ConcurrentUpdater[T]{
		ID:    objectID,
		timer: commitTimer,
		data:  createFunc(),
		Done:  make(chan bool),
	}

	go func(obj *ConcurrentUpdater[T]) {
		for {
			<-obj.timer.C
			obj.timer.Stop()
			store.OnCommit(obj)
			obj.Done <- true
		}
	}(object)

	store.objects[objectID] = object
	return object
}

// func sample_test() {
// 	objectStore := NewObjectStore[string]()
// 	objectStore.OnCommit = func(object *ConcurrentUpdater[string]) {
// 		// In the end, you can write things into the database here
// 	}

// 	var wg sync.WaitGroup
// 	for id := 0; id <= 1000; id++ {

// 		wg.Add(1)
// 		go func() {

// 			store := objectStore.GetOrCreateObject(fmt.Sprintf("%v", id), func() string {
// 				return ""
// 			})

// 			for i := 0; i <= 100; i++ {
// 				go func(i int) {
// 					randomDelay := time.Duration(100+rand.Intn(200)) * time.Millisecond
// 					time.Sleep(randomDelay)
// 					store.Update(func(s *string) {
// 						// do something with string
// 					})
// 				}(i)
// 			}

// 			<-store.Done
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// }
