package conf

import (
	. "github.com/aceberg/WMService/models"
	"github.com/spf13/viper"
)

const configPath = "/data/wmservice/config"

func GetConfig() (config Conf) {
	viper.SetDefault("DB_PATH", "/data/wmservice/rs1.db")
	viper.SetDefault("GUI_IP", "0.0.0.0")
	viper.SetDefault("GUI_PORT", "8843")
	viper.SetDefault("THEME", "cosmo")
	viper.SetDefault("SHOW", "15")

	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DbPath = viper.Get("DB_PATH").(string)
	config.GuiIP = viper.Get("GUI_IP").(string)
	config.GuiPort = viper.Get("GUI_PORT").(string)
	config.Theme = viper.Get("THEME").(string)
	config.Show = viper.Get("SHOW").(string)

	return config
}

func WriteConfig(theme string) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.Set("THEME", theme)
	viper.WriteConfig()
}
