# Generator Sample

Install the cmd tool and run it:

```bash
go install github.com/sourcefellows/mongo-query/cmd/mongo-query-gen@latest
```

Run it either by ruinning on command line...

```bash
mongo-query-gen -in Types.go -outDir .
```

.. or use `go generate`

```bash
go generate ./...
```