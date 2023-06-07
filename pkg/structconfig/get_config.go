package structconfig

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var _env *Env // mutex sync.Mutex

// Get return Env struct contains all environment variable
func Get() *Env {
	if _env == nil {
		// mutex.Lock()
		// defer mutex.Unlock()
		if _env == nil {
			_env = loadEnv()
		}
	}

	return _env
}

func loadEnv() *Env {
	env := new(Env)

	dir := os.Getenv("DIR")

	if len(dir) == 0 {
		dir = "."
	}
	viper.SetConfigFile(fmt.Sprintf("%s/config.json", dir))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("an error happen when load config file: %v", err))
	}

	err := viper.Unmarshal(env)
	if err != nil {
		panic(fmt.Errorf("an error happen when unmarshal env config: %v", err))
	}

	return env
}
