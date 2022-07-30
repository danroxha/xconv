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