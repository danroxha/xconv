# CZEN - Commit ZEN
Expired Study Project on [Committizen](https://github.com/commitizen-tools/commitizen)
## Requirements
- [Git] `^2.`+
## Usage
### Help
Run in your terminal

```bash
czen help
```

or the shortcut

```bash
czen h
```
- output
```
NAME:
   czen - Commit ZEN is a cli tool to generate conventional commits.

USAGE:
   czen [-h] {init,commit,example,info,tag,schema,bump,changelog,version}

AUTHOR:
   Rocha da Silva, Daniel <rochadaniel@acad.ifma.edu.br>

COMMANDS:
   init, i        init commitizen configuration
   commit, c      create new commit
   changelog, ch  generate changelog (note that it will overwrite existing file)
   bump, b        bump semantic version based on the git log
   rollback, r    revert commit to a specific tag
   tag, t         show tags
   schema, s      show commit schema
   example, e     show commit example
   version, v     get the version of the installed czen or the current project
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)


COPYRIGHT:
   (c) 2022 MIT
```
### Committing

```bash
czen commit
```

or the shortcut

```bash
czen c
```

### Rollback to Tag

```bash
czen rollback
```

or the shortcut

```bash
czen r
```

### Show tags

```bash
czen tag
```

or the shortcut

```bash
czen t
```
Select the tag for rollback and confirm