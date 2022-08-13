package core

import "sort"

type Collection interface {
	sort.Interface
	Json
	First(keys ...string) interface{}
	IsEmpty() bool
	Last(keys ...string) interface{}
	Offset(index int, item interface{}) Collection
	Sort(sorter interface{}) Collection
	Pluck(key string) Fields
	Prepend(item ...interface{}) Collection
	Pull(defaultValue ...interface{}) interface{}
	Push(items ...interface{}) Collection
	Put(index int, item interface{}) Collection
	Shift(defaultValue ...interface{}) interface{}
}

type Json interface {
	ToJson() string
}
