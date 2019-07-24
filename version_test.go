package main_test

import (
	"testing"

	. "github.com/jredhooper/version"
)

func reset() {
	nope := false
	yup := true
	nothing := ""

	Major = &nope
	Minor = &yup
	Release = &nope
	Manual = &nothing
	Snapshot = &nope
}

func pass(input string, expected string) bool {
	output, err := UpdateVersion(input)

	if err != nil {
		return false
	}

	if output != expected {
		return false
	}

	return true
}

func TestUpdateVersion(suite *testing.T) {

	suite.Run("Improper Formatting", func(t *testing.T) {
		if pass("0.0", "") {
			t.Fail()
		}
	})

	suite.Run("Minor Update", func(t *testing.T) {
		if !pass("0.0.0", "0.0.1") {
			t.Fail()
		}
	})

	suite.Run("Minor Non-Number", func(t *testing.T) {
		if pass("0.0.a", "") {
			t.Fail()
		}
	})

	suite.Run("Minor with Snapshot", func(t *testing.T) {
		if !pass("0.0.0-SNAPSHOT", "0.0.1") {
			t.Fail()
		}
	})

	suite.Run("Major Update", func(t *testing.T) {
		toggle := true

		Major = &toggle

		if !pass("0.0.0", "0.1.0") {
			t.Fail()
		}

		t.Run("With Minor", func(t *testing.T) {
			if !pass("0.1.1", "0.2.0") {
				t.Fail()
			}
		})

		if pass("0.a.0", "") {
			t.Fail()
		}
	})

	suite.Run("Release Update", func(t *testing.T) {
		toggle := true

		Release = &toggle

		if !pass("0.0.0", "1.0.0") {
			t.Fail()
		}

		t.Run("With Minor", func(t *testing.T) {
			if !pass("0.0.1", "1.0.0") {
				t.Fail()
			}
		})

		t.Run("With Major", func(t *testing.T) {
			if !pass("0.1.0", "1.0.0") {
				t.Fail()
			}
		})

		if pass("a.0.0", "") {
			t.Fail()
		}
	})

	suite.Run("Manual", func(t *testing.T) {
		update := "0.0.0"

		Manual = &update

		if !pass("3.2.1", "0.0.0") {
			t.Fail()
		}

		t.Run("Bad Manual Input", func(t *testing.T) {
			bad := "a.b.c"
			Manual = &bad
			if pass("3.2.1", "0.0.0") {
				t.Fail()
			}
		})
	})

	suite.Run("Snapshot", func(t *testing.T) {
		reset()

		toggle := true

		Snapshot = &toggle

		if !pass("0.0.0", "0.0.0-SNAPSHOT") {
			t.Fail()
		}

	})
}
