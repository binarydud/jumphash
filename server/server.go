// Package server provides functions for running a hash server.
package server

import (
    "github.com/labstack/echo"
    "net/http"

    "syscall"
    "log"
    "time"
    "crypto/sha512"
    "encoding/base64"
)

// Sleep function is middleware that causes the request to sleep for given number of seconds.
// Only POST requests sleep as GET requests are immediately returned a 404
func Sleep(seconds int) echo.HandlerFunc {
    return func(c *echo.Context) error {
        if c.Request().Method == "POST" {
            time.Sleep(time.Duration(seconds) * time.Second)
        }
        return nil
    }
}

// Shutown middleware handles http based shutdown.
// If "command" is included via a POST, then the current goroutine will continue, but all
// subsequent requests should fail. This relies on the graceful package to handle the signal call.
func Shutdown() echo.HandlerFunc {
    return func(c *echo.Context) error {
        command := c.Request().PostFormValue("command")
        if command == "shutdown" {
            log.Printf("commencing shutdown")
            syscall.Kill(syscall.Getpid(), syscall.SIGINT)
        }
        return nil
    }
}

// hashPash takes a password and returns a hash representation.
// The password is first hashed with sha512, then encoded in base64.
func hashPass(password string) string {
    if password == ""{
        return ""
    }
    sha_512 := sha512.New()
    sha_512.Write([]byte(password))
    return base64.StdEncoding.EncodeToString(sha_512.Sum(nil))
}

// Hash handler gets a POST form value and returns a hash.
func Hash(c *echo.Context) error {
    password := c.Request().PostFormValue("password")
    body := hashPass(password)
    return c.String(http.StatusOK, body)
}
