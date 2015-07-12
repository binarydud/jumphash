package server

import (
    "github.com/labstack/echo"
    "github.com/stretchr/testify/assert"
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
    "os"
    "os/signal"
    "time"
)

func TestHashPass(t *testing.T){
    assert.Equal(t, "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==", hashPass("angryMonkey"), "should be equal")
}

func TestEmptyPassword(t *testing.T){
    assert.Equal(t, hashPass(""), "", "should be empty")
}

func TestHash(t *testing.T){
    e := echo.New()
    req, _ := http.NewRequest(echo.POST, "/", strings.NewReader("password=angryMonkey"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")
    rec := httptest.NewRecorder()
    c := echo.NewContext(req, echo.NewResponse(rec), e)
    Hash(c)
    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==", rec.Body.String())
}

func TestShutdown(t *testing.T){
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, os.Interrupt, os.Kill)
    e := echo.New()
    req, _ := http.NewRequest(echo.POST, "/", strings.NewReader("command=shutdown"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")
    rec := httptest.NewRecorder()
    c := echo.NewContext(req, echo.NewResponse(rec), e)
    Shutdown()(c)
    s := <-sigchan
    assert.Equal(t, os.Interrupt, s)

}

func TestSleep(t *testing.T) {
    req, _ := http.NewRequest(echo.POST, "/", nil)
    rec := httptest.NewRecorder()
    c := echo.NewContext(req, echo.NewResponse(rec), echo.New())
    start := time.Now()
    Sleep(1)(c)
    duration := time.Now().Sub(start)
    assert.True(t, duration > time.Second*1)


    req, _ = http.NewRequest(echo.GET, "/", nil)
    rec = httptest.NewRecorder()
    c = echo.NewContext(req, echo.NewResponse(rec), echo.New())
    start = time.Now()
    Sleep(1)(c)
    duration = time.Now().Sub(start)
    assert.True(t, duration < 1*time.Second)
}
