package emble

import "time"

type someStruct struct {
	UUID string
	Name string
	CreationTime time.Time
	Money uint
}

func newSomeStruct() *someStruct{
	return &someStruct{ 
		UUID: "uuid",
		Name: "Alex",
		CreationTime: time.Now(),
		Money: 0,
	}
}