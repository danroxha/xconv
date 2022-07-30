package subcmd

import (
	"fmt"
	"strings"

	"github.com/dannrocha/xconv/src/gitscm"
	"github.com/urfave/cli/v2"
)

func TagCurrent(c *cli.Context) error {

	git, err := gitscm.New()

	if err != nil {
		panic(err.Error())
	}

	if git.IsTagsEmpty() {
		fmt.Println("There are no tag in this repository")
		return nil
	}

	tag := git.GitTags[0]


	if strings.Contains(c.String("format"), `%V`) {
		fmt.Printf("%v\n", tag.Annotation)
		return nil		
	}

	fmt.Printf("%v\t %v\t %v\n", tag.Date, tag.Annotation, tag.Author)
	return nil
}