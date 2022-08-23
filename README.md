[![Tests on Linux, MacOS and Windows](https://github.com/gohugoio/hugoreleaser-plugins-api/workflows/Test/badge.svg)](https://github.com/gohugoio/hugoreleaser-plugins-api/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/github.com/gohugoio/hugoreleaser-plugins-api)](https://goreportcard.com/report/github.com/gohugoio/hugoreleaser-plugins-api)
[![GoDoc](https://godoc.org/github.com/gohugoio/hugoreleaser-plugins-api?status.svg)](https://godoc.org/github.com/gohugoio/hugoreleaser-plugins-api)


Plugins API for https://github.com/gohugoio/hugoreleaser

A plugin is a [Go Module](https://go.dev/blog/using-go-modules) with a main func, e.g., using the API provided by the [archiveplugin](https://pkg.go.dev/github.com/gohugoio/hugoreleaser-plugins-api/archiveplugin) package:

```go
func main() {
	server, err := server.New(
		func(d server.Dispatcher, req archiveplugin.Request) archiveplugin.Response {
			d.Infof("Creating archive %s", req.OutFilename)

			if err := req.Init(); err != nil {
				// ... handle error.
			}

			if err := createArchive(req); err != nil {
				// ... handle error.
			}
			// Empty response is a success.
			return archiveplugin.Response{}
		},
	)
	if err != nil {
		log.Fatalf("Failed to create server: %s", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

	_ = server.Wait()
}
```

See the [Deb Plugin](https://github.com/gohugoio/hugoreleaser-archive-plugins/tree/main/deb) for a more complete example.
