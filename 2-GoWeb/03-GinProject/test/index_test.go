package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"studygolang/2-GoWeb/03-GinProject/initRouter"
	"testing"
)

func TestIndexRouter(t *testing.T){
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req,_ := http.NewRequest(http.MethodGet,"/",nil)
	router.ServeHTTP(w,req)
	fmt.Println(w.Code,w.Body.String())
	assert.Equal(t,http.StatusOK,w.Code)
	assert.Equal(t,"hello gin",w.Body.String())
}