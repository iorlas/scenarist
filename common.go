package main

import "github.com/pkg/errors"

func must(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}
