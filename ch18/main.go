package main

import (
	"context"
	"fmt"
	"github.com/Andrew-M-C/go.timeconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

type keys struct{}
type traceIDCtx struct{}

var incrNum uint64

func main() {
	//t := time.Date(2019, 1, 31, 0, 0, 0, 0, time.UTC)   // 2019-01-31

	//nt := t.AddDate(0, 1, 0)    // 后推一个月
	//fmt.Printf("%v\n", nt)        // 2019-03-03 00:00:00 +0000 UTC, 不是我们想要的

	nt := timeconv.AddDate(time.Now(), 0, -1, 0)
	fmt.Printf("%v\n", nt) // 2019-02-28 00:00:00 +0000 UTC，这就对了
	ctx := context.Background()
	ctx = context.WithValue(ctx, keys{}, "__test__")
	fmt.Println(ctx.Value(keys{}))
	r := gin.Default()
	r.GET("/", hello)
	r.Run(":9988")
}
func hello(c *gin.Context) {
	traceID := c.GetHeader("X-Request-Id")
	if traceID == "" {
		traceID = newTraceID()
	}
	c.Request = c.Request.WithContext(NewTraceID(c.Request.Context(), traceID))
	c.JSON(http.StatusOK, c.Request.Context().Value(traceIDCtx{}))
}
func NewTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtx{}, traceID)
}

func newTraceID() string {
	return fmt.Sprintf("trace-id-%d-%s-%d",
		os.Getpid(),
		time.Now().Format("2006.01.02.15.04.05.999"),
		atomic.AddUint64(&incrNum, 1))
}
