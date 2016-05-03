package cmd

import (
	"github.com/elct9620/granblue.api/crawler"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(cmdCrawler)
	cmdCrawler.AddCommand(cmdCharacterList)
}

var cmdCrawler = &cobra.Command{
	Use:   "crawler",
	Short: "Run crawler to fetch data",
	Long:  `A set of commands which can fetch data from gbf-wiki.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Empty command, show help to tell user what to do
		cmd.Help()
	},
}

var (
	cmdCharacterList = &cobra.Command{
		Use:   "characters",
		Short: "Fetch character list",
		Run:   fetchCharacterList,
	}
)

func fetchCharacterList(cmd *cobra.Command, args []string) {
	// TODO(elct9620): Call crawler to get character data

	c := crawler.New()
	c.SetParser(&crawler.CharacterParser{})
	c.SetPage(crawler.CHARACTERS_PAGE_NAME)
	c.Fetch()
	c.Parse()
}
