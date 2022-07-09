package semver

import (
	"reflect"
	"testing"
)

func TestM(t *testing.T) {
	semver := New("v0.0.0")
	if !semver.IsVersionValid() {
		t.Fail()
	}
}

func TestO(t *testing.T) {
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

func TestP(t *testing.T) {
	semver := New("v0.0.0INVALID")

	if semver.IsVersionValid() {
		t.Fail()
	}
}