package cmd

import (
	"context"
	"log"
	"time"

	"github.com/sickyoon/daddyddns/ddns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "refresh DNS record",
	Long:  "refresh DNS record",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c := ddns.New("y00ns.com")
		err := c.Refresh(ctx, viper.GetString("subdomain"))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	refreshCmd.PersistentFlags().StringP("subdomain", "d", "dev", "Subdomain")
	viper.BindPFlag("subdomain", refreshCmd.PersistentFlags().Lookup("subdomain"))
	rootCmd.AddCommand(refreshCmd)
}
