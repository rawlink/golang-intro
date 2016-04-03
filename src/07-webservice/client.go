package main

import (
    "io/ioutil"
    "net/http"
    "fmt"
    "errors"
    "github.com/afex/hystrix-go/hystrix"
)

func main() {
    for i := 0 ; i < 10 ; i ++ {
        if body, err := fallbackGetHello() ; err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Success:", body)
        }
    }
}

func fallbackGetHello() (string, error) {
    output := make(chan string)
    errors := hystrix.Go(
        "getHello",
        func() error {
            // Contact the hello service
            if body, err := get("http://localhost:8080/hello/gopher") ; err != nil {
                return err
            } else {
                output <- body
                return nil
            }
        },
        func(err error) error {
            output <- "Fallback response"
            return nil
            //return errors.New("aaagh!")
        },
    )

    select {
    case out := <-output:
        //success
        return out, nil
    case err := <-errors:
        // failure
        return "", err
    }
}

func get(url string) (string, error) {
    if resp, err := http.Get(url) ; err != nil {
        return "", err
    } else {
        defer resp.Body.Close()
        if statusCode := resp.StatusCode ; statusCode != http.StatusOK {
            return "", errors.New(fmt.Sprintf("Unexpected status code %d returned.", statusCode))
        }
        if body, err := ioutil.ReadAll(resp.Body) ; err != nil {
            return "", err
        } else {
            return string(body), nil
        }
    }
}