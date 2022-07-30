package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dannrocha/xconv/src/gitscm"
	"github.com/dannrocha/xconv/src/setup"
	"github.com/urfave/cli/v2"
)

func Tag(c *cli.Context) error {

	script := setup.NewScript()

	for _, task := range script.Task {
		if task.Bind == TAG && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
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
