package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sony-nurdianto/test_driven_development/chapter05/db"
	"github.com/sony-nurdianto/test_driven_development/chapter05/handler"
)

var _ = Describe("Handlers integrations", func() {
	var svr *httptest.Server
	var book db.Book

	BeforeEach(func() {
		book = db.Book{
			ID:     uuid.New().String(),
			Name:   "My first integration test",
			Status: db.Available.String(),
		}
		bs := db.NewBookService([]db.Book{book}, nil)
		ha := handler.NewHandler(bs, nil)
		svr = httptest.NewServer(http.HandlerFunc(ha.Index))
	})

	AfterEach(func() {
		svr.Close()
	})

	Describe("Index endpoint", func() {
		Context("with one existing book", func() {
			It("should return book", func() {
				r, err := http.Get(svr.URL)
				Expect(err).To(BeNil())
				Expect(r.StatusCode).To(Equal(http.StatusOK))

				body, err := io.ReadAll(r.Body)
				r.Body.Close()
				Expect(err).To(BeNil())

				var resp handler.Response
				err = json.Unmarshal(body, &resp)

				Expect(err).To(BeNil())
				Expect(len(resp.Books)).To(Equal(1))
				Expect(resp.Books).To(ContainElement(book))
			})
		})
	})
})
