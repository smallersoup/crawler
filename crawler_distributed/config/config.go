package config

const (
	//Saver ES index
	Index = "dating_profile"

	//Service Saver port
	ItemSaverServicePort = 1234

	//rating limit
	Qps = 20

	//WorkerPort0
	WorkerPort0 = 9000
	//Saver ES Endpoints
	EsUrl = "http://192.168.1.101:9200"

	//RPC Service EndPoints
	ItemSaverRpc = "ItemSaveService.Save"

	//Crawler Worker ServiceRpc
	CrawlerServiceRpc = "CrawlerService.Process"

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"
)
