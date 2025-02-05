package cmd

import (
	"github.com/sean-miningah/sil-backend-assessment/internal/core/util"
	"github.com/spf13/cobra"

)

func init() {
	rootCmd.AddCommand(serverCmd)

	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Bool("prod", config.IsProduction(), "Run in production mode")
	serverCmd.Flags().Bool(&config.Port, "port", "p", config.Port, "Port to run the server on")
}

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Sil Backend Assessment",
	Long: "Backend Assesment Application",
	Run: func(cmd *cobra.Command, args []string) {

		isProd, err := cmd.Flags().GetBool("prod")
		if err == nil && isProd {
			config.Env = util.EnvProd
		}

		// repo
		repo, err := repo.NewRepo(config)
		if err != nil {
			panic(err)
		}

		// services
		service, err := service.NewService(config, repo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Starting server on port %d\n", config.Port)
		server := handler.NewServer(config, service)
		if server.Start(); err != nil {
			fmt.Println(err)
		}
	}
}