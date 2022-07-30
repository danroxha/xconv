package setup

const (
	LUA     = "lua"
	SH      = "sh"
	BEFORE  = "before"
	AFTER   = "after"
	LIST    = "list"
	CONFIRM = "confirm"
	INPUT   = "input"
)

type Option struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
	Name  string `yaml:"name"`
}

type Question struct {
	Type       string   `yaml:"type"`
	Message    string   `yaml:"message"`
	Name       string   `yaml:"name"`
	Default    string   `yaml:"default"`
	Middleware []string `yaml:"middleware"`
	Filter     string   `yaml:"filter"`
	Editor     bool     `yaml:"editor"`
	Choices    []Option `yaml:"choices"`
}

type ExitCode struct {
	Exception   string
	ExitCode    int
	Description string
}

type Configuration struct {
	Rule   Rule   `yaml:"rule"`
	Script Script `yaml:"script"`
}

type Rule struct {
	Version       string    `yaml:"version"`
	ActiveProfile string    `yaml:"active_profile"`
	ChangelogFile string    `yaml:"changelog_file"`
	Profiles      []Profile `yaml:"profiles"`
}

type Tag struct {
	Format string `yaml:"format"`
	Mode   string `yaml:"mode"`
}

type Bump struct {
	Map     map[string]string `yaml:"map"`
	Pattern string            `yaml:"pattern"`
}

type Profile struct {
	Bump            Bump       `yaml:"bump"`
	Name            string     `yaml:"name"`
	CommitParser    string     `yaml:"commit_parser"`
	ChangeTypeOrder []string   `yaml:"change_type_order"`
	Example         string     `yaml:"example"`
	MessageTemplate string     `yaml:"message_template"`
	Questions       []Question `yaml:"questions"`
	Schema          string     `yaml:"schema"`
	Tag             Tag        `yaml:"tag"`
	Extends         string     `yaml:"extends"`
}

type Message struct {
	Content string `yaml:"content"`
	Color   bool   `yaml:"color"`
}

type ScriptBase struct {
	Name   string `yaml:"name"`
	Enable bool   `yaml:"enable"`
	Script string `yaml:"script"`
	Type   string `yaml:"type"`
}

type Filter struct {
	ScriptBase `yaml:",inline"`
	Retry      bool    `yaml:"retry"`
	Message    Message `yaml:"message"`
}

type Task struct {
	ScriptBase `yaml:",inline"`
	Bind       string `yaml:"bind"`
	Language   string `yaml:"language"`
	When       string `yaml:"when"`
}

type Middleware struct {
	ScriptBase `yaml:",inline"`
}

type Script struct {
	Filter     []Filter     `yaml:"filters"`
	Middleware []Middleware `yaml:"middlewares"`
	Task       []Task       `yaml:"tasks"`
}

// https://github.com/commitizen-tools/commitizen/blob/master/docs/exit_codes.md
var ExitCodeStardard map[string]ExitCode = map[string]ExitCode{
	"ExpectedExit": {
		Exception:   "ExpectedExit",
		ExitCode:    0,
		Description: "Expected exit",
	},

	"FileSetupExist": {
		Exception:   "FileSetupExist",
		ExitCode:    1,
		Description: "there is .xconv.yaml configuration in the repository",
	},

	"NotAGitProjectError": {
		Exception:   "NotAGitProjectError",
		ExitCode:    2,
		Description: "not a git repository (or any of the parent directories): .git",
	},

	"NoCommitsFoundError": {
		Exception:   "NoCommitsFoundError",
		ExitCode:    3,
		Description: "no commit found",
	},

	"NoVersionSpecifiedError": {
		Exception:   "NoCommitsFoundError",
		ExitCode:    4,
		Description: "version can not be found in configuration file [.xconv.yaml]",
	},

	"NoPermissionOnDir": {
		Exception:   "NoPermissionOnDir",
		ExitCode:    5,
		Description: ".xconv.yaml file cannot be create in the current directory",
	},

	"BumpRegexInvalid": {
		Exception:   "BumpRegexInvalid",
		ExitCode:    6,
		Description: "bump.pattern in .xconv.yaml invalid",
	},
	"GitNotFound": {
		Exception:   "GitNotFound",
		ExitCode:    7,
		Description: "git not found!. visit <https://git-scm.com/> to install",
	},

	"NothingToCommitError": {
		Exception:   "NothingToCommitError",
		ExitCode:    11,
		Description: "nothing in staging to be committed",
	},

	"MissingConfigError": {
		Exception:   "MissingConfigError",
		ExitCode:    15,
		Description: "configuration missed for .xconv.yaml",
	},

	"CurrentVersionNotFoundError": {
		Exception:   "CurrentVersionNotFoundError",
		ExitCode:    17,
		Description: "current version cannot be found in version_files",
	},

	"InvalidConfigurationError": {
		Exception:   "InvalidConfigurationError",
		ExitCode:    19,
		Description: "an error was found in the xconv configuration",
	},

	"NoneIncrementExit": {
		Exception: "InvalidConfigurationError",
		ExitCode:  21,
		Description: "the commits found are not elegible to be bumped",
	},

	"InvalidScriptFilter": {
		Exception:   "InvalidScriptFilter",
		ExitCode:    22,
		Description: "an error was found in the script",
	},
	"InvalidProfile": {
		Exception: "InvalidProfile",
		ExitCode: 23,
		Description: "profiles have recursive inheritance",
	},
}
