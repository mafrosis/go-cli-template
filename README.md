A template for a Golang CLI app
==========

Fork and modify to quickly bootstrap a CLI app.

Features of this template:

* Uses the best-in-the-game [spf13/cobra](https://github.com/spf13/cobra) as the CLI library
* Supports config via environment variables or a YAML file in `~/.config`
* Includes a logging library for debug statements
* Includes a simple `make` build system
* Comes baked in with a sensible versioned release workflow


Getting Started
----------

1. Clone this repo
2. Run `make build`
3. Test with `./bin/go-cli-template version`
4. Test with `GO_CLI_TEMPLATE_EXAMPLE=foobar ./bin/go-cli-template demo`
5. Consider the code in `cmd/demo.go`, and `pkg/demo/demo.go`


How To Use
----------

0. Decide on a catchy name for your app
1. Use the template from github
2. Find and replace all references to the following in the source code:
  i. `mafrosis/go-cli-template`
  ii. `go-cli-template`
  iii. `GO_CLI_TEMPLATE`
3. Remove `cmd/demo.go`, and the directory `pkg/demo`
4. Start adding your own commands


### How to add a new command

Use the `cobra` library to add a new command (you will need `$GOPATH/bin` on your `$PATH` for this).

```
> cobra add foobar
foobar created at /Users/mblack/Development/me/go-cli-template
> ls cmd
> foobar.go  root.go  version.go
```


### Releasing a version

Git tags should be used for version release. The environment variable `RELEASE_TAG` is picked up
by the `Makefile` and injected into `go build`.

The following is a basic example. You will likely want to put this into a build & release pipeline
with Github Actions or other CI.

```
> git tag v1.0.0
> export RELEASE_TAG=v1.0.0
> make build
go build -v -tags netgo -ldflags '-s -w -extldflags "-static" -X "github.com/mafrosis/go-cli-template/pkg/version.String=1.0.0" -X "github.com/mafrosis/go-cli-template/pkg/version.Revision=180e376" -X "github.com/mafrosis/go-cli-template/pkg/version.Date=20230423"' -o bin/go-cli-template main.go
command-line-arguments
> ./bin/go-cli-template version
1.0.0 (180e376) built on 20230423
```
