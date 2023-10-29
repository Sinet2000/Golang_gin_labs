# My Go Project

This is a simple Go project that demonstrates how to manage dependencies, build, test, and run a Go application.

## Prerequisites

Before you start, make sure you have Go installed. You can download and install it from [https://golang.org/dl/](https://golang.org/dl/).

Based on https://github.com/demo-apps/go-gin-app
https://go.dev/doc/tutorial/web-service-gin

```bash
$ go get -u github.com/gin-gonic/gin
$ go build -o app.out
$ ./app.out
```

## Getting Started

1. **Initialize module**:
This will create a go.mod file to manage dependencies.

   ```shell
   go mod init my-go-project
   ```

2. **Get All Dependencies**:
To install all listed dependencies, use the following command:

   ```shell
   go get ./...
   ```

2. **Get All Dependencies**:
To install all listed dependencies, use the following command:

   ```shell
   go get ./...
   ```

---
## Development

1. **Building the Project:**:
To install all listed dependencies, use the following command:

   ```shell
   go build -o app.out
   ```

2. **Testing:**:

   ```shell
   go test
   ```

3. **Running the Application:**:

   ```shell
   ./app.out
   ```

4. **Get a Package:**:

   ```shell
   go get package-name
   ```

5. **Update Dependencies:**:

   ```shell
   go get -u ./...
   ```
6.  **Remove a Package:**:

   ```shell
   go list -u | xargs go get -u
   ```

---

## Other Useful Commands

**Generate Documentation:**:

   ```shell
   godoc -http :6060
   ```