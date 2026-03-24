package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/sony-nurdianto/test_driven_development/chapter05/db"
	"github.com/sony-nurdianto/test_driven_development/chapter05/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexIntegration(t *testing.T) {
	// arange
	book := db.Book{
		ID:     uuid.New().String(),
		Name:   "My first integration test",
		Status: db.Available.String(),
	}

	bs := db.NewBookService([]db.Book{book}, nil)
	h := handler.NewHandler(bs, nil)

	svr := httptest.NewServer(http.HandlerFunc(h.Index))
	defer svr.Close()

	// act
	r, err := http.Get(svr.URL)

	// assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	require.Nil(t, err)
	var resp handler.Response
	err = json.Unmarshal(body, &resp)
	require.Nil(t, err)
	assert.Equal(t, 1, len(resp.Books))
	assert.Contains(t, resp.Books, book)
}
