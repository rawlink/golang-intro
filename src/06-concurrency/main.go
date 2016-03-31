package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    goroutines()
    blockingChannels()
    bufferedChannels()
    channelsWithRangeAndClose()
    channelsWithSelect()
}

func toBeGoroutined(i int) {
    time.Sleep(1 * time.Second)
    fmt.Println("Goroutine", i)
}

func goroutines() {
    fmt.Println("Goroutines")
    fmt.Println("==========")

    start := time.Now()

    for i := 0 ; i < 20 ; i++ {
        go toBeGoroutined(i)
    }

    fmt.Println("Starting up all goroutines took", time.Since(start))

    // Wait for all our goroutines to finish. There are better ways to do this, but that comes with channels...
    time.Sleep(2 * time.Second)

    fmt.Println()
}

func summation(c chan int, n int) {
    time.Sleep(1 * time.Second)
    c <- (n * (n + 1) / 2)
}

func blockingChannels() {
    fmt.Println("Blocking channels")
    fmt.Println("=================")

    // This is a blocking channel. Sends or receives to this channel will block until the other side is ready.
    sumChannel := make(chan int)

    // Creating a job to do a summation
    fmt.Println("Starting summation job.")
    go summation(sumChannel, 15)
    fmt.Println("Waiting for job to complete.")
    // Seeing as we created a blocking channel, the statement below will block until the answer comes back.
    fmt.Println("The summation from 1 to 15 is", <- sumChannel)

    // Go will detect when all threads are deadlocked.
    //<- sumChannel
    //sumChannel <- 1

    // And yes, you can totally use a channel as a threadsafe queue.

    fmt.Println()
}

func bufferedChannels() {
    fmt.Println("Buffered channels")
    fmt.Println("=================")

    // Buffered channels block for a send when they are full or a receive when they are empty.
    bufferedChannel := make(chan int, 2)

    bufferedChannel <- 1
    bufferedChannel <- 2
    //bufferedChannel <- 3

    fmt.Println("Channel item is", <- bufferedChannel)
    fmt.Println("Channel item is", <- bufferedChannel)
    //fmt.Println("Channel item is", <- bufferedChannel)

    fmt.Println()
}
func runThenClose(c chan int) {
    for i := 0 ; i < 10 ; i++ {
        c <- rand.Intn(100)
    }

    // Only the sender should close
    close(c);

    // A channel send after a close will result in a panic
    //c <- 1
}

func channelsWithRangeAndClose() {
    fmt.Println("Channels, range, and close")
    fmt.Println("==========================")

    rand.Seed(time.Now().UnixNano())
    c := make(chan int)

    go runThenClose(c)

    for i := range c {
        fmt.Println("Received channel value", i)
    }

    fmt.Println()
}

func channelsWithSelect() {
    fmt.Println("Channels and select")
    fmt.Println("===================")

    events := time.Tick(500 * time.Millisecond) // Creates a channel that fires at a regular interval
    shutdown := time.After(3 * time.Second) // Creates a channel that fires after a duration of time has passed

    fmt.Println("Starting receive...")
    for {
        // Select allows you to wait on several channels at once. If there is a default case it will fire if no other
        // case is ready. If there is no default case, the select will block until 1 case is available. If multiple
        // cases are available, one will be randomly chosen.
        select {
        case <- events:
            fmt.Println("\tReceived an event")
        case <- shutdown:
            fmt.Println("Received shutdown - exiting")
            fmt.Println()
            return
        default:
            fmt.Println("\tHeartbeat...")
            time.Sleep(250 * time.Millisecond)
        }
    }
}
