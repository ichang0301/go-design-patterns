package barrier

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBarrier(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/headers":
			if timeoutMilliseconds == 1 {
				time.Sleep(10 * time.Millisecond) // induce timeout
			}
			w.Write([]byte("Accept-Encoding: gzip"))
		case "/user-agent":
			if timeoutMilliseconds == 1 {
				time.Sleep(10 * time.Millisecond) // induce timeout
			}
			w.Write([]byte("user-agent: test"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{
			mockServer.URL + "/headers",
			mockServer.URL + "/user-agent",
		}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "user-agent") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{
			"http://malformed-url",
			mockServer.URL + "/user-agent",
		}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{
			mockServer.URL + "/headers",
			mockServer.URL + "/user-agent",
		}

		timeoutMilliseconds = 1

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		return ""
	}

	originalStdout := os.Stdout
	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	os.Stdout = originalStdout

	temp := <-out
	return temp
}
