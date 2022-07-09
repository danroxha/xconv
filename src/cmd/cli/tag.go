package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dannrocha/czen/src/gitscm"
	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Tag(c *cli.Context) error {

	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {
		if auto.Bind == TAG && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	git, err := gitscm.New()

	if err != nil {
		panic(err.Error())
	}

	if git.IsTagsEmpty() {
		fmt.Println("There are no tag in this repository")
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', tabwriter.Debug)
	defer writer.Flush()

	fmt.Fprintf(writer, " * Data\t * Version\t * Author\n")

	for _, tag := range git.GitTags {
		fmt.Fprintf(writer, "%v\t %v\t %v\n", tag.Date, tag.Annotation, tag.Author)
	}

	return nil
}
