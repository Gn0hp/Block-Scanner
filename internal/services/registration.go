package services

import (
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

//AppManagement use remote config like consul to manage service. (It would be better to access a registered app)
var AppManagement map[string]interface{}

func init() {
	if AppManagement == nil {
		AppManagement = make(map[string]interface{})
		logrus.Info("Init new app management")
	}
}

func RegisterApp(name string, app interface{}) {
	mux := sync.RWMutex{}
	mux.Lock()
	defer mux.Unlock()
	name = strings.ToLower(name)
	AppManagement[name] = app
}
