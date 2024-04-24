package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

var ntpTimeFunc func(string) (time.Time, error)

func localMain() {
	time, err := ntpTimeFunc("0.beevik-ntp.pool.ntp.org")
	if err != nil {

		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(time)
	}
}

func TestLocalMainErrorHandling(t *testing.T) {
	old := os.Stderr // Сохраняем оригинальный os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w // Перенаправляем вывод в поток

	defer func() {
		w.Close()
		os.Stderr = old // Восстанавливаем оригинальный os.Stderr
	}()

	// Моделируем ситуацию, когда ntp.Time() возвращает ошибку
	ntpTimeFunc = func(address string) (time.Time, error) {
		return time.Time{}, errors.New("error message")
	}

	localMain()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	expected := "error message\n"
	fmt.Println(buf.String())
	if buf.String() != expected {
		t.Errorf("Expected error message to be %q, but got %q", expected, buf.String())
	}
}
