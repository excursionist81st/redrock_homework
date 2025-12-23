package main

import "github.com/gin-gonic/gin"

type Student struct {
	Name  string    `json:"name"`
	Score []float64 `json:"score"`
}

func average_get(scores []float64) float64 {
	sum := 0.0
	for _, score := range scores {
		sum += score
	}
	return sum / float64(len(scores))
}
func average_do(c *gin.Context) {
	var request Student
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	res_average := average_get(request.Score)
	c.JSON(200, gin.H{
		"average": res_average,
	})
}
func main() {
	r := gin.Default()
	r.POST("/", average_do)
	r.Run(":80")
}
