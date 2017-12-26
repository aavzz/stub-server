package cfgfile

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"errors"
	"github.com/go-ini/ini"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)


type section1 struct {
	field1 string
	field2 string
}

// information is copied to this struct and is publicly available
type Section1 struct {
	Field1 string
	Field2 string
}

type configurationFile struct {
	sect1 section1
}

// information is copied to this struct and is publicly available
type ConfigurationFile struct {
	Sect1 Section1
}


var cfgFile1, cfgFile2 configurationFile
var cfgFile1ok, cfgFile2ok bool


func ReadConfig() {

	cfg, err := ini.Load(ConfigCfgFile())

	if err != nil {
		daemonState := os.Getenv("_STUB_SERVER_DAEMON_STATE")
		if daemonState == "" {
			SysLog.Err(err.Error())
			RemovePidFile()
		}
		os.Exit(1)
	}

	cfgFile1ok=false
	_, err = cfg.GetSection("section1")
	if err == nil {
		if cfg.Section("section1").HasKey("field1") {
			cfgFile1.sect1.field1 = cfg.Section("section1").Key("field1").String()
		}
		if cfg.Section("section1").HasKey("field2") {
			cfgFile1.sect1.field2 = cfg.Section("section1").Key("field2").String()
		}
	}
	// true only means that we finished updating config data, not that the data is ok
	cfgFile1ok=true
	
	// update second copy of config data
	cfgFile2ok=false
	if err == nil {
		if cfg.Section("section1").HasKey("field1") {
			cfgFile2.sect1.field1 = cfg.Section("section1).Key("field1").String()
		}
		if cfg.Section("section1").HasKey("field2") {
			cfgFile2.sect1.field2 = cfg.Section("section1").Key("field2").String()
		}
	}
	cfgFile2ok=true
}

func CfgFileContent() (*CfgFile, error) {
	if cfgFile1ok == true {
		c := &CfgFile{
			Sect1: Section1{
				Field1: cfgFile1.sect1.field1,
				Field2: cfgFile1.sect1.field2,
			},
		}
		return c, nil
	}
	if cfgFile2ok == true {
		c := &CfgFile{
			Sect1: section1{
				Field1: cfgFile2.sect1.field1,
				Field2: cfgFile2.sect1.field1,
			},
		}
		return c, nil
	}
	return nil, errors.New("Error retrieving configuration file content")
}

