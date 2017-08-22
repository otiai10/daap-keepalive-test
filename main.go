package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/otiai10/daap"
)

func main() {

	fmt.Println(time.Now())

	machine := daap.NewEnvMachine()
	proc := daap.NewProcess("otiai10/daap-keepalive-test", daap.Args{
		Machine: machine,
	})

	ctx := context.Background()
	if err := proc.Run(ctx); err != nil {
		panic(err)
	}

	o, err := ioutil.ReadAll(proc.Stdout)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(bytes.Trim(o, " \n"), []byte("\n"))
	if bytes.HasPrefix(lines[len(lines)-1], []byte("STARTED: ")) {
		fmt.Println("ISSUE #3 POSITIVE!")
		t, _ := time.Parse("STARTED: 01/02/06 15:04", string(lines[len(lines)-1]))
		fmt.Println("Duration since last stdout:", time.Since(t))
		fmt.Println("See https://github.com/otiai10/daap/issues/3 for more information")
	} else {
		fmt.Println("Negative", time.Now())
	}
}
