package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyParams struct {
	Spliced  bool `json:"spliced"`
	Left     int  `json:"left"`
	X        int  `json:"x"`
	Offset   int  `json:"offset"`
	Verified bool `json:"verified"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.Static("/static", "./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"status": "ok"})
	})
	r.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "这是login页面",
		})
	})
	r.GET("/web1", func(c *gin.Context) {
		c.Header("my-header-value", "test000")
		c.JSON(403, gin.H{
			"response": "这是web1页面",
		})
	})
	r.POST("/isVerify", func(c *gin.Context) {
		datazz := make([]int, 0)
		c.ShouldBindJSON(&datazz)
		datas := make([]int, 0)
		for i := 0; i < len(datazz); i++ {
			c := math.Abs(float64(datazz[i]))
			if c >= 0 && c <= 9 {
				datas = append(datas, int(c))
			}
		}
		sum := 0
		for _, data := range datas {
			sum += data
		}
		avg := float64(sum) / float64(len(datas))
		sum2 := 0.0
		for _, data := range datas {
			sum2 += math.Pow(float64(data)-avg, 2)
		}
		stddev := sum2 / float64(len(datas))
		c.JSON(http.StatusOK, stddev != 0)
		if stddev != 0 {
			fmt.Println("ok")
		} else {
			fmt.Println("fail")
		}
	})
	r.POST("/check", func(c *gin.Context) {
		var pm VerifyParams
		c.ShouldBindJSON(&pm)
		spliced := pm.Spliced
		left := pm.Left     //当前拖动块的左边缘距离其父元素左边缘的距离
		x := pm.X           //当前滑块的中心点距离其父元素左边缘的距离
		offset := pm.Offset //允许的最大偏差值 滑块中心点与目标位置中心点的最大允许距离
		verify := pm.Verified
		serverCompute := math.Abs(float64(left-x)) < float64(offset)
		if spliced && verify && serverCompute {
			c.JSON(http.StatusOK, gin.H{
				"status": "true",
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "false",
			})
		}
	})
	r.Run(":8080")
}
