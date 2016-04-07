package main

import (
    "fmt"
    "flag"
    "time"
    "net/http"
    "log"
    "github.com/gin-gonic/gin" // An http routing library similar to sinatra
    "github.com/Sirupsen/logrus" // A logging library - we could have used the built in logger, but this one has appenders, etc.
    metrics "github.com/rcrowley/go-metrics" // A Go port of the Coda Hale metrics library
    "math/rand"
    "errors"
)

const FLAG_FAULTY = "faulty"

func main() {
    // Parse command line arguments
    arguments := parseCommandLine()

    // setup a logrus logger as a writer for our metrics logging
    logger := logrus.New()
    writer := logger.Writer()
    defer writer.Close()

    // Kick off background metrics logger that will spit out metric info every 5 seconds
    go metrics.Log(metrics.DefaultRegistry, 12 * time.Second, log.New(writer, "metrics", log.Lmicroseconds))

    // Create a router with default middleware
    router := gin.Default()

    // Add middlewares. They are like servlet filters
    router.Use(counterMiddleware())
    if faulty, _ := arguments[FLAG_FAULTY].(bool) ; faulty {
        router.Use(faultyMiddleware())

    }

    // Create rest endpoints
    router.GET("/hello/:name", metricsWrapper("hello"), hello)
    router.GET("/goodbye/:name", metricsWrapper("goodbye"), goodbye)

    // Start listening
    router.Run()
}

func hello(ctx *gin.Context) {
    name := ctx.Param("name")
    ctx.String(http.StatusOK, "Server says, 'Hello, %s!'", name)
}

func goodbye(ctx *gin.Context) {
    name := ctx.Param("name")
    ctx.String(http.StatusOK, "Server says, 'Goodbye, %s.'", name)
}

func parseCommandLine() map[string]interface{} {
    arguments := make(map[string]interface{})
    faulty :=  flag.Bool(FLAG_FAULTY, false, "Should this server randomly 'fail'")

    flag.Parse()

    arguments[FLAG_FAULTY] = *faulty

    return arguments
}

func metricsWrapper(name string) gin.HandlerFunc {
    // Using a closure. This allows us to create middleware scoped objects and register them once.
    counter := metrics.NewCounter()
    timer := metrics.NewTimer()

    metrics.Register(fmt.Sprintf("api.%s.counter", name), counter)
    metrics.Register(fmt.Sprintf("api.%s.timer", name), timer)

    return func(ctx *gin.Context) {
        counter.Inc(1)
        start := time.Now()
        ctx.Next()
        timer.UpdateSince(start)
    }
}

func counterMiddleware() gin.HandlerFunc {
    // Using a closure. This allows us to create middleware scoped objects and register them once.
    counter := metrics.NewCounter()
    metrics.Register("api.all.counter", counter)

    return func(ctx *gin.Context) {
        counter.Inc(1)
        ctx.Next()
    }
}

func faultyMiddleware() gin.HandlerFunc {
    // Random failure middleware
    return func(ctx *gin.Context) {
        const randMax = 100
        const oneInWhat = 3
        chance := rand.Intn(randMax)
        if chance < randMax / oneInWhat { // Fail 1 in oneInWhat times
            ctx.AbortWithError(http.StatusInternalServerError, errors.New("Random failure"))
        } else {
            ctx.Next()
        }
    }
}
