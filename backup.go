package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
)

const CmdPGDUMP = "pg_dump"

func backup(name string) (bool, error) {
	if _, err := os.Stat("scenarios"); os.IsNotExist(err) {
		must(os.Mkdir("scenarios", os.ModePerm))
	}

	var path = fmt.Sprintf("scenarios/%s.sql", name)

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return false, nil
	}

	var cmd = exec.Command(
		CmdPGDUMP,
		"--host="+viper.GetString("host"),
		"--port="+viper.GetString("port"),
		"--user="+viper.GetString("user"),
		"--schema="+viper.GetString("schema"),

		"--no-owner", "--clean",

		"--file="+path,

		viper.GetString("database"),
	)

	stdin, err := cmd.StdinPipe()
	must(err)
	stdout, err := cmd.StdoutPipe()
	must(err)
	stderr, err := cmd.StderrPipe()
	must(err)
	defer stdin.Close()
	defer stdout.Close()
	defer stderr.Close()

	must(cmd.Start())

	_, err = stdin.Write([]byte(viper.GetString("password") + "\n"))
	must(err)

	output, err := ioutil.ReadAll(stdout)
	must(err)
	output, err = ioutil.ReadAll(stderr)
	must(err)

	if err := cmd.Wait(); err != nil {
		println(string(output))
		println(string(output))
		panic(err)
	}

	must(preprocess(path))
	return true, nil
}
