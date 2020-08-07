// +build !skipForTest

package autoload

import (
	"github.com/eyolas/wellness-ws/config"
)

func init() {
	config.LoadConfig()
	config.InitializeLog()
	config.InitDb()
}
