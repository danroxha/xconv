package setup

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var rule Rule = NewRule()

func TestNewRule(t *testing.T) {
	t.Run("should load active profile from xconv.yaml", func(t *testing.T) {
		rule := NewRule()
		profile, err := rule.FindCurrentProfileEnable()

		assert.Equal(t, err, nil, "error on load profile")

		expected := rule.ActiveProfile
		result := profile.Name

		assert.Equal(t, expected, result, "exptected %v, but result was %v to active profile", expected, result)

	})


	t.Run("shouldn't load active profile from xconv.yaml", func(t *testing.T) {
		rule := NewRule()
		rule.ActiveProfile = "undefinedProfile"
		_, err := rule.FindCurrentProfileEnable()

		expected := "Profile setup not found or disabled"

		assert.NotEqual(t, err, nil, "profile can't be undefined %v", nil)
		assert.Equal(t, expected, err.Error(), "expected error message '%v', but message was '%v'", expected, err.Error())
	
	})
}

func BenchmarkNewRule(b *testing.B) {
	NewRule()
}

func BenchmarkFindProfileByName(b *testing.B) {
	rule.FindProfileByName("xconv_default")
}

func BenchmarkFindCurrentProfileEnable(b *testing.B) {
	rule.FindCurrentProfileEnable()
}

func BenchmarkReplaceProfile(b *testing.B) {
	profile := Profile {
		Name: "xconv_default",
	}
	rule.ReplaceProfile(profile)
}