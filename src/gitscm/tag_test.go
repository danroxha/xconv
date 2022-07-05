package gitscm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/dannrocha/czen/src/cmd"
	"github.com/dannrocha/czen/src/util"
)

func TestGivenATagCommitShowDetailsThenShouldToConvertToGitTagStruct(t *testing.T) {
	commit := `
		tag v0.1.1
		Tagger: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:27:27 2022 -0300


		commit 4fed22f43ea4c6eb9cc875fcb30bc68582ad9610
		Author: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:26:02 2022 -0300

			chore(.czen.yaml): new attributes added

			BREAKING CHANGE:
	`

	tag := ParseTag(commit)

	expected := GitTag{
		Annotation: "v0.1.1",
		Date:       "Mon Jul 4 18:27:27 2022 -0300",
		Author:     "Daniel Rocha <rochadaniel@acad.ifma.edu.br>",
		Commit: GitCommit{
			Hash:   "4fed22f43ea4c6eb9cc875fcb30bc68582ad9610",
			Date:   "Mon Jul 4 18:26:02 2022 -0300",
			Author: "Daniel Rocha <rochadaniel@acad.ifma.edu.br>",
		},
	}

	if !reflect.DeepEqual(expected, tag) {
		t.Errorf("Fail")
	}
}

func TestGivenATagCommitShowDetailswithoutStandardThenShouldToConvertToGitTagStruct(t *testing.T) {
	commit := `
		tag without-standard
		Tagger: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:27:27 2022 -0300


		commit 4fed22f43ea4c6eb9cc875fcb30bc68582ad9610
		Author: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:26:02 2022 -0300

			chore(.czen.yaml): new attributes added

			BREAKING CHANGE:
	`

	tag := ParseTag(commit)

	expected := GitTag{
		Annotation: "without-standard",
		Date:       "Mon Jul 4 18:27:27 2022 -0300",
		Author:     "Daniel Rocha <rochadaniel@acad.ifma.edu.br>",
		Commit: GitCommit{
			Hash:   "4fed22f43ea4c6eb9cc875fcb30bc68582ad9610",
			Date:   "Mon Jul 4 18:26:02 2022 -0300",
			Author: "Daniel Rocha <rochadaniel@acad.ifma.edu.br>",
		},
	}

	if !reflect.DeepEqual(expected, tag) {
		t.Errorf("Fail")
	}
}

func TestGivenATagCommitShowDetailsWithAuthorsNotTheSameThenShouldToConvertToGitTagStruct(t *testing.T) {
	commit := `
		tag v0.1.1
		Tagger: Rocha Daniel <danielrocha@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:27:27 2022 -0300


		commit 4fed22f43ea4c6eb9cc875fcb30bc68582ad9610
		Author: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:26:02 2022 -0300

			chore(.czen.yaml): new attributes added

			BREAKING CHANGE:
	`

	tag := ParseTag(commit)

	expected := GitTag{
		Annotation: "v0.1.1",
		Date:       "Mon Jul 4 18:27:27 2022 -0300",
		Author:     "Rocha Daniel <danielrocha@acad.ifma.edu.br>",
		Commit: GitCommit{
			Hash:   "4fed22f43ea4c6eb9cc875fcb30bc68582ad9610",
			Date:   "Mon Jul 4 18:26:02 2022 -0300",
			Author: "Daniel Rocha <rochadaniel@acad.ifma.edu.br>",
		},
	}

	if !reflect.DeepEqual(expected, tag) {
		t.Errorf("Fail")
	}
}



func TestGivenATagGroupTheLoadTagGroup(t *testing.T) {
	tags := generateTag()
	rm := cmd.InternalCommand {
		Application: "rm",
		Args: []string{
			"-rf",
			"repository/",
		},
	}

	defer rm.Execute()

	git := Git{
		Command: map[string]cmd.InternalCommand{
			"tag": {
				Application: "sh",
				Args: []string{
					"-c",
					`cd repository/ && git tag -l`,
				},
			},
			"show": {
				Application: "sh",
				Args: []string{
					"-c",
					`cd repository/ && git show -q %v`,
				},
			},
		},		
	}

	git.LoadGitTags()

	if !reflect.DeepEqual(git.Tags, tags) {
		t.Errorf("Fail")
	}

}



func generateTag() []string {

	build := cmd.InternalCommand{
		Application: "sh",
		Args: []string{
			"-c",
			`mkdir -p repository \
				&& cd repository \
				&& git init \
				&& touch README.md \
				&& git add README.md \
				&& git commit -m "message"
			`,
		},
	}
	
	build.Execute()


	for i := 0; i < 5; i++ {
		generate := cmd.InternalCommand{
			Application: "sh",
			Args: []string{
				"-c",
				fmt.Sprintf(`cd repository \
					&& git tag -a v0.0.%v -m "T"
				`, i),
			},
		}

		generate.Execute()
	}

	tagList := cmd.InternalCommand{
		Application: "sh",
		Args: []string{
			"-c",
			`cd repository/ && git tag -l`,
		},
	}

	out, _ := tagList.Execute()

	return util.RemoveContains(strings.Split(string(out), "\n"), "")

}