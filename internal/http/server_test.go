package http

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type (
	ReaderMock struct{}
)

func (reader ReaderMock) ReadLine(lineNumber int) (string, error) {
	if lineNumber <= 0 {
		return "", errors.New("Line number must be a postive number")
	}
	if lineNumber >= 10 {
		return "", errors.New("Line numbers exceeds the number of lines of this file")
	}

	if lineNumber == 1 {
		return "kožušček", nil
	}
	return "Some line", nil
}

// TestHanderShouldReturnNotFoundForNonGetMethod
// Given an non get request
// When  handler
// Then  should return 404 status code
func TestHanderShouldReturnNotFoundForNonGetMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/lines/2", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	if w.Result().StatusCode != 404 {
		t.Error("Non Method posts should return 404 not found")
	}
}

// TestHanderShouldReturnNotFoundForNonValidNumbers
// Given an invalid number
// When  handler
// Then  should return 404 status code
func TestHanderShouldReturnNotFoundForNonValidNumbers(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/lines/soup", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	if w.Result().StatusCode != 404 {
		t.Error("Non valid number should return 404 not found")
	}
}

// TestHanderShouldReturnNotFoundForNegativeNumbers
// Given an non positive number
// When  handler
// Then  should return 404 status code
func TestHanderShouldReturnNotFoundForNegativeNumbers(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/lines/0", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	if w.Result().StatusCode != 404 {
		t.Error("Non postive number should return 404 not found")
	}
}

// TestHanderShouldReturnALineWhenTheNumberIsValidAndExisting
// Given an positive number present on the reader target
// When  handler
// Then  should output a string line
func TestHanderShouldReturnALineWhenTheNumberIsValidAndExisting(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/lines/6", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if string(body) != "Some line" {
		t.Error("Couldn't read a line from a valid number")
	}
}

// TestHanderShouldReturnStatusOKWhenTheNumberIsValidAndExisting
// Given an positive number present on the reader target
// When  handler
// Then  should return status 200
func TestHanderShouldReturnStatusOKWhenTheNumberIsValidAndExisting(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/lines/6", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Error("Didn't return status code 200 from a valid number")
	}
}

// TestHanderShouldReturnStatus413ForLineNumberAboveTheExistingOnes
// Given an positive number not present on the reader target
// When  handler
// Then  should return status 413
func TestHanderShouldReturnStatus413ForLineNumberAboveTheExistingOnes(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/lines/11", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	resp := w.Result()

	if resp.StatusCode != 413 {
		t.Error("Didn't return status code 413 for a number bigger than the number of lines")
	}
}

// TestHanderShouldReturnValidASCII
// Given an positive number not present on the reader target
// When  handler
// Then  should return status 413
func TestHanderShouldReturnValidASCII(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/lines/1", nil)
	w := httptest.NewRecorder()
	server := New(ReaderMock{})
	server.handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if string(body) != "kozuscek" {
		t.Error("Couldn't read a line from a valid number")
	}
}
