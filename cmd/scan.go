package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/jedipunkz/ecrscan/pkg/myecr"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var setFlags struct {
	image  string
	region string
}

// Image is struct for Container Image
type Image struct {
	imageName string
	tag       string
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan image",
	Run: func(cmd *cobra.Command, args []string) {
		e := myecr.Ecr{}

		i := Image{}
		if !i.getImageTag(setFlags.image) {
			log.Fatal("That image not found.")
		}

		e.Repositories = [][]string{
			{i.imageName, i.tag},
		}
		e.Resion = setFlags.region

		finding, _, err := e.ListFindings()
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf("\x1b[32m%s\x1b[0m", "ImageName: "+i.imageName+", ")
		// fmt.Printf("\x1b[32m%s\x1b[0m", "Tag: "+i.tag+"\n")

		data := [][]string{}
		for k, v := range finding.FindingSeverityCounts {
			data = append(data, []string{k, strconv.FormatInt(*v, 10)})
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Severity Level", "Counts"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&setFlags.image, "image", "", "", "Image Name")
	if err := scanCmd.MarkFlagRequired("image"); err != nil {
		log.Fatal(err)
	}
	scanCmd.Flags().StringVarP(&setFlags.region, "region", "", "", "Region Name")
	if err := scanCmd.MarkFlagRequired("region"); err != nil {
		log.Fatal(err)
	}
}

func (i *Image) getImageTag(image string) bool {
	if !strings.Contains(image, ":") {
		return false
	}
	slice := strings.Split(image, ":")
	i.imageName = slice[0]
	i.tag = slice[1]
	return true
}
