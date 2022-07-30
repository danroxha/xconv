package setup

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/dannrocha/xconv/src/cmd"
	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v3"
)

func NewScript() Script {
	script := Script{}

	script.loadScriptFromFile()

	return script
}

func (sc *Script) loadScriptFromFile() error {
	file := struct {
		Script Script `yaml:"script"`
	}{}

	blob, err := ioutil.ReadFile(".xconv.yaml")

	if err != nil {
		panic(err)
	}

	parseError := yaml.Unmarshal(blob, &file)

	if parseError != nil {
		panic(parseError)
	}

	*sc = file.Script

	return nil
}

func (sc *Script) FindAllFilters() map[string]Filter {
	filterGroup := make(map[string]Filter)

	for _, filter := range sc.Filter {
		filterGroup[filter.Name] = filter
	}

	return filterGroup
}

func (sc *Script) FindAllMiddlewares() map[string]Middleware {
	middlewareGroup := make(map[string]Middleware)

	for _, middleware := range sc.Middleware {
		middlewareGroup[middleware.Name] = middleware
	}

	return middlewareGroup
}

func (sc *Script) FindByFilterName(name string) (Filter, bool) {
	filterGroup := sc.FindAllFilters()
	filter, ok := filterGroup[name]
	return filter, ok
}

func (sc *Script) FindByMiddlewareName(name string) (Middleware, bool) {
	filterGroup := sc.FindAllMiddlewares()
	middleware, ok := filterGroup[name]
	return middleware, ok
}

func (middleware Middleware) Run(args ...string) string {
	L := lua.NewState()
	defer L.Close()

	err := L.DoString(middleware.Script)
	if err != nil {
		throw := ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description, ":", middleware.Name)
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

	result, ok := ret.(lua.LString)

	if !ok {
		throw := ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description, ":", middleware.Name)
		os.Exit(throw.ExitCode)
	}

	return string(result)
}

func (filter Filter) Run(args ...string) bool {

	L := lua.NewState()
	defer L.Close()

	err := L.DoString(filter.Script)
	if err != nil {
		throw := ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description, ":", filter.Name)
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

	result, ok := ret.(lua.LBool)

	if !ok {
		throw := ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description, ":", filter.Name)
		os.Exit(throw.ExitCode)
	}

	return bool(result)
}

func (task Task) Run(args ...string) {

	if task.Language == SH {

		var binarySh string

		if runtime.GOOS == "windows" {
			binarySh = findShExecutable()
		} else {
			binarySh = `sh`
		}

		command := cmd.InternalCommand{
			Application: binarySh,
			Args: []string{
				"-c",
				task.Script,
			},
		}

		output, err := command.Execute()

		if err != nil {
			throw := ExitCodeStardard["InvalidScriptFilter"]
			fmt.Printf("%v: \n - automation::%v::%v\n", throw.Description, task.Language, task.Name)
			os.Exit(throw.ExitCode)
		}

		fmt.Println(string(output))

		return
	}

	L := lua.NewState()
	defer L.Close()

	err := L.DoString(task.Script)
	if err != nil {
		throw := ExitCodeStardard["InvalidScriptFilter"]
		fmt.Println(throw.Description, ":", task.Name)
		os.Exit(throw.ExitCode)
	}

	arguments := []lua.LValue{}

	for _, arg := range args {
		arguments = append(arguments, lua.LString(arg))
	}

	L.CallByParam(
		lua.P{
			Fn:      L.GetGlobal("run"),
			NRet:    0,
			Protect: true,
		},
		arguments...,
	)
}

func findShExecutable() string {

	shLocale := cmd.InternalCommand{
		Application: `where`,
			Args: []string{
				`git`,
			},
	}

	output, err := shLocale.Execute()

	if err != nil {
		panic(err)
	}

	binaryGitPath := string(output)
	slicePath := strings.Split(binaryGitPath, `\`)
	gitFullPath := strings.Join(slicePath[:len(slicePath)-2], `\`)

	return gitFullPath + `\bin\sh.exe`

}
