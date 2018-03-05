package model

import (
	"t180web/pkg/types"
	"fmt"
	"github.com/golang/glog"
)

func TeamPersonInfo()([]types.TeamPersonInfo, error) {
	var teamPersonInfo []types.TeamPersonInfo
	rows, herr := DB.Query(fmt.Sprintf(`select person_address, person_name, person_simg, person_content, person_gy, add_time from t180_person`))
	if herr != nil{
		glog.Errorf("TeamPersonInfo: sql return error %v\n", herr)
	}

	for rows.Next(){
		row := new(types.TeamPersonInfo)
		if herr = rows.Scan(&row.PersonAddr, &row.PersonName, &row.PersonImg, &row.PersonContent, &row.PersonGy, &row.AddTime); herr != nil{
			glog.Errorf("rows.Scan error %v\n", herr)
			return  nil, herr
		}
		teamPersonInfo = append(teamPersonInfo, *row)
	}

	if rows.Err() != nil {
		glog.Errorf("TeamPersonInfo: sql row.Err %v\n", herr)
		return nil, rows.Err()
	}
	return teamPersonInfo, nil
}
