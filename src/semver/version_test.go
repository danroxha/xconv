package semver

import (
	"reflect"
	"testing"
)

func TestGivenThreeReasonsToUpgradeTheVersionThenItShouldIncrementCorrectly(t *testing.T) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	expected := Version{
		Major: 1,
		Minor: 1,
		Path:  0,
	}

	version.IncrementVersion("major", STANDARD)
	version.IncrementVersion("patch", STANDARD)
	version.IncrementVersion("minor", STANDARD)

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func TestGivenTheVersionChangeThenMajorShouldTakePriority(t *testing.T) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	expected := Version{
		Major: 1,
		Minor: 0,
		Path:  0,
	}

	version.IncrementVersion("patch", STANDARD)
	version.IncrementVersion("minor", STANDARD)
	version.IncrementVersion("major", STANDARD)

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func TestGivenTheMAJOR_MIJORAndPATCHSequenceChangeVersionThenItShouldIncrementCorrectly(t *testing.T) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	expected := Version{
		Major: 1,
		Minor: 1,
		Path:  1,
	}

	version.IncrementVersion("major", STANDARD)
	version.IncrementVersion("minor", STANDARD)
	version.IncrementVersion("patch", STANDARD)

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func TestGivenMultipleInterpolatedVersionUpThenItShouldIncrementCorrectly(t *testing.T) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	expected := Version{
		Major: 2,
		Minor: 0,
		Path:  1,
	}

	version.IncrementVersion("minor", STANDARD)
	version.IncrementVersion("patch", STANDARD)
	version.IncrementVersion("patch", STANDARD)
	version.IncrementVersion("minor", STANDARD)
	version.IncrementVersion("major", STANDARD)
	version.IncrementVersion("minor", STANDARD)
	version.IncrementVersion("patch", STANDARD)
	version.IncrementVersion("major", STANDARD)
	version.IncrementVersion("patch", STANDARD)

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func TestGivenThreeReasonsToUpgradeTheVersionThenItShouldIncrementCorrectlyOnAlphaMode(t *testing.T) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	expected := Version{
		Major: 0,
		Minor: 2,
		Path:  0,
	}

	version.IncrementVersion("major", ALPHA)
	version.IncrementVersion("patch", ALPHA)
	version.IncrementVersion("minor", ALPHA)

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func TestGivenThreeReasonsToUpgradeTheVersionThenItShouldIncrementCorrectlyOnBetaMode(t *testing.T) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	expected := Version{
		Major: 0,
		Minor: 2,
		Path:  0,
	}

	version.IncrementVersion("major", BETA)
	version.IncrementVersion("patch", BETA)
	version.IncrementVersion("minor", BETA)

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func BenchmarkIcrementVersion(b *testing.B) {
	version := Version{
		Major: 0,
		Minor: 0,
		Path:  0,
	}

	version.IncrementVersion("major", BETA)
}
