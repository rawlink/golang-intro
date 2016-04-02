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

const FLAG_SERVER_NAME = "servername"
const FLAG_PORT = "port"
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
    router.Use(configurationMiddleware(arguments))
    router.Use(counterMiddleware())
    if faulty, _ := arguments[FLAG_FAULTY].(bool) ; faulty {
        router.Use(faultyMiddleware())

    }

    // Create a rest endpoint
    router.GET("/hello/:name", metricsWrapper("hello"), hello)
    router.GET("/goodbye/:name", metricsWrapper("goodbye"), goodbye)

    // Start listening
    router.Run(fmt.Sprintf(":%d", arguments[FLAG_PORT]))
}

func hello(ctx *gin.Context) {
    servername, _ := ctx.Get(FLAG_SERVER_NAME)
    name := ctx.Param("name")
    ctx.String(http.StatusOK, "%s says, 'Hello, %s!'", servername, name)
}

func goodbye(ctx *gin.Context) {
    servername, _ := ctx.Get(FLAG_SERVER_NAME)
    name := ctx.Param("name")
    ctx.String(http.StatusOK, "%s says, 'Goodbye, %s.'", servername, name)
}

func parseCommandLine() map[string]interface{} {
    arguments := make(map[string]interface{})
    serverName := flag.String(FLAG_SERVER_NAME,"Unknown","The name of the server")
    port := flag.Int(FLAG_PORT, 8080, "The port to listen on")
    faulty :=  flag.Bool(FLAG_FAULTY, false, "Should this server randomly 'fail'")

    flag.Parse()

    arguments[FLAG_SERVER_NAME] = *serverName
    arguments[FLAG_PORT] = *port
    arguments[FLAG_FAULTY] = *faulty

    return arguments
}

func metricsWrapper(name string) gin.HandlerFunc {
    // Using a closure. This allows us to create singletons and register them once.
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

func configurationMiddleware(config map[string]interface{}) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        for key, value := range config {
            ctx.Set(key, value)
        }
        ctx.Next()
    }
}

func counterMiddleware() gin.HandlerFunc {
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
        const oneInWhat = 2
        chance := rand.Intn(randMax)
        if chance < randMax / oneInWhat { // Fail 1 in oneInWhat times
            ctx.AbortWithError(http.StatusInternalServerError, errors.New("Random failure"))
        } else {
            ctx.Next()
        }
    }
}
