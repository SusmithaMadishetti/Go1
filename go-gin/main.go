package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//Joke structure which contains single joke
type Joke struct {
	ID    int    `json:"id" binding "required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding :"required"`
}

var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./views/js", true)))
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", jokeHandler)
		api.POST("/jokes/like/:jokeId", likeJokeHandler)

	}
	r.Run(":3000")
}
func jokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}
func likeJokeHandler(c *gin.Context) {
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes++
			}
		}
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
