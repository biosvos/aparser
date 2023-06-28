package aparser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAParser(t *testing.T) {
	ret := NewAParser([]*Argument{
		{
			Flags:       []string{"language", "l"},
			Required:    true,
			Description: "언어를 설정한다.",
			Default:     "ko",
			Mandatory:   "",
		},
	})

	require.NotNil(t, ret)
}

func TestMandatory(t *testing.T) {
	ret := NewAParser([]*Argument{
		{
			Description: "언어를 설정한다.",
			Mandatory:   "Language",
		},
	})

	t.Run("pass", func(t *testing.T) {
		_, err := ret.Parse([]string{"Program", "ko"})
		require.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		_, err := ret.Parse([]string{"Program"})
		require.Error(t, err)
	})
}

func TestAliases(t *testing.T) {
	ret := NewAParser([]*Argument{
		{
			Flags:       []string{"language", "l"},
			Description: "언어를 설정한다.",
		},
	})

	for _, tt := range []struct {
		name      string
		arguments []string
		result    string
	}{
		{
			arguments: []string{"Program", "-l", "ko"},
			result:    "ko",
		},
		{
			arguments: []string{"Program", "-language", "ko"},
			result:    "ko",
		},
		{
			arguments: []string{"Program"},
			result:    "",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ret.Parse(tt.arguments)
			require.NoError(t, err)
			require.EqualValues(t, tt.result, *result["l"])
		})
	}
}

func TestRequired(t *testing.T) {
	ret := NewAParser([]*Argument{
		{
			Flags:       []string{"language", "l"},
			Required:    true,
			Description: "언어를 설정한다.",
		},
	})

	t.Run("pass", func(t *testing.T) {
		result, err := ret.Parse([]string{"Program", "-l", "ko"})
		require.NoError(t, err)
		require.EqualValues(t, "ko", *result["l"])
	})

	t.Run("fail", func(t *testing.T) {
		_, err := ret.Parse([]string{"Program"})
		require.Error(t, err)
	})
}

func TestDefault(t *testing.T) {
	parser := NewAParser([]*Argument{
		{
			Flags:       []string{"language", "l"},
			Description: "언어를 설정한다.",
			Default:     "ko",
		},
	})

	t.Run("default", func(t *testing.T) {
		result, err := parser.Parse([]string{"Program"})
		require.NoError(t, err)
		require.EqualValues(t, "ko", *result["l"])
	})

	t.Run("en", func(t *testing.T) {
		result, err := parser.Parse([]string{"Program", "-l", "en"})
		require.NoError(t, err)
		require.EqualValues(t, "en", *result["l"])
	})

	t.Run("en", func(t *testing.T) {
		result, err := parser.Parse([]string{"Program", "-language", "en"})
		require.NoError(t, err)
		require.EqualValues(t, "en", *result["l"])
	})
}
