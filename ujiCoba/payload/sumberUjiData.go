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
    {"uuid": "1", "data": {"name":"satu","status":"like"}, "conf":{"childofuuid": "0","index":"1","version":"1.00"}},
    {"uuid": "2", "data": {"name":"sutis","status":"like"}, "conf":{"childofuuid": "0","index":"2","version":"1.00"}},
    {"uuid": "3", "data": {"name":"sutarman","status":"like"}, "conf":{"childofuuid": "2","index":"1","version":"1.00"}},
    {"uuid": "4", "data": {"name":"sukiyam","status":"like"}, "conf":{"childofuuid": "1","index":"2","version":"1.00"}},
    {"uuid": "5", "data": {"name":"amir","status":"like"}, "conf":{"childofuuid": "1","index":"1","version":"1.00"}}
]`

//====================================================================================
