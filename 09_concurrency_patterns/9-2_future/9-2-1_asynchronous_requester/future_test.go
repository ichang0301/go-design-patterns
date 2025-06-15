package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success_result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello, World!", nil
		})

		wg.Wait()
	})

	t.Run("Error_result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e)
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("an error occurred")
		})

		wg.Wait()
	})

	t.Run("Closure_Success_result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future := &MaybeString{}
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(setContext("Hello"))

		wg.Wait()
	})
}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	t.Helper()

	time.Sleep(2 * time.Second)

	t.Error("Test timed out")
	wg.Done()
}
