package security

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sayamphoo/microservice/enum"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type logFile struct {
	fileLog  *os.File
	fileName string
}

func (lf *logFile) createFile() {
	data := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))

	if lf.fileName != data {
		lf.fileLog.Close()
	}

	logFileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
	fileLog, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการสร้างไฟล์ log:", err)
		return
	}

	lf.fileLog = fileLog
}
func LogDetect() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var requestData map[string]interface{}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			return
		}
		
		allParams := c.Params
		paramsMap := make(map[string]string)
		for _, param := range allParams {
			paramsMap[param.Key] = param.Value
		}

		jsonHeader, _ := json.Marshal(c.Request.Header)
		jsonParam, _ := json.Marshal(paramsMap)
		jsonBody, _ := json.Marshal(requestData)

		c.Set(enum.REQUEST_DATA, string(jsonBody))

		// ดำเนินการต่อไป
		c.Next()

		elapsed := time.Since(start)

		req := fmt.Sprintf("\n--> %-5s | %-5s |\t'%s'\n",
			c.Request.Method,
			c.ClientIP(),
			c.Request.URL.Path,
		)

		req += fmt.Sprintf("\t\tHEARDER => [%s]\n",
			strings.ReplaceAll(string(jsonHeader), "\n", ""),
		)
		req += fmt.Sprintf("\t\tBODY\t=> [%s]\n",
			strings.ReplaceAll(string(jsonBody), "\n", ""),
		)

		req += fmt.Sprintf("\t\tPARAMS\t=> [%s]",
			strings.ReplaceAll(string(jsonParam), "\n", ""),
		)

		req += fmt.Sprintf("\n<-- CODE %d  | %s\n", c.Writer.Status(), elapsed)

		log.Println(req)
	}
}

func FileLog() *os.File {
	logData := logFile{}
	logData.createFile()
	return logData.fileLog
}
