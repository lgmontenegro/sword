package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	//cfgFilePath string

	rootCmd = &cobra.Command{
		Use:   "sword",
		Short: "Sword is an assessment to Sword Health Company",
		Long:  `Sword is an API developer practical exercise assessment to Sword Health development staff team`,
	}
)

func init(){
	//cobra.OnInitialize(configApp)

	//rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file (default is config.yml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

/*func configApp(){

}*/

func Execute() error{
	return rootCmd.Execute()
}