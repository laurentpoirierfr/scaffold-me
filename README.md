# Scaffold-me

### Usage

```bash
scaffold-me -h
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  scaffold-me [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get and execute scaffolder from url.
  help        Help about any command

Flags:
  -h, --help     help for scaffold-me
  -t, --toggle   Help message for toggle

Use "scaffold-me [command] --help" for more information about a command.
```

During your work, position yourself in a working directory.

```bash
mkdir exemple
cd exemple
```
Run the scaffold-me tool

```bash
scaffold-me get -h
Get and execute scaffolder from url.

Usage:
  scaffold-me get [flags]

Flags:
      --branch string     Branch name off scaffolder. (default "main")
  -h, --help              help for get
      --password string   User password to git repository.
      --tag string        Tag version off scaffolder.
      --url string        Git url off scaffolder. (default "https://github.com/laurentpoirierfr/default-scaffold.git")
      --user string       User login to git repository.
```

### Creation of a scaffolder

Directory Organization Example

```bash
.
├── cmd
│   ├── main.go
│   └── README.md.tpl
├── filename
│   └── %filename%-exemple.md
├── go.mod.tpl
├── README.md.tpl
└── scaffold.yml
```

### Note 

*.tpl files will be renamed. So it's not mandatory to use it, unless you want to escape some files. Example of file escaped: .gitlab-ci.yml by taking .gitlab-ci.yml.tpl this avoids the activation of the file.


### scaffold.yml 

```yaml
version: "1"
description: "[ Création d'un projet backend golang ]"
fields:
  - name : "project-name"
    description: "Nom du projet     : "
    default: "gitlab.com/homezone/project-name"
  - name : "project-version"
    description: "Version du projet : "
    default: "v0.1.0"
  - name : "filename"
    description: "Nom du fichier : "
    default: "filename"
```

Example of integrating a field into a file

* https://github.com/laurentpoirierfr/default-scaffold




