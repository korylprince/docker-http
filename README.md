# About

This is a simple program (with [Docker image](https://hub.docker.com/r/korylprince/docker-http/)) to serve the static contents of a folder. It can be configured to use HTTP Basic Auth with a supplied username and password.

# Usage

```
Usage of serve:
  -addr string
    	address to listen on, e.g. [<ip>]:<port> (default ":80")
  -dir string
    	folder to serve (default "/http")
  -passwd string
    	set password for basic auth
  -user string
    	set username for basic auth
```

Parameters can be set with flags, or environment variables: `HTTP_ADDRESS`, `HTTP_DIRECTORY`, `HTTP_USERNAME`, `HTTP_PASSWORD`. If `_FLAG` is appended to the environment variable (e.g. `HTTP_PASSWORD_FILE`) the parameter is set to the contents of the path given by the environment variable. Environment variables take precedence over flags.

If the username and password are set, HTTP Basic Auth will be used.
