/*
Package cfgfile parses configuration file for the project
and creates internal representation of the parsed file.

If the configuration file cannot be read, the whole process terminates.

Internal representation is meant to be updated on signal (e.g. SIGHUP)
and is exported read-only in a consistent way.
*/
package cfgfile

import (
	"errors"
	. "github.com/aavzz/stub-server/setup/cmdlnopts"
	. "github.com/aavzz/stub-server/setup/pidfile"
	. "github.com/aavzz/stub-server/setup/syslog"
	"github.com/go-ini/ini"
	"os"
)

// section1 represents a section in the configuration file.
// There may be several types representing different sections
// each one named after the section it represents.
// This is a private version, information is copied here from
// the configuration file and is exported to its public counterpart.
type section1 struct {
	//field1 represents a key in the configuration file.
	//Its name and type should reflect the nature of the key.
	field1 string
	//field2 is another key (see field1)
	field2 string
}

// Section1 represents a section in the configuration file.
// There may be several types representing different sections
// each one named after the section it represents.
// This is a public version, information is copied here from
// its private counterpart.
type Section1 struct {
	//Field1 represents a key in the configuration file.
	//Its name and type should reflect the nature of the key.
	Field1 string
	//Field2 is another key (see Field1)
	Field2 string
}

// configurationFile represents configuration file.
// This is a private version, information is copied here from
// configuration file and is exported to its public counterpart.
type configurationFile struct {
	// sect1 represents a section in the configuration file.
	sect1 section1
}

// ConfigurationFile represents configuration file.
// This is a public version, information is copied here from its
// its public counterpart.
type ConfigurationFile struct {
	// Sect1 represents a section in the configuration file.
	Sect1 Section1
}

// Private structures and guardian vars.
var cfgFile1, cfgFile2 configurationFile
var cfgFile1ok, cfgFile2ok bool

// ReadConfig reads configuration file on startup or signal and fills
// private structures with the information it finds. In case of failure
// it logs a message and stops the process. Logging must have been
// initialized when ReadConfig is called.
func ReadConfig() {

	cfg, err := ini.Load(ConfigCfgFile())

	if err != nil {
		SysLog.Err(err.Error())
		RemovePidFile()
		os.Exit(1)
	}

	// update the first copy of config data
	cfgFile1ok = false
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
	cfgFile1ok = true

	// update the second copy of config data
	cfgFile2ok = false
	if err == nil {
		if cfg.Section("section1").HasKey("field1") {
			cfgFile2.sect1.field1 = cfg.Section("section1").Key("field1").String()
		}
		if cfg.Section("section1").HasKey("field2") {
			cfgFile2.sect1.field2 = cfg.Section("section1").Key("field2").String()
		}
	}
	// true only means that we finished updating config data, not that the data is ok
	cfgFile2ok = true
}

// CfgFileContent exports internal representation of the configuration file
// in a safe and consistent way.
func CfgFileContent() (*ConfigurationFile, error) {
	if cfgFile1ok == true {
		c := &ConfigurationFile{
			Sect1: Section1{
				Field1: cfgFile1.sect1.field1,
				Field2: cfgFile1.sect1.field2,
			},
		}
		return c, nil
	}
	if cfgFile2ok == true {
		c := &ConfigurationFile{
			Sect1: Section1{
				Field1: cfgFile2.sect1.field1,
				Field2: cfgFile2.sect1.field1,
			},
		}
		return c, nil
	}
	return nil, errors.New("Error retrieving configuration file content")
}
