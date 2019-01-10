/* copy from github.com/shen100/golang123 , thanks  */
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"unicode/utf8"
)

var jsonData map[string]interface{}

func initJSON(conf string) {
	if conf == "" {
		conf = "./config.json"
	}
	bytes, err := ioutil.ReadFile(conf)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type dBConfig struct {
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	URL          string
}

// DBConfig 数据库相关配置
var DBConfig dBConfig

func initDB() {
	SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}

//
type serverConfig struct {
	SiteName       string
	LogDir         string
	LogFile        string
	Port           int
	LogMaxLines    int
	LogMaxSize     int
	LogDaily       bool
	LogMaxDays     int
	LogEnableDepth bool
	LogDepth       int
	LogLevel       string
	TokenSecret    string
	TokenMaxAge    int
	PassSalt       string
	DcorpSecret    string
	DcorpId        string
	DacToken       string
	DchatId        string
	ResouLevel     float64
}

// ServerConfig 服务器相关配置
var ServerConfig serverConfig

func initServer() {
	SetStructByJSON(&ServerConfig, jsonData["server"].(map[string]interface{}))
	sep := string(os.PathSeparator)
	execPath, _ := os.Getwd()
	length := utf8.RuneCountInString(execPath)
	lastChar := execPath[length-1:]
	if lastChar != sep {
		execPath = execPath + sep
	}

	//	ymdStr := time.Now().Format("2006-01-02")

	if ServerConfig.LogDir == "" {
		ServerConfig.LogDir = execPath
	} else {
		length := utf8.RuneCountInString(ServerConfig.LogDir)
		lastChar := ServerConfig.LogDir[length-1:]
		if lastChar != sep {
			ServerConfig.LogDir = ServerConfig.LogDir + sep
		}
	}
	// 	fmt.Println(ServerConfig.LogDir)
	ServerConfig.LogFile = ServerConfig.LogDir + ServerConfig.LogFile
}

func ConfigInit(conf string) {
	initJSON(conf)
	initDB()
	initServer()
}
