package designpatterns

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

// var of Single instance
var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating singleton instance ...")
		instance = &Singleton{data: "Singleton data"}
	})
	return instance
}

func Singletonmain() {
	// Create a instance of Singleton:
	singleton1 := GetInstance()
	fmt.Println(singleton1.data)

	// Try Creating a 2nd instance of Singleton:
	singleton2 := GetInstance()
	fmt.Println(singleton2)
}
