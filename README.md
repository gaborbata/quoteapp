# quoteapp
A responsive web application which displays a random quote from its embedded database.

Used libraries/frameworks:

* [Gin](https://github.com/gin-gonic/gin) - a HTTP web framework written in Go
* [Bolt](https://github.com/boltdb/bolt) - an embedded key/value database for Go
* [Milligram](https://github.com/milligram/milligram) - a minimalist CSS framework
* [PicoModal](https://github.com/Nycto/PicoModal) - a small, self-contained JavaScript modal library

## Docker support
An example `Dockerfile` is available to help building and running images containing the application.

Examples:
* Build image: `docker build -t quoteapp .`
* Run image: `docker run -i -t -p 3030:3030 quoteapp`
