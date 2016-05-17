# path-param-server

This is an example of how to extract parameters from the path
using the Go standard library.

Mostly, it is a convenient testbed for me to get some Github
Webhooks working on a Jenkins build server, but that's not
too meaningful to anyone reading this, although one thing that
I can say is that the Jenkins Github plugin does not seem to
like parameterized repository names, which I was trying to use
to allow the Jenkins job to build from any of a team's remote
repositories.

## running

```sh
go run main.go
```

or

```sh
go install
path-param-server
```

## options

Set the listening port with the `PORT` environment variable, or
set the entire bind address (e.g,. "0.0.0.0:5000") with the
`ADDR` environment variable.
