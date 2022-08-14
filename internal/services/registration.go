package services

import (
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

var AppManagement map[string]interface{}

func init() {
	if AppManagement == nil {
		AppManagement = make(map[string]interface{})
		logrus.Info("Init new app management")
	}
}

func RegisterApp(name string, app interface{}) {
	mx := sync.RWMutex{}
	mx.Lock()
	defer mx.Unlock()
	name = strings.ToLower(name)
	AppManagement[name] = app
}
