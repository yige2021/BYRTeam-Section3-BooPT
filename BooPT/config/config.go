package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type DatabaseConf struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type LDAPConf struct {
	Addr     string
	User     string
	Password string
	BaseDN   string
}

type S3Conf struct {
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
}

type Config struct {
	Database DatabaseConf
	LDAP     LDAPConf
	S3       S3Conf
	LogLevel int // 0: error, 1: warn, 2: info, 3: debug
	JwtSalt  string
}

var CONFIG = Config{
	Database: DatabaseConf{
		Host:     "localhost",
		Port:     5432,
		User:     "BooPT",
		Password: "BooPTPassword",
		DbName:   "BooPT",
	},
	LDAP: LDAPConf{
		Addr:     "localhost:389",
		User:     "cn=Manager,dc=maxcrc,dc=com",
		Password: "secret",
		BaseDN:   "dc=maxcrc,dc=com",
	},
	S3: S3Conf{
		AccessKey: "",
		SecretKey: "",
		Region:    "bupt-selfhost-bucket",
		Bucket:    "boopt",
	},
	LogLevel: 3,
	JwtSalt:  "BooPTJWTSalt",
}

var JWTSALT = []byte(CONFIG.JwtSalt)

func Read(filename string) error {
	//声明变量yamlFile和err，赋值为文件名
	yamlFile, err := ioutil.ReadFile(filename)

	//未找到文件，返回对应信息
	if err != nil {
		logrus.Errorf("yamlFile.Get err %#v ", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &CONFIG)
	if err != nil {
		logrus.Errorf("Unmarshal: %#v ", err)
		return err
	}

	//读取成功
	return nil
}
