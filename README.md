jsonconfig
==============
Package Name: jsonconfig
Author      : jhoelzel
URK         : github.com/jhoelzel/jsonconfig
License     : MIT
Version     : 0.0.1 

Summary
--------------

jsonconfig is a simple golang package that reads a simple JSON configuration file from any path into a strongly typed struct and vice versa.

Why a JSON config file
--------------

Most of the time changing minor settings in your application is paramount, but most of the time they cause downtime. 
jsonconfig enables you to adapt your configuration however you like by simply reloading the configuration as needed.
Furthermore, using the JSON Format, enables you to not only change settings yourself with a simple text editor, but also have machines to the jobs for you ;)

Usage
--------------

*in your code *

```golang
import (
	"github.com/jhoelzel/jsonconfig"
)

var config jsonconfig.Configuration

func main() {
    config = jsonconfig.Read("./config/config.json")
}
```

Methods
--------------

| Command | Description |
| --- | --- |
| jsonconfig.Read("./config/config.json") | Read the json file from "./config/config.json" and return Configuration |
| jsonconfig.Write(Configuration, "./config/config.json") | Write the Configuration to "./config/config.json" |
| jsonconfig.Configuration | The "Configuration" Type |

The configuration File
--------------

*HowEverYouChooseToNameIt.json*

```json
{
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

}
```

But your configuration flags suck!
--------------

While that topic is up for debate, this repository has been created to quickly build a new Package using Golang that works for me.
Your needs may be different, but do not despair! You may simply hit the fork button and create your own version.

To do that change the following files:

| models.go            | main_test.go                                 |
| -------------        | ------------- 
| add new Models here  | make sure your new Models are also tested    |


I have a better idea for X
--------------

Pull requests are allways welcome =)