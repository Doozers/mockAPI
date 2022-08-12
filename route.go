package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type workspaces struct {
	Name    string   `json:"workspace"`
	Channel []string `json:"channel"`
}

func AskRights() func(c *gin.Context) {
	data := []workspaces{
		{
			Name:    "berty",
			Channel: []string{"general"},
		},
		{
			Name:    "berty",
			Channel: []string{"vip"},
		},
		{
			Name:    "berty",
			Channel: []string{"admin"},
		},
		{
			Name:    "territori",
			Channel: []string{"general"},
		},
		{
			Name:    "territori",
			Channel: []string{"vip"},
		},
		{
			Name:    "berty",
			Channel: []string{"admin"},
		},
	}

	return func(c *gin.Context) {
		if c.Request.Header.Get("pubKey") == "" {
			c.JSON(401, gin.H{
				"message": "missing pubKey",
			})
			return
		}

		rand.Seed(time.Now().Unix())
		var perm []workspaces
		perm = append(perm, data[0])
		perm = append(perm, data[3])

		for _, v := range []int{rand.Intn(4), rand.Intn(4)} {
			perm = append(perm, data[[]int{1, 2, 4, 5}[v]])
		}
		c.JSON(200, gin.H{
			"access": perm,
		})
	}
}
