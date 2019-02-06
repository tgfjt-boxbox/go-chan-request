package utils

import (
	"net/http"
	"sync"
)

var once sync.Once
var c *http.Client

func GetClient() *http.Client {
	once.Do(func() {
		c = &http.Client{}
	})

	return c
}
