package main

import (
	"bytes"
	"log"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func logMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Authorization: %v", c.Request.Header["Authorization"]) // DO NOT use in production
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		log.Println("Response body: " + blw.body.String()) // DO NOT use in production. se below code instead.
		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			//ok this is an request with error, let's make a record for it
			// now print body (or log in your preferred way)
			log.Println("Response body: " + blw.body.String())
		}
	}
}
