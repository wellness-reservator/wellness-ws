// +build !skipForTest

package autoload

import (
	"github.com/wellness-reservator/wellness-ws/config"
)

func init() {
	config.LoadConfig()
	config.InitializeLog()
	config.InitDb()
}
