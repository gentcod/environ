# environ
This is a simple Go package that helps you load environment configurations from your .env file

#

### Usage:
- Create your .env file: `touch .env`:
~~~
  PORT=5000
  URL=hello.com
  KEY=secret
~~~
- Import the module:
~~~
import 	environ "github.com/gentcod/environ"
~~~
- Create your Config struct based on your environment variables:
~~~
type Config struct {
	Port string `mapstructure:"PORT"`
	Url string `mapstructure:"URL"`
	Key string `mapstructure:"KEY"`
}
~~~
- Use the `Init` func to initialize and get your environment variables set:
~~~
filepath := "./env/test.env"
var config Config

err := environ.Init(filepath, &config)
if err != nil {
  log.Fatal("error laoding configurations")
}
~~~
