package setting

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Application struct {
	Evn string `yaml:"env"`
}

type BaseS struct {
	RunMode string `yaml:"run_mode"`
}

type AppS struct {
	PageSize  int    `yaml:"page_size"`
	JwtSecret string `yaml:"jwt_secret"`
}

type ServerS struct {
	HttpPort     int           `yaml:"http_port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type DatabaseS struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table_prefix"`
}

var (
	confData []byte
	Base     BaseS
	App      AppS
	Server   ServerS
	Database DatabaseS
)

func init() {

	var err error

	str, _ := os.Getwd()

	confData, err := ioutil.ReadFile(str + "/conf/application.yaml")
	if err != nil {
		log.Fatalf("Fatal to parse '/conf/application.yaml' : %v", err)
	}
	application := Application{}
	yaml.Unmarshal(confData, &application)

	if application.Evn == "" {
		log.Fatalf("Please configure the shipping environment in '/conf/application.yaml'")
	}

	confData, err = ioutil.ReadFile(str + "/conf/conf_" + application.Evn + ".yaml")
	if err != nil {
		log.Fatalf("Fatal to parse '/conf/conf_%v.yaml'", application.Evn)
	}

	Base = LoadBase(confData)
	App = LoadApp(confData)
	Server = LoadServer(confData)
	Database = LoadDatabase(confData)

}

func LoadBase(data []byte) BaseS {
	base := BaseS{}
	yaml.Unmarshal(data, &base)
	return base
}

func LoadApp(data []byte) AppS {
	app := AppS{}
	err := yaml.Unmarshal(data, &app)
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	return app
}

func LoadServer(data []byte) ServerS {
	server := ServerS{}
	err := yaml.Unmarshal(data, &server)
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	return server
}

func LoadDatabase(data []byte) DatabaseS {
	database := DatabaseS{}
	err := yaml.Unmarshal(data, &database)
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	return database
}
