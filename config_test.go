package main

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var configString = `
app_id: foobar2
user_id: myuserid
key: mykeyid1
output_file: ~/mygitcrypt.key
vault_address: "https://127.0.0.1:8200"
    `

func TestConfigLocation(t *testing.T) {

	Convey("Test Config Location", t, func() {
		So(userHome()+"/.lockpick", ShouldEqual, configLocation())
	})

	Convey("Test Config Location set with env", t, func() {
		os.Setenv("LOCKPICK_CONF", "foobar")
		So("foobar", ShouldEqual, configLocation())
		os.Setenv("LOCKPICK_CONF", "")
	})
}

func TestConfig(t *testing.T) {

	Convey("Test Reading Config", t, func() {
		configFile, err := ioutil.TempFile("", "lockpick")
		So(err, ShouldEqual, nil)
		configFile.WriteString(configString)
		os.Setenv("LOCKPICK_CONF", configFile.Name())
		readConfig()
		So(config.AppID, ShouldEqual, "foobar2")
		So(config.UserID, ShouldEqual, "myuserid")
		So(config.Key, ShouldEqual, "mykeyid1")
		So(config.OutputFile, ShouldEqual, "~/mygitcrypt.key")
		So(config.VaultAddress, ShouldEqual, "https://127.0.0.1:8200")

	})
}
