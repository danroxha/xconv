package script

import (
	"fmt"
	"os"

	"github.com/dannrocha/czen/src/config"
	lua "github.com/yuin/gopher-lua"
)

type ScriptHandle struct {
	Script string
	Args []string
}


func (s ScriptHandle) Run() string {
	return RunScriptLuaReturnString(s.Script, s.Args...)
}

func (s ScriptHandle) RunFilter() bool {
	return RunScriptLuaReturnBool(s.Script, s.Args...)
}


func RunScriptLuaReturnBool(script string, args ...string) bool {

	L := lua.NewState()
	defer L.Close()

	err := L.DoString(script)
	if err != nil {
		throw := config.ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description)
		os.Exit(throw.ExitCode)
	}

	arguments := []lua.LValue{}

	for _, arg := range args {
		arguments = append(arguments, lua.LString(arg))
	}

	err = L.CallByParam(
		lua.P{
			Fn:      L.GetGlobal("run"),
			NRet:    1,
			Protect: true,
		},
		arguments...,
	)

	if err != nil {
		panic(err)
	}

	ret := L.Get(-1)

	result, success := ret.(lua.LBool)

	if !success {
		throw := config.ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description)
		os.Exit(throw.ExitCode)
	}


	return bool(result)

}


func RunScriptLuaReturnString(script string, args ...string) string {

	L := lua.NewState()
	defer L.Close()

	err := L.DoString(script)
	if err != nil {
		throw := config.ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description)
		os.Exit(throw.ExitCode)
	}

	arguments := []lua.LValue{}

	for _, arg := range args {
		arguments = append(arguments, lua.LString(arg))
	}

	err = L.CallByParam(
		lua.P{
			Fn:      L.GetGlobal("run"),
			NRet:    1,
			Protect: true,
		},
		arguments...,
	)

	if err != nil {
		panic(err)
	}

	ret := L.Get(-1)

	result, _ := ret.(lua.LString)

	return string(result)

}
