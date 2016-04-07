# A Web Service
In this section, you will work with a “real” program. This program is a web service using third-party libraries. Upon completion of this section you should have a basic understanding of the following:

* Vendored package management - Introduced as an expirement in Go 1.5. Turned on by default in Go 1.6.
* Glide package manager - A Go package manager that works with the new vendored package management.
  * Project initialization - glide create/init
  * Dependency installation - glide install
  * Other dependency management - glide get/update/remove/list/tree
* Some third-party libraries
  * [logrus](https://github.com/Sirupsen/logrus) - Logging
  * [gin](https://github.com/gin-gonic/gin) - REST API and middleware
  * [go-metrics](https://github.com/rcrowley/go-metrics) - Go port of DropWizard Metrics
  * [hystrix-go](https://github.com/afex/hystrix-go) - Go port of Netflix Hystrix

## Building This Project

Before you can compile this example, you will need to install its third-party dependencies using Glide. To install the dependencies and build this project perform the following steps:

0. In a terminal/cmd window, change directories into this project.
0. Type `glide install` and press enter. You should see output similar to the following:

```
    [INFO] Downloading dependencies. Please wait...
    [INFO] Fetching updates for gopkg.in/go-playground/validator.v8.
    [INFO] Fetching updates for golang.org/x/sys.

    ...

    [INFO] Setting version for github.com/manucorporat/sse to ee05b128a739a0fb76c7ebd3ae4810c1de808d6d.
    [INFO] Setting version for golang.org/x/sys to a60af9cbbc6ab800af4f2be864a31f423a0ae1f2.
    [INFO] Setting version for golang.org/x/net to f315505cf3349909cdf013ea56690da34e96a451.
    [INFO] Setting version for gopkg.in/go-playground/validator.v8 to c193cecd124b5cc722d7ee5538e945bdb3348435.
```

0. Type `go build server.go` and press enter.
0. Type `go build client.go` and press enter.

You should now have two binaries, ‘server’ and ‘client’. ‘server’ has an argument ‘--faulty’ that will cause it to fail 1/3 of the requests that it serves, exercising the hystrix usage in ‘client’.
