package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func Generate(version string, hash plumbing.Hash) object.Tag {
	return object.Tag{
		Name:       fmt.Sprintf("v%s", version),
		Message:    fmt.Sprintf("Release of v%s", version),
		Target:     hash,
		TargetType: plumbing.CommitObject,
		Tagger: object.Signature{
			Name: "Sharecare Versioning CLI",
			When: time.Now(),
		},
	}
}

func Tag(versionNumber string) error {
	version := fmt.Sprintf("v%v", versionNumber)

	wd, err := os.Getwd()

	check(err)

	repo, err := git.PlainOpen(wd)

	check(err)

	target, err := repo.Head()

	check(err)

	e := repo.Storer.NewEncodedObject()

	tag := Generate(version, target.Hash())

	tag.Encode(e)

	hash, err := repo.Storer.SetEncodedObject(e)

	check(err)

	ref := fmt.Sprintf("refs/tags/%s", version)

	check(err)

	err = repo.Storer.SetReference(plumbing.NewReferenceFromStrings(ref, hash.String()))

	check(err)

	return nil
}
