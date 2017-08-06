package jsonconfig

import (
	"bufio"
	"log"
	"os"
	"testing"
)

var debugconfigpath = "./jsconfig_testconfig.json"

// TestConfigRead tests if the default configuration can be read
func TestConfigRead(t *testing.T) {
	//create dummy file in root
	writeDummyConfiguration()
	//read it
	config := Read(debugconfigpath)
	//validate it
	validateConfiguration(t, config)
	//clean up
	teardown()
}

// TestConfigWrite tests if the configuration json is writeable and modifiable
func TestConfigWrite(t *testing.T) {
	//create dummy file in root
	writeDummyConfiguration()
	//read it
	config := Read(debugconfigpath)
	//validate it
	validateConfiguration(t, config)
	config.Admin.Email = "IamNowChangedForTheTest"
	Write(config, debugconfigpath)
	//read it again
	config = Read(debugconfigpath)
	//validate it
	validateConfigurationChanged(t, config)
	//clean up
	teardown()
}

// validateConfigurationChanged for the test
func validateConfigurationChanged(t *testing.T, config Configuration) {
	if config.Admin.Email != "IamNowChangedForTheTest" {
		t.Errorf("Admin.Email was incorrect!  got: %s, want: %s.", config.Admin.Email, "IamNowChangedForTheTest")
	}
}
func teardown() {
	var err = os.Remove(debugconfigpath)
	checkError(err)
}

func validateConfiguration(t *testing.T, config Configuration) {
	if config.Admin.Email != "admin@yourdomain.tld" {
		t.Errorf("Admin.Email was incorrect! got: %s, want: %s.", config.Admin.Email, "admin@yourdomain.tld")
	}
	if config.Globals.URL != "https://yourdomain.tld" {
		t.Errorf("Globals.URL was incorrect! got: %s, want: %s.", config.Admin.Email, "https://yourdomain.tld")
	}
	if config.Database.Username != "root" {
		t.Errorf("Database.Username was incorrect! got: %s, want: %s.", config.Admin.Email, "root")
	}
}

func writeDummyConfiguration() {
	testconf := `{
        "Globals": {
                "url": "https://yourdomain.tld",
				"theme": "default",
				"date_format": "m/d/Y",
				"time_format" : "h:ia",				
				"decimal_seperator" : ".",
				"language" : "en_us",
				"version" : "0.0.1"
        },
        "Admin": {
                "email": "admin@yourdomain.tld",
                "name": "Admin Adminicus"
        },
        "Emails": {
                "name": "no-reply Server",
                "host": "yourdomain.tld",
                "port": "587",
                "email": "no-reply@yourdomain.tld",
                "username": "no-reply@yourdomain.tld",
                "password": "...",
                "mechanism": "TLS"
        },
        "Login": {
                "username": "admin@yourdomain.tld",
                "password": "VeryUnsecureClearTextPW"
        },
        "Database": {
                "username": "root",
                "password": "root",
                "dbname": "testdb",
                "host": "mysql.yourdomain.tld",
                "port": "587"
        }

}`
	f, err := os.Create(debugconfigpath)
	checkError(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, writeerr := w.WriteString(testconf)
	checkError(writeerr)
	w.Flush()
}

//	facilitate error handling and logging
func checkError(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
