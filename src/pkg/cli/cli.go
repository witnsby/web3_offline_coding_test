package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/helper"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/rps"
)

func init() {
	rootCmd.AddCommand(pvpCmd)
	rootCmd.AddCommand(pveCmd)
}

// rootCmd represents the General game information
var rootCmd = &cobra.Command{
	Use:   helper.Use,
	Short: helper.Short,
	Long:  helper.Long,
}

// pvpCmd represents the pvp command
var pvpCmd = &cobra.Command{
	Use:   helper.PvPUse,
	Short: helper.PvPShort,
	Run: func(cmd *cobra.Command, args []string) {
		rps.PlayPvP()
	},
}

// pveCmd represents the pve command
var pveCmd = &cobra.Command{
	Use:   helper.PvEUse,
	Short: helper.PvEShort,
	Run: func(cmd *cobra.Command, args []string) {
		rps.PlayPvBot()
	},
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
