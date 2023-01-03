package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	NOTE   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func TestGetAllExpense(t *testing.T) {
	seedExpense(t)
	var e []expense

	res := request(http.MethodGet, uri("expenses"), nil)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(e), 0)
}
func TestCreateCustomer(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
	    "amount": 79,
	   	"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)
	var exp expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&exp)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, exp.ID)
	assert.Equal(t, "strawberry smoothie", exp.Title)
	assert.Equal(t, 79.00, exp.Amount)
	assert.Equal(t, "night market promotion discount 10 bath", exp.NOTE)
	assert.Equal(t, []string{"food", "beverage"}, exp.Tags)
}

func TestGetCustomerByID(t *testing.T) {
	e := seedExpense(t)

	var lastExp expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(e.ID)), nil)
	err := res.Decode(&lastExp)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, e.ID, lastExp.ID)
	assert.NotEmpty(t, lastExp.Title)
	assert.NotEmpty(t, lastExp.NOTE)
	assert.NotEmpty(t, lastExp.Tags)
	assert.NotEmpty(t, lastExp.Amount)
}

func seedExpense(t *testing.T) expense {
	var c expense
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create customer:", err)
	}
	return c
}
func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	token := os.Getenv("AUTH_TOKEN")
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
