# lstn

This is a small Go program that listens on a port, and responds with whatever you give it via the flags.

```sh
$ lstn -h
Usage of lstn:
  -body string
			The http response body. Defaults to an empty string.
  -path string
			The request path to respond to. Defaults to /. (default "/")
  -port string
			The port to listen on. Listens on 3000 by default. (default "3000")
  -status int
			The http status to return. Returns 200 by default. (default 200)

$ lstn --port 5000 --status 500 --body '{"foo":"bar"}' --path /
```

It might be useful for testing things.
