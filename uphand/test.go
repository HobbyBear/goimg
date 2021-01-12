package uphand

import (
	"github.com/gin-gonic/gin"
)

// 测试上传
func Test(c *gin.Context) {

	w := c.Writer

	html := `
<html>
    <head>
        <title>测试上传</title>
    </head>
    <body>
        <div style="padding: 20px;">
            <p>测试上传</p>
            <form action="/" method="post" enctype="multipart/form-data">
                <p><input type="file" name="userfile" /></p>
                <p><input type="submit" name="submit" value="上传" /></p>
            </form>
        </div>
    </body>
</html>`

	w.Write([]byte(html))

}
