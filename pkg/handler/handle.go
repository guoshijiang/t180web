package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/1851616111/util/message"
	"github.com/golang/glog"
	"t180web/pkg/model"
)

func GetTeamPersonInfo(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	PersonInfo, err := model.TeamPersonInfo()
	if err != nil {
		glog.Errorf("Get Team Person Info Error")
		return
	}
	if err := json.NewEncoder(w).Encode(PersonInfo); err != nil {
		message.InnerError(w)
		return
	}
	return
}

