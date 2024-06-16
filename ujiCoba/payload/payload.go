package payload

import "errors"

//====================================================================================

type Data1 struct {
	UuId        string `form:"uuid" json:"uuid" bson:"uuid"`
	Name        string `form:"kodename" json:"kodename" bson:"kodename"`
	Childofuuid string `form:"childofuuid" json:"childofuuid" bson:"childofuuid"`
}

func (model *Data1) ConvertToMapInterface() interface{} {
	dataMap := map[string]interface{}{
		"uuId":        model.UuId,
		"kodename":    model.Name,
		"childofuuid": model.Childofuuid,
	}
	return dataMap
}
func (model *Data1) ConvertToStruct(data interface{}) error {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("conversion failed")
	}
	dataStruct := &Data1{
		UuId:        dataMap["uuid"].(string),
		Name:        dataMap["name"].(string),
		Childofuuid: dataMap["childofuuid"].(string),
	}
	model = dataStruct
	return nil
}

//====================================================================================

type Data2 struct {
	UuId        string `form:"uuid" json:"uuid" bson:"uuid"`
	Kodename    string `form:"kodename" json:"kodename" bson:"kodename"`
	Childofuuid string `form:"childofuuid" json:"childofuuid" bson:"childofuuid"`
}

func (model *Data2) ConvertToMapInterface() interface{} {
	dataMap := map[string]interface{}{
		"uuId":        model.UuId,
		"name":        model.Kodename,
		"childofuuid": model.Childofuuid,
	}
	return dataMap
}
func (model *Data2) ConvertToStruct(data interface{}) error {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("conversion failed")
	}
	dataStruct := &Data2{
		UuId:        dataMap["uuid"].(string),
		Kodename:    dataMap["kodename"].(string),
		Childofuuid: dataMap["childofuuid"].(string),
	}
	model = dataStruct
	return nil
}

//====================================================================================
