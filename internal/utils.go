package utils

import (
	"math/rand"
	"reflect"
)

func IteratorStruct(s interface{}, iteratorFn func(key string, value reflect.Value)) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i).Name
		value := reflect.ValueOf(s).FieldByName(key)

		iteratorFn(key, value)
	}
}

func UniqueId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
