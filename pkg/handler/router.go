package handler


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.GET("/api/teamPersonInfo/personInfo", httprouter.Handle(GetTeamPersonInfo))
	return r
}


