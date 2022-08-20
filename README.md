# XCONV - X Conventional
Expired Study Project on [Committizen](https://github.com/commitizen-tools/commitizen)
## Requirements
- [Git] `^2.`+
## Usage
### Help
Run in your terminal

```bash
xc help
```

or the shortcut

```bash
xc h
```
- output
```
NAME:
   xc - X Conventional is a cli tool to generate conventional commits and versioning.

USAGE:
   xc [-h] {init,commit,example,info,tag,schema,bump,changelog,version}

AUTHOR:
   Rocha da Silva, Daniel <rochadaniel@acad.ifma.edu.br>

COMMANDS:
   init, i        init xconv configuration
   commit, c      create new commit
   changelog, ch  generate changelog (note that it will overwrite existing file)
   bump, b        bump semantic version based on the git log
   rollback, r    revert commit to a specific tag
   tag, t         show tags
   schema, s      show commit schema
   example, e     show commit example
   version, v     get the version of the installed xconv or the current project
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)


COPYRIGHT:
   (c) 2022 MIT
```
### Committing

```bash
xc commit
```

or the shortcut

```bash
xc c
```

### Rollback to Tag

```bash
xc rollback
```

or the shortcut

```bash
xc r
```

### Show tags

```bash
xc tag
```

or the shortcut

```bash
xc t
```
Select the tag for rollback and confirm


## Configuration
### Default
```yml
rule:
  changelog_file: CHANGELOG.md
  active_profile: xconv_default

  profiles:
  - name: xconv_default
    extends: nil
    tag:
      restricted: true
      mode: standard
      format: v$version
      
    bump:
      map:
        'BREAKING CHANGE': MAJOR
        feature: MINOR
        bugfix: PATCH
        hotfix: PATCH
        pattern: ^(BREAKING[\-\ ]CHANGE|feature|hotfix|docs|bugfix|refactor|perf)(\(.+\))?(!)?

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

    schema: |
      <prefix>(<scope>): <subject>
      <BLANK LINE>
      <body>
      <BLANK LINE>
      (BREAKING CHANGE: <footer> )
      <footer>
    
    example: |
      bugfix: correct minor typos in code

      see the work item for details on the typos fixed

      related work items #12
    
    info: This is default info from xconv
    info_path: xconv_info.txt

    message_template: | 
      {{prefix}}({{scope}}): {{subject}}
      
      {% if body != '' %}
      {{body}}
      {% endif %}

      {% if is_breaking_change %}
      BREAKING CHANGE: 
      {% endif %}
      {% if footer != '' %}
      Related work items: #{{footer}}
      {% endif %}
    questions:
      - type: list
        message: "Select the type of change you are committing:"
        name: prefix
        choices:
          - value: feature
            key: f
            name: "feature: A new feature."

          - value: bugfix
            name: "bugfix: A bug fix. Correlates with PATCH in SemVer"
            key: b

          - value: hotfix
            name: "hotfix: A bug fix in PROD"
            key: h

          - value: docs
            name: "docs: Documentation only changes"
            key: d

          - value: style
            name: "style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"
            key: s

          - value: refactor
            name: "refactor: A code change that neither fixes a bug nor adds a feature"
            key: r

          - value: perf
            name: "perf: A code change that improves performance"
            key: p

          - value: test
            name: "test: Adding missing or correcting existing tests"
            key: t

          - value: chore
            name: "chore: Changes to configuration files  (example scopes: .gitignore, .xconv.yaml)"
            key: z

          - value: build
            name: "build: Changes that affect the build system or external dependencies (example scopes: pip, docker, npm)"
            key: x

          - value: ci
            name: "ci: Changes to our CI configuration files and scripts (example scopes: AzureDevOps)"
            key: c

      - type: input
        message: "What is the scope of this change? (class or file name): (press [enter] to skip): "
        name: scope
        middleware: 
        - to lower case
        - trim
        filter: is empty

      - type: input
        message: "Write a short and imperative summary of the code changes: (lower case and no period): "
        name: subject
        middleware: 
        - to lower case
        - trim
        filter: is empty

      - type: input
        message: "Provide additional contextual information about the code changes: (press [enter] to skip): "
        name: body
        middleware:
        - to lower case

      - type: confirm
        name: is_breaking_change
        message: "Is this a BREAKING CHANGE? Correlates with MAJOR in SemVer (press [enter] to skip): "
        default: false

      - type: input
        message: "Related work items (PBI, Task IDs, Issue): (press [enter] to skip)"
        name: footer
        middleware:
        - to lower case

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
    - name: push current tag
      enable: true
      language: sh
      bind: bump
      when: after
      script: |
        CURRENT_TAG=$(xc tag current --format=%V)
        git push origin $CURRENT_TAG

```

### Custom
```yml
rule:
  changelog_file: CHANGELOG.md
  active_profile: alpha profile

  profiles:
  - name: alpha profile
    extends: xconv_default
    tag:
      mode: alpha # alpha beta | default: standard
      format: v$version # v$major.$minor.$patch$"

    message_template: | 
      {{prefix}}({{scope}}): {{subject}}
      
      {% if body != '' %}
      {{body}}
      {% endif %}

      {% if is_breaking_change %}
      BREAKING CHANGE: 
      {% endif %}
      {% if footer != '' %}
      issue: {{footer}}
      {% endif %}
script:
  tasks:
    - name: log bump
      enable: true
      language: sh
      bind: bump
      when: after
      script: |
        CURRENT_TAG=$(xc tag current --format=%V)
        
        if [ ! -e log-version.txt ]
        then
          echo $CURRENT_TAG > log-version.txt
          exit
        fi
        
        echo $CURRENT_TAG >> log-version.txt
```


### Description
- Rule

   | Variable | Type | Default | implemented | Description |
   |-|-|-|-|-|
   | changelog_file | string | CHANGELOG.md | :x: | |
   | active_profile | string | xconv_default | :heavy_check_mark: | |
   | [profiles](#profiles) | list<[profile](#profiles)> |  | :heavy_check_mark: | |


   - <span id="profiles">Profile</span>
      | Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | name | string | xconv_default | :heavy_check_mark: | |
      | extends | string | nil | :heavy_check_mark: | |
      | [tag](#tag) | object |  | :heavy_check_mark: | |
      | [bump](#bump) | object |  | :heavy_check_mark: | |
      | changelog_pattern | string | ^(feature\|bugfix\|hotfix\|perf\|refactor)?(!)? | :x: | |
      | change_type_map | map<string, string> |  { feature: Feature, bugfix: Bugfix, hotfix: Hotfix, perf: Performance, docs: Documentation, refactor: Refactor } | :x: | |
      | change_type_order | list&lt;string&gt; | BREAKING[\-\ ]CHANGE, feature, bugfix,  hotfix, refactor, perf | :x: | |
      | commit_parser | string | ^(?P<change_type>docs\|feature\|bugfix\|hotfix\|refactor\|perf\|BREAKING CHANGE)(\?:\\((?P<scope>[^()\r\n]*)\)\|\()?(?P<breaking>!)?:\s(?P<message>.*) | :x: | |
      | version_parser | string | (?P$version([0-9]+)\.([0-9]+)\.([0-9]+)(?:-\\([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+[0-9A-Za-z-]+)?(\w+)?) | :heavy_check_mark: | |
      | schema_pattern | string | (build\|ci\|docs\|feature\|bugfix\|hotfix\|perf\|refactor\|style\|test\|chore\|revert\|bump):\\(\\\(\S+\))?!?:\\(\s.*) | :x: | |
      | schema | string |  <code>&lt;prefix&gt;(&lt;scope&gt;): &lt;subject&gt;&lt;BLANK LINE&gt;<br/></br>&lt;body&gt;<br/><br/>&lt;BLANK LINE&gt;(BREAKING CHANGE: &lt;footer&gt; )&lt;footer&gt;<code> | :heavy_check_mark: | |
      | example | string | <code>bugfix(main.go): correct minor typos in code<br/><br/>see the work item for details on the typos fixed<br/><br/>related work items #12<code> | :heavy_check_mark: | |
      | info | string | xconv_default | :heavy_check_mark: | |
      | info_path | string | xconv_default | :heavy_check_mark: | |
      | message_template | string | <code>{{prefix}}({{scope}}): {{subject}}<br/><br/>{% if body != '' %}{{body}}{% endif %}<br/><br/>{% if is_breaking_change %}<br/>BREAKING CHANGE: <br/>{% endif %}<br/><br/>{% if footer != '' %}<br/>Related work items: #{{footer}}<br/>{% endif %}<code> | :heavy_check_mark: | |
      | [questions](#questions) | list | | :heavy_check_mark: | |

   - <span id="tag">Tag</span>
      | Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | restricted | string | true | :x: | |
      | mode | string | standard | :heavy_check_mark: | |
      | format | object | v$version | :heavy_check_mark: | |

   - <span id="bump">Bump</span>
      | Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | map | map<string, string> | {'BREAKING CHANGE': MAJOR, feature: MINOR, bugfix: PATCH, hotfix: PATCH } | :heavy_check_mark: | |
      | pattern | string | "^\(BREAKING\[\- \]CHANGE\|feature\|hotfix\|docs\|bugfix\|refactor\|perf\)\(\(.+\))?(!)?" | :heavy_check_mark: | |

   - <span id="questions">Questions</span>
      | Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | type | string | ```list\|input```| :heavy_check_mark: | |
      | message | string | | :heavy_check_mark: | |
      | name | string | | :heavy_check_mark: | |


- script
   | Variable | Type | Default | implemented | Description |
   |-|-|-|-|-|
   | [filter](#filters) | list<[filter](#filters)> | | :heavy_check_mark: | |
   | [middleware](#middlewares) | list<[middleware](#middlewares)> | | :heavy_check_mark: | |
   | [task](#tasks) | list<[task](#tasks)> | | :heavy_check_mark: | |


   - <span id="filters"> filter</span>
      | Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | name | string | | :heavy_check_mark: | |
      | retry | boolean | false | :heavy_check_mark: | |
      | enable | boolean | false | :heavy_check_mark: | |
      | message | object | | :heavy_check_mark: | |
      | script | string | | :heavy_check_mark: | |

   - <span id="middlewares"> middleware</span>
      | Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | name | string | | :heavy_check_mark: | |
      | enable | boolean | false | :heavy_check_mark: | |
      | script | string | | :heavy_check_mark: | |
   - <span id="tasks"> task</span>
      Variable | Type | Default | implemented | Description |
      |-|-|-|-|-|
      | name | string | | :heavy_check_mark: | |
      | enable | boolean | false | :heavy_check_mark: | |
      | language | string | ```lua``` | :heavy_check_mark: | |
      | bind | string | | :heavy_check_mark: | |
      | when | string | ```after```| :heavy_check_mark: | |
      | script | string | | :heavy_check_mark: | |
