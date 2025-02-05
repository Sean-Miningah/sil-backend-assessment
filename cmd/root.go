package cmd

var config util.Config

var rootCmd = &cobra.Command{
	Use: "server",
	Short: "Sil Backend Assessment",
	Long: "Backend Assesment Application"
}

func init() {
	var err error

	config, err = util.LoadConfig(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to laod env config: %s", err)
		os.Exit(1)
	}
}

func Execute() {
	return rootCmd.Execute()
}