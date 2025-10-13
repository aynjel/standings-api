package logs

import (
	"net/http"

	"anggi.tabulation/utils/logs"
	"github.com/gin-gonic/gin"
)

// TestLogs fires sample Info/Warning/Error logs and returns status.
func TestLogs(c *gin.Context) {
	logs.InfoPrintln("Test info log entry from /api/logs/test")
	logs.WarningPrintln("Test warning log entry from /api/logs/test")
	logs.ErrorPrintln("Test error log entry from /api/logs/test")
	c.JSON(http.StatusOK, gin.H{"status": "logged"})
}
