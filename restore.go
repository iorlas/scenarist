package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
)

const CmdPSQL = "psql"

func restore(name string) (bool, error) {
	var path = fmt.Sprintf("scenarios/%s.sql", name)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, nil
	}

	var cmd = exec.Command(
		CmdPSQL,
		"--host="+viper.GetString("host"),
		"--port="+viper.GetString("port"),
		"--user="+viper.GetString("user"),

		"--single-transaction",

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

	stdoutResult, err := ioutil.ReadAll(stdout)
	must(err)

	stderrResult, err := ioutil.ReadAll(stderr)
	must(err)

	if err := cmd.Wait(); err != nil {
		println(string(stdoutResult))
		println(string(stderrResult))
		panic(err)
	}

	must(preprocess(path))
	return true, nil
}
