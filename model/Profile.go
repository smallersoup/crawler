package model

import "encoding/json"

type Profile struct {
	//姓名
	Name string
	//性别
	Gender string
	//年龄
	Age int
	//身高
	Height int
	//体重
	Weight int
	//收入
	Income string
	//婚姻状况
	Marriage string
	//学历
	Education    string
	WorkLocation string
	//职业
	Occupation string
	//户口
	Hukou string
	//星座
	Xinzuo string
	//房
	House string
	//车
	Car string
}

func FromJsonObj(obj interface{}) (Profile, error) {

	objJson, err := json.Marshal(obj)

	profile := Profile{}
	if err != nil {
		return profile, nil
	}
	return profile, json.Unmarshal(objJson, &profile)
}
