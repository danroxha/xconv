package semver

import (
	"reflect"
	"testing"
)

func TestShouldCheckIfTheTagIsValidWhitoutSufix(t *testing.T) {
	semver := New("v0.0.0")
	if !semver.IsVersionValid() {
		t.Fail()
	}
}

func TestShouldCheckIfTheTagIsValid(t *testing.T) {
	semver := New("v10.0.0-SNAPSHOT")
	version, _ := semver.FindVersion()
	expected := Version{
		Major: 10,
		Minor: 0,
		Path:  0,
	}

	if !semver.IsVersionValid() || !reflect.DeepEqual(expected, version) {
		t.Fail()
	}
}

func TestShouldCheckIfTheTagIsInvalid(t *testing.T) {
	semver := New("v0.0.0INVALID")

	if semver.IsVersionValid() {
		t.Fail()
	}
}

func BenchmarkSemVerNew(b *testing.B) {
	New("v10.0.0-SNAPSHOT")
}