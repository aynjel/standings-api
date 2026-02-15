package app

import (
	"github.com/gin-gonic/gin"
)

// ServeScalarUI serves the Scalar API documentation UI
func ServeScalarUI() gin.HandlerFunc {
	return func(c *gin.Context) {
		html := `<!doctype html>
<html>
  <head>
    <title>Tabulation API Documentation</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script id="api-reference" data-url="/swagger/doc.json"></script>
    <script>
      var configuration = {
        theme: 'purple',
      }

      var apiReference = document.getElementById('api-reference')
      apiReference.dataset.configuration = JSON.stringify(configuration)
    </script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html)
	}
}
