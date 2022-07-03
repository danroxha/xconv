package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dannrocha/czen/src/git"
	"github.com/urfave/cli/v2"
)

func Tag(c *cli.Context) error {
	tags, err := git.LoadGitTags()

	if err != nil {
		panic(err.Error())
	}

	if len(tags) == 0 {
		fmt.Println("There are no tag in this repository")
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', tabwriter.Debug)
	defer writer.Flush()

	fmt.Fprintf(writer, " * Data\t * Version\t * Author\n")

	for _, tag := range tags {
		fmt.Fprintf(writer, "%v\t %v\t %v\n", "Sat Jul 2 20:06:22 2022", tag.Annotation, tag.Author)
	}

	return nil
}
