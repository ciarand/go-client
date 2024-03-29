package main

import (
	"fmt"
	"github.com/keybase/go-libkb"
	"os"
)

// Keep this around to simplify things
var G = &libkb.G

func parseArgs() (libkb.Command, error) {
	p := libkb.PosixCommandLine{}
	cmd, err := p.Parse(os.Args)
	if err != nil {
		err = fmt.Errorf("Error parsing command line arguments: %s\n", err.Error())
		return nil, err
	}
	G.SetCommandLine(p)
	return cmd, nil
}

func testLogging() {
	G.Log.Debug("hello debug")
	G.Log.Info("hello info")
	G.Log.Notice("hello notice")
	G.Log.Warning("hello warning")
	G.Log.Error("hello error")
}

func main() {
	G.Init()
	err := main2()
	e2 := G.Shutdown()
	if err == nil {
		err = e2
	}
	if err != nil {
		G.Log.Error(err.Error())
		os.Exit(2)
	}
}

func main2() error {

	cmd, err := parseArgs()
	if cmd == nil || err != nil {
		return err
	}
	return G.RunCmdline(cmd)
}
