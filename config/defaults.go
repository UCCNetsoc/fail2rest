package config

import (
	"github.com/spf13/viper"
)

func loadDefaults() {
	// Fail2Rest
	viper.SetDefault("fail2rest.production", false)
	viper.SetDefault("fail2rest.auth_enabled", true)
	viper.SetDefault("fail2rest.secret", "abcd1234")

	// HTTP
	viper.SetDefault("http.port", 8080)

	// Fail2Ban
	viper.SetDefault("fail2ban.socket", "/var/run/fail2ban/fail2ban.sock")

	// Consul
	viper.SetDefault("consul.enabled", true)
	viper.SetDefault("consul.host", "netsoc-consul:8500")
	viper.SetDefault("consul.token", "")
	viper.SetDefault("consul.fail2rest", "fail2rest-token")
}
