package cmd

import (
	"os"

	"github.com/jedipunkz/ecrscan/pkg/myecr"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listSetFlags struct {
	image  string
	region string
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list image's vulnerability",
	Run: func(cmd *cobra.Command, args []string) {
		e := myecr.Ecr{}

		i := Image{}
		if !i.getImageTag(listSetFlags.image) {
			log.Fatal("That image not found.")
		}

		e.Repositories = [][]string{
			{i.imageName, i.tag},
		}
		e.Resion = listSetFlags.region

		_, vulFindings, err := e.ListFindings()
		if err != nil {
			log.Fatal(err)
		}

		data := [][]string{}
		for _, v := range vulFindings {
			data = append(data, []string{v.Name, v.Severity, v.URI, v.Description})
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"CVE Name", "Severity", "URI", "Description"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listSetFlags.image, "image", "", "", "Image Name")
	listCmd.MarkFlagRequired("image")
	listCmd.Flags().StringVarP(&listSetFlags.region, "region", "", "", "Region Name")
	listCmd.MarkFlagRequired("region")
}
