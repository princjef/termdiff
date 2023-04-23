//go:build mage
// +build mage

package main

import (
	"github.com/princjef/mageutil/bintool"
	"github.com/princjef/mageutil/shellcmd"
)

var linter = bintool.Must(bintool.New(
	"golangci-lint{{.BinExt}}",
	"1.51.1",
	"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
))

var documenter = bintool.Must(bintool.New(
	"gomarkdoc{{.BinExt}}",
	"0.4.1",
	"https://github.com/princjef/gomarkdoc/releases/download/v{{.Version}}/gomarkdoc_{{.Version}}_{{.GOOS}}_{{.GOARCH}}{{.ArchiveExt}}",
))

func Lint() error {
	if err := linter.Ensure(); err != nil {
		return err
	}

	return linter.Command(`run`).Run()
}

func Generate() error {
	return shellcmd.Command(`go generate .`).Run()
}

func Doc() error {
	if err := documenter.Ensure(); err != nil {
		return err
	}

	return documenter.Command(`.`).Run()
}

func DocVerify() error {
	if err := documenter.Ensure(); err != nil {
		return err
	}

	return documenter.Command(`-c .`).Run()
}

func Test() error {
	return shellcmd.Command(`go test -count 1 -coverprofile=coverage.txt ./...`).Run()
}

func Coverage() error {
	return shellcmd.Command(`go tool cover -html=coverage.txt`).Run()
}
