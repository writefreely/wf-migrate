# wf-import

[![GoDoc](https://godoc.org/github.com/writeas/wf-import?status.svg)](https://godoc.org/github.com/writeas/wf-import)

wf-import provides helper functions and a command-line utility for migrating posts between [WriteFreely](https://writefreely.org) instances.

## Command-line

Install the command-line utility with:

```
go get github.com/writeas/wf-import/cmd/wfimport
```

`wfimport` takes a username `-u`, optional WriteFreely instance hostname `-h`, and the filename of the JSON data you want to import.

By default, `wfimport` publishes posts to Write.as:

```
wfimport -u username exported-data.json
```

But you can also supply another WriteFreely instance to import to:

```
wfimport -u username -h pencil.writefree.ly exported-data.json
```
