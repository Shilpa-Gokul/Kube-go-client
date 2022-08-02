package cmd

import (

	"github.com/spf13/cobra"
	"github.com/Shilpa-Gokul/Kube-go-client/my-cli/deploy"
)

// deployNginxCmd represents the details of deployNginx Command
var deployNginxCmd = &cobra.Command{
	Use:   "deployNginx",
	Short: "Deploy Nginx image in Kubernetes",
	Long: `Deploy Nginx image with the provided scale count and version.
	If no scale and version arguments are provided, default values will be considered.
	Default Values:
	Scale- 1
	Version- 1.12`,
	Run: func(cmd *cobra.Command, args []string) {
		scale,_ := cmd.Flags().GetInt("scale")
		version,_ := cmd.Flags().GetString("version") 
		deploy.DeployImage(scale, version)
	},
}

func init() {
	rootCmd.AddCommand(deployNginxCmd)
	// Add the flags required for deployNginx command
	deployNginxCmd.PersistentFlags().Int("scale", 1, "Count of the replica set")
	deployNginxCmd.PersistentFlags().String("version", "1.12", "Version of the nginx image to be deployed")

}
