# lstn

This is a small Go program that listens on a port, and responds with whatever you give it via the flags.

```sh
$ lstn -h
Usage of lstn:
  -body string
			The http response body.
  -path string
			The request path to respond to. (default "/")
  -port string
			The port to listen on. (default "3000")
  -status int
			The http status to return. (default 200)

$ lstn --port 5000 --status 500 --body '{"foo":"bar"}' --path /
```

It might be useful for testing things.
