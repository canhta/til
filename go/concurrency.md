# Go Concurrency

Go provides robust concurrency primitives that allow developers to write efficient and concurrent programs.

## Goroutines

Goroutines enable concurrent execution of functions.

```go
func main() {
    go func() {
        // Concurrent task
    }()
    // Main task
}
```

## Channels

Channels enable communication and synchronization between goroutines.

#### Unbuffered Channels

```go
ch := make(chan int) // Create an unbuffered channel

go func() {
    ch <- 42 // Send data into the channel
}()

data := <-ch // Receive data from the channel
```

#### Buffered Channel

```go
ch := make(chan int, 3) // Create a buffered channel with capacity 3

go func() {
    ch <- 1
    ch <- 2
    close(ch) // Close the channel after sending data
}()

for num := range ch {
    // Receive data from the channel until it's closed
}
```

## Select Statement

The select statement allows goroutines to wait on multiple communication operations.

```go
select {
case data := <-ch1:
    // Handle data from ch1
case data := <-ch2:
    // Handle data from ch2
}
```

## WaitGroup

`sync.WaitGroup` waits for a collection of goroutines to finish.

```go
var wg sync.WaitGroup
wg.Add(2) // Number of goroutines to wait for

go func() {
    defer wg.Done()
    // Task 1
}()

go func() {
    defer wg.Done()
    // Task 2
}()

wg.Wait() // Wait until all goroutines finish

```

## Worker Pools

Worker pools distribute tasks among multiple workers.

```go
var jobs = make(chan int, 10)
var results = make(chan int, 10)

func worker(id int) {
    for job := range jobs {
        // Process job
        results <- job * 2
    }
}

func main() {
    const numWorkers = 3
    for i := 1; i <= numWorkers; i++ {
        go worker(i)
    }

    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs)

    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Println("Result:", result)
    }
}
```

## Context

The `context` package manages cancellation signals, deadlines, and values across API boundaries.

```go
func process(ctx context.Context) {
    select {
    case <-time.After(2 * time.Second):
        fmt.Println("Task completed successfully")
    case <-ctx.Done():
        fmt.Println("Task canceled:", ctx.Err())
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        time.Sleep(1 * time.Second)
        cancel() // Cancel the task after 1 second
    }()

    process(ctx)
}
```

## Realtime Examples

### Goroutines and Channels: real-time chat application

Imagine you are building a real-time chat application. Each incoming message from different users needs to be processed concurrently. Goroutines and channels can help achieve this.

```go
type Message struct {
    User    string
    Content string
}

func processMessages(messages <-chan Message) {
    for msg := range messages {
        // Process incoming message asynchronously
        fmt.Printf("Received message from %s: %s\n", msg.User, msg.Content)
    }
}

func main() {
    messages := make(chan Message, 100) // Buffered channel for incoming messages

    // Start message processing goroutine
    go processMessages(messages)

    // Simulate receiving messages
    go func() {
        for i := 1; i <= 10; i++ {
            user := fmt.Sprintf("User%d", i)
            content := fmt.Sprintf("Hello from %s!", user)
            messages <- Message{User: user, Content: content}
            time.Sleep(time.Second) // Simulate delay between messages
        }
        close(messages) // Close the channel after sending all messages
    }()

    // Keep the main goroutine running
    select {}
}
```

In this example, messages from different users are concurrently processed in the `processMessages` goroutine using an unbuffered channel. This real-time scenario demonstrates the power of goroutines and channels in concurrent applications.

### Worker Pools: Handling HTTP Requests Concurrently

Consider a web server handling incoming HTTP requests. Using worker pools, you can process these requests concurrently.

```go
type Request struct {
    ID  int
    URL string
}

func worker(id int, requests <-chan Request, results chan<- string) {
    for req := range requests {
        // Process HTTP request asynchronously
        response := fmt.Sprintf("Worker %d processed request %d to URL: %s", id, req.ID, req.URL)
        results <- response
    }
}

func main() {
    numWorkers := 3
    requests := make(chan Request, 10)
    results := make(chan string, 10)

    // Start worker pool
    for i := 1; i <= numWorkers; i++ {
        go worker(i, requests, results)
    }

    // Simulate incoming HTTP requests
    go func() {
        for i := 1; i <= 10; i++ {
            url := fmt.Sprintf("http://example.com/%d", i)
            requests <- Request{ID: i, URL: url}
            time.Sleep(time.Second) // Simulate request delay
        }
        close(requests) // Close the channel after sending all requests
    }()

    // Collect and handle results
    go func() {
        for response := range results {
            fmt.Println(response)
        }
    }()

    // Keep the main goroutine running
    select {}
}
```

In this example, HTTP requests are processed concurrently by multiple workers. The main goroutine sends requests to the worker pool, which processes them concurrently and sends back the responses for further handling.

### Context: Graceful Shutdown of a Server

When shutting down a server, it's essential to gracefully handle in-flight requests. Contexts can help manage the lifecycle of server operations.

```go
func handleRequest(ctx context.Context, req *http.Request) {
    // Simulate request processing
    select {
    case <-time.After(time.Second):
        fmt.Println("Request processed successfully")
    case <-ctx.Done():
        fmt.Println("Request canceled:", ctx.Err())
    }
}

func main() {
    server := &http.Server{
        Addr: ":8080",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            ctx, cancel := context.WithCancel(context.Background())

            // Handle request asynchronously
            go func() {
                defer cancel() // Ensure cancel is called after request processing
                handleRequest(ctx, req)
            }()

            // Respond to the client
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Request received and processing"))
        }),
    }

    // Start the server in a goroutine
    go func() {
        if err := server.ListenAndServe(); err != nil {
            fmt.Println("Server error:", err)
        }
    }()

    // Gracefully shutdown the server on SIGINT or SIGTERM
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    <-sigCh

    // Shutdown the server gracefully
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        fmt.Println("Error during server shutdown:", err)
    } else {
        fmt.Println("Server shut down gracefully")
    }
}
```

In this example, the server listens for incoming requests. When a shutdown signal (SIGINT or SIGTERM) is received, the server is gracefully shut down. In-flight requests are allowed to finish processing, ensuring a smooth transition during server shutdown.
