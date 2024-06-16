package payload

//====================================================================================
var DataMap1 = []map[string]interface{}{
	{
		"uuId":        "1",
		"data":        `{"name": "satu", "status": "like"}`,
		"childofuuid": "0",
	},
	{
		"uuId":        "2",
		"data":        `{"name": "sutis", "status": "like"}`,
		"childofuuid": "0",
	},
	{
		"uuId":        "3",
		"data":        `{"name": "sutarman", "status": "like"}`,
		"childofuuid": "2",
	},
	{
		"uuId":        "4",
		"data":        `{"name": "sukiyam", "status": "like"}`,
		"childofuuid": "1",
	},
	{
		"uuId":        "5",
		"data":        `{"name": "amir", "status": "like"}`,
		"childofuuid": "1",
	},
}

//====================================================================================

var DataString1 = `[
	{"uuid": "1", "data": {"name":"satu","status":"like"}, "childofuuid": "0"},
	{"uuid": "2", "data": {"name":"sutis","status":"like"}, "childofuuid": "0"},
	{"uuid": "3", "data": {"name":"sutarman","status":"like"}, "childofuuid": "2"},
	{"uuid": "4", "data": {"name":"sukiyam","status":"like"}, "childofuuid": "1"},
	{"uuid": "5", "data": {"name":"amir","status":"like"}, "childofuuid": "1"}
]`

//====================================================================================
