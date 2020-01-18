package settings

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

//Settings is the struct that contains all the configs
type Settings struct {
	Mongo *MongoSettings `yaml:"mongo"`
}

//MongoSettings is the struct that contains the config for the db connection
type MongoSettings struct {
	Database   string `yaml:"database"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Port       string `yaml:"port"`
	Collection string `yaml:"collection"`
}

//FromYAML reads the configs from the YAML file
func FromYAML(file string) *Settings {
	filename, _ := filepath.Abs(file)

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	settings := new(Settings)

	err = yaml.Unmarshal(data, settings)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(settings)
	return settings
}

//BuildConnectionString takes the parameters and builds the db uri
func BuildConnectionString(settings *MongoSettings) string {
	// mongodb://<dbuser>:<dbpassword>@ds263808.mlab.com:63808/zipcode

	ConnectionString := fmt.Sprintf(
		"%s://<%s>:<%s>@ds263808.mlab.com:%s/%s",
		settings.Database,
		settings.User,
		settings.Password,
		settings.Port,
		settings.Collection,
	)
	fmt.Print(ConnectionString)
	return ConnectionString
}

//CreateConnectionString ...
func CreateConnectionString() string {

	settings := new(MongoSettings)
	ConnectionString := BuildConnectionString(settings)

	return ConnectionString
}
