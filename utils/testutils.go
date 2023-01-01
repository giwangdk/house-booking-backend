package utils

import (
	"encoding/json"
	"final-project-backend/server"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeReq(opts *server.RouterConfig, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
	router := server.CreateRouter(opts)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return router, rec
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}
