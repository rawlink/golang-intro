# Web Service
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
