package main

import (
	"github.com/Strum355/log"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"github.com/strangeman/fail2go"

	"github.com/UCCNetsoc/fail2rest/api"
	"github.com/UCCNetsoc/fail2rest/config"
	"github.com/UCCNetsoc/fail2rest/services"

	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Load Config
	config.Load()

	// Initialise logger
	if viper.GetBool("fail2rest.production") {
		log.InitJSONLogger(&log.Config{})
	} else {
		log.InitSimpleLogger(&log.Config{})
	}

	// Print settings
	config.PrintSettings()

	// Initialise Fail2Ban connection
	log.WithFields(log.Fields{
		"fail2ban socket": viper.GetString("fail2ban.socket"),
	}).Info("Initialising fail2ban connection")

	conn := fail2go.Newfail2goConn(viper.GetString("fail2ban.socket"))

	// Start HTTP Server
	log.WithFields(log.Fields{
		"port": viper.GetInt("http.port"),
	}).Info("Initialising HTTP Server")

	r := chi.NewRouter()

	var err error
	// Register service with Consul
	if viper.GetBool("consul.enabled") {
		consul := services.ConsulService{
			ConsulHost:  viper.GetString("consul.host"),
			ConsulToken: viper.GetString("consul.token"),
			ServiceAddr: "127.0.0.1",
			Port:        viper.GetInt("http.port"),
			TTL:         time.Second * 5,
		}
		err = consul.Setup()
		if err != nil {
			log.WithError(err).Error("Could not setup Consul service")
			os.Exit(1)
		}
		err = consul.Register()
		if err != nil {
			log.WithError(err).Error("Could not register with Consul service")
			os.Exit(1)
		}
	}
	// Initialise API
	api := api.API{Fail2Conn: conn}
	api.Register(r)

	err = http.ListenAndServe(":"+fmt.Sprint(viper.GetInt("http.port")), r)
	if err != nil {
		log.WithError(err).Error("Error serving HTTP")
	}
}
