package main
import (
    "github.com/binarydud/jumphash/server"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/tylerb/graceful"

    "flag"
    "log"
)

// main runs the jumphash program.
// There are two optional arguments, addr and seconds.
func main() {
    var addr = flag.String("address", "127.0.0.1:8080", "tcp address to listen on")
    var seconds = flag.Int("sleep", 5, "time in seconds to sleep each POST request")
    flag.Parse()
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(server.Shutdown())
    e.Use(server.Sleep(*seconds))
    e.Post("/", server.Hash)
    log.Printf("listening on %s", *addr)
    graceful.ListenAndServe(e.Server(*addr), 0)
}
