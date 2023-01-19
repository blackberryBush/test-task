package sth

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HandleYear(c *gin.Context) {
	if year, err := strconv.Atoi(c.Param("year")); err == nil {
		today := time.Now()
		thatTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
		difference := int(today.Sub(thatTime).Hours() / 24)
		if difference < 0 {
			c.String(http.StatusOK, "Осталось %v дней", -difference)
		} else {
			c.String(http.StatusOK, "Прошло %v дней", difference)
		}
	} else {
		log.Println("year converting error: ", err)
	}
}

func CheckPingHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-PING") == "ping" {
			c.Header("X-PONG", "pong")
		}
	}
}
