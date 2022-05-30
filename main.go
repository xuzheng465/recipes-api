package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.Run()
}

type Recipe struct {
	Name         string    `json:"name"`
	Tag          []string  `json:"tag"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"published_at"`
}
