package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/mmcloughlin/cite"
	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process citations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return processFiles(args)
	},
}

func init() {
	RootCmd.AddCommand(processCmd)
}

func processFiles(filenames []string) error {
	if len(filenames) == 0 {
		return errors.New("no files specified")
	}

	for _, filename := range filenames {
		err := processFile(filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func processFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	src := cite.ParseCode(f)
	f.Close()

	builders := []cite.ResourceBuilder{
		cite.BuildGithubResourceFromCitation,
		cite.BuildPlainResourceFromCitation,
	}
	processor := cite.NewProcessor(builders)
	processor.AddHandler("insert", cite.InsertHandler)

	processed, err := processor.Process(src)
	if err != nil {
		return err
	}

	data := []byte(processed.String())
	return ioutil.WriteFile(filename, data, 0)
}
