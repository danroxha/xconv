package gitscm

import (
	"reflect"
	"testing"
)

func TestGivenATagCommitShowDetailsThenShouldToConvertToGitTagStruct(t *testing.T) {
	commit := `
		tag v0.1.1
		Tagger: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:27:27 2022 -0300


		commit 4fed22f43ea4c6eb9cc875fcb30bc68582ad9610
		Author: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:26:02 2022 -0300

			chore(.xconv.yaml): new attributes added

			BREAKING CHANGE:
	`

	tag := parseTag(commit)

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

			chore(.xconv.yaml): new attributes added

			BREAKING CHANGE:
	`

	tag := parseTag(commit)

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

			chore(.xconv.yaml): new attributes added

			BREAKING CHANGE:
	`

	tag := parseTag(commit)

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

func BenchmarkParseTag(b *testing.B) {
	commit := `
		tag v0.1.1
		Tagger: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:27:27 2022 -0300


		commit 4fed22f43ea4c6eb9cc875fcb30bc68582ad9610
		Author: Daniel Rocha <rochadaniel@acad.ifma.edu.br>
		Date:   Mon Jul 4 18:26:02 2022 -0300

			chore(.xconv.yaml): new attributes added

			BREAKING CHANGE:
	`

	parseTag(commit)
}