[![Tests on Linux, MacOS and Windows](https://github.com/gohugoio/hugoreleaser-plugins-api/workflows/Test/badge.svg)](https://github.com/gohugoio/hugoreleaser-plugins-api/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/github.com/gohugoio/hugoreleaser-plugins-api)](https://goreportcard.com/report/github.com/gohugoio/hugoreleaser-plugins-api)
[![GoDoc](https://godoc.org/github.com/gohugoio/hugoreleaser-plugins-api?status.svg)](https://godoc.org/github.com/gohugoio/hugoreleaser-plugins-api)


Plugins API for https://github.com/gohugoio/hugoreleaser

A plugin is a [Go Module](https://go.dev/blog/using-go-modules) with a main func, e.g., using the API provided by the [archiveplugin](https://pkg.go.dev/github.com/gohugoio/hugoreleaser-plugins-api/archiveplugin) package:

```go
func main() {
	var archiveClient archiveClient
	server, err := server.New(
		server.Options[model.Config, archiveplugin.Request, any, model.Receipt]{
			Init: func(c model.Config, prococol execrpc.ProtocolInfo) error {
				archiveClient.cfg = c
				return nil
			},
			Handle: func(call *execrpc.Call[archiveplugin.Request, any, model.Receipt]) {
				model.Infof(call, "Creating archive %s", call.Request.OutFilename)
				var receipt model.Receipt
				if !archiveClient.cfg.Try {
					if err := archiveClient.createArchive(call.Request); err != nil {
						receipt.Error = model.NewError(name, err)
					}
				}
				receipt = <-call.Receipt()
				call.Close(false, receipt)
			},
		},
	)
	if err != nil {
		log.Fatalf("Failed to create server: %s", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

```

See the [Deb Plugin](https://github.com/gohugoio/hugoreleaser-archive-plugins/tree/main/deb) for a more complete example.
