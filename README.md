# wf-migrate

[![GoDoc](https://pkg.go.dev/badge/github.com/writefreely/wf-migrate?status.svg)](https://pkg.go.dev/github.com/writefreely/wf-migrate)

wf-migrate provides helper functions and a command-line utility for migrating posts between [WriteFreely](https://writefreely.org) instances.

## Command-line

Install the command-line utility with:

```
go install github.com/writefreely/wf-migrate/cmd/wfimport
```

`wfimport` takes a username `-u`, optional WriteFreely instance hostname `-h`, and the filename of the JSON data you want to import.

By default, `wfimport` publishes posts to [Write.as](https://write.as):

```
wfimport -u username exported-data.json
```

But you can also supply another WriteFreely instance to import to:

```
wfimport -u username -h pencil.writefree.ly exported-data.json
```
