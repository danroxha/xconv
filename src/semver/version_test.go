package semver

import (
	"reflect"
	"testing"
)

func TesgivenThreeReasonsToUpgradeTheVersionThenItShouldIncrementCorrectly(t *testing.T) {
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

	version.IncrementVersion("major")
	version.IncrementVersion("patch")
	version.IncrementVersion("minor")

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

	version.IncrementVersion("patch")
	version.IncrementVersion("minor")
	version.IncrementVersion("major")

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

	version.IncrementVersion("major")
	version.IncrementVersion("minor")
	version.IncrementVersion("patch")

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

	version.IncrementVersion("minor")
	version.IncrementVersion("patch")
	version.IncrementVersion("patch")
	version.IncrementVersion("minor")
	version.IncrementVersion("major")
	version.IncrementVersion("minor")
	version.IncrementVersion("patch")
	version.IncrementVersion("major")
	version.IncrementVersion("patch")

	if !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}
