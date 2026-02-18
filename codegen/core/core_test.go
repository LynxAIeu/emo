package core_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lynxai-team/emo/codegen/core"
)

func TestSanitize(t *testing.T) {
	t.Parallel()

	input := []core.Ref{{
		Name:    "Info",
		Emoji:   "i",
		IsError: false,
	}, {
		Name:    "PanicBear",
		Emoji:   "p",
		IsError: true,
	}, {
		Name:    "DataBaseDebug",
		Emoji:   "",
		IsError: false,
	}, {
		Name:    "Query",
		Emoji:   "q",
		IsError: false,
	}}

	want := []core.Ref{{
		Name:    "DataBaseDebug",
		Emoji:   "",
		IsError: false,
	}, {
		Name:    "Info",
		Emoji:   "i",
		IsError: false,
	}, {
		Name:    "PanicBear",
		Emoji:   "p",
		IsError: true,
	}, {
		Name:    "Query",
		Emoji:   "q",
		IsError: false,
	}, {
		Name:    "QueryError",
		Emoji:   "q",
		IsError: true,
	}}

	got := core.Sanitize(input)
	if !cmp.Equal(got, want) {
		t.Errorf("got  %v", got)
		t.Errorf("want %v", want)
	}
}

func TestSanitizePanic(t *testing.T) {
	t.Parallel()

	input := []core.Ref{{
		Name:    "SameName",
		Emoji:   "i",
		IsError: false,
	}, {
		Name:    "PanicBear",
		Emoji:   "p",
		IsError: true,
	}, {
		Name:    "DataBaseDebug",
		Emoji:   "",
		IsError: false,
	}, {
		Name:    "SameName",
		Emoji:   "q",
		IsError: false,
	}}

	defer func() {
		err := recover()
		if err != nil {
			return
		}
		t.Error("expected panic")
	}()

	_ = core.Sanitize(input)
}
