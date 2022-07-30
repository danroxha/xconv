package setup

var XCONVFileContent string = `
rule:
  version: 0.0.0
  changelog_file: CHANGELOG.md
  active_profile: xconv_default

  profiles: 
  - name: xconv_default
    tag:
      mode: standard # alpha beta | default: standard
      format: v$version-SNAPSHOT # v$major.$minor.$patch$"
      
    bump_map:
      'BREAKING CHANGE': MAJOR
      feature: MINOR
      bugfix: PATCH
      hotfix: PATCH

    bump_message: "version $current_version \u2192 $new_version"
    bump_pattern: ^(BREAKING[\-\ ]CHANGE|feature|hotfix|docs|bugfix|refactor|perf)(\(.+\))?(!)?

    changelog_pattern: ^(feature|bugfix|hotfix|perf|refactor)?(!)?

    change_type_map:
      feature: Feature
      bugfix: Bugfix
      hotfix: Hotfix
      perf: Performance
      docs: Documentation
      refactor: Refactor

    change_type_order:
		- BREAKING[\-\ ]CHANGE
		- feature
		- bugfix
		- hotfix
		- refactor
		- perf

    commit_parser: ^(?P<change_type>docs|feature|bugfix|hotfix|refactor|perf|BREAKING CHANGE)(?:\((?P<scope>[^()\r\n]*)\)|\()?(?P<breaking>!)?:\s(?P<message>.*)
    version_parser: (?P$version([0-9]+)\.([0-9]+)\.([0-9]+)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+[0-9A-Za-z-]+)?(\w+)?)
    
    schema_pattern: (build|ci|docs|feature|bugfix|hotfix|perf|refactor|style|test|chore|revert|bump):(\(\S+\))?!?:(\s.*)

    schema: "
      <prefix>(<scope>): <subject>\n
      <BLANK LINE>\n
      <body>\n
      <BLANK LINE>\n 
      (BREAKING CHANGE: <footer> ) \n
      <footer>
    "
    example: "
      bugfix: correct minor typos in code\n\n
      see the work item for details on the typos fixed\n\n
      related work items #12
    "
    info: This is default info from xconv
    info_path: xconv_info.txt

    message_template: "{{prefix}}({{scope}}): {{subject}}\n\n{% if body != '' %}{{body}}\n\n{% endif %}{% if is_breaking_change %}BREAKING CHANGE: {% endif %}{% if footer != '' %}Related work items: #{{footer}}{% endif %}"
    questions:
		- type: list
			message: "Select the type of change you are committing:"
			name: prefix
			choices:
			- value: feature
				name: "feature: A new feature."

			- value: bugfix
				name: "bugfix: A bug fix. Correlates with PATCH in SemVer"

			- value: hotfix
				name: "hotfix: A bug fix in PROD"

			- value: docs
				name: "docs: Documentation only changes"

			- value: style
				name: "style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"

			- value: refactor
				name: "refactor: A code change that neither fixes a bug nor adds a feature"

			- value: perf
				name: "perf: A code change that improves performance"

			- value: test
				name: "test: Adding missing or correcting existing tests"

			- value: chore
				name: "chore: Changes to configuration files  (example scopes: .gitignore, .xconv.yaml)"

			- value: build
				name: "build: Changes that affect the build system or external dependencies (example scopes: pip, docker, npm)"

			- value: ci
				name: "ci: Changes to our CI configuration files and scripts (example scopes: AzureDevOps)"

		- type: input
			message: "What is the scope of this change? (class or file name): (press [enter] to skip): "
			name: scope
			middleware: 
			- to lower case
			- trim
			filter: is empty
			# editor: true # Open editor. Given undefined GIT_EDITOR and EDITOR try open [code, nano, vim, vi]
			# default: "Default value for name field"

		- type: input
			message: "Write a short and imperative summary of the code changes: (lower case and no period): "
			name: subject
			middleware: 
			- to lower case
			- trim
			filter: is_empty
			# editor: true # Open editor. Given undefined GIT_EDITOR and EDITOR try open [code, nano, vim, vi]
			# default: "Default value for name field"

		- type: input
			message: "Provide additional contextual information about the code changes: (press [enter] to skip): "
			name: body
			middleware:
			- to lower case
			# editor: true # Open editor. Given undefined GIT_EDITOR and EDITOR try open [code, nano, vim, vi]
			# filter: is empty
			# default: "Default value for name field"

		- type: confirm
			name: is breaking change
			message: "Is this a BREAKING CHANGE? Correlates with MAJOR in SemVer (press [enter] to skip): "
			default: false

		- type: input
			message: "Related work items (PBI, Task IDs, Issue): (press [enter] to skip)"
			name: footer
			middleware:
			- to lower case
			# editor: true # Open editor. Given undefined GIT_EDITOR and EDITOR try open [code, nano, vim, vi]
			# default: "Default value for name field"

script:
  filters:
	- name: is empty
		retry: true
		enable: true
		message:
			content: "[ALERT]: this field cannot be empty or only contain spaces"
			color: true
		script: |
			function run(argument)
				return argument == nil or argument == ''
			end

  middlewares:
	- name: to lower case
		enable: true
		script: |
			function run(argument)
				return string.lower(argument)
			end

	- name: trim
		enable: true
		script: |
			function run(argument)
				return (string.gsub(argument, "^%s*(.-)%s*$", "%1"))
			end

  tasks:
  - name: task example
    enable: true
    language: lua # sh | lua
    bind: schema
    when: before
    script: |
      print("Hello World")
`

var Filename string = ".xconv.yaml"

func NewConfiguration() Configuration {
	return Configuration{
		Rule: NewRule(),
		Script: NewScript(),
	}
}
