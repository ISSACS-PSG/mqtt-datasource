package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/ISSACS-PSG/mqtt-datasource/pkg/plugin"
)

func main() {
	if err := datasource.Manage("issacs-mqtt-datasource", plugin.NewMQTTInstance, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}