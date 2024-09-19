# environ
This is a simple Go package that helps you load environment configurations from your .env file

#

### Usage:
- Create your .env file: `touch .env`:
~~~
PORT_ADDRESS=0.0.0.0:5000
GRPC_ADDRESS=0.0.0.0:3000
DB_DRIVER=postgres
DB_URL=postgres://root:secret@localhost:5431/dummy_bank?sslmode=disable
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
TOKEN_SECRET_KEY=12345678901234567890123456789012
TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=24h
MIGRATION_URL=file://sql/migrations
~~~

- Import the module:
~~~
import 	environ "github.com/gentcod/environ"
~~~

- Create your Config struct based on your environment variables:
~~~
type Config struct {
	PortAddress string
	GrpcAddress string
	DBDriver string
	DBUrl string
	TokenSymmetricKey string
	TokenSecretKey string
	TokenDuration time.Duration
	RefreshTokenDuration time.Duration
	MigrationUrl string
}
~~~
#### NB: Make sure that the struct field name matches the env key, in the form; field name: `FirstName`, env key: `FIRST_NAME`

- Use the `Init` func to initialize and get your environment variables set:
~~~
filepath := "./env/test.env"
var config Config

err := environ.Init(filepath, &config)
if err != nil {
  log.Fatal("error laoding configurations")
}
~~~
#### NB: filepath should be relative to the directory in which the file containing the calling module is located