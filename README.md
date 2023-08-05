# Scaffold-me

### Usage

```bash
git clone https://github.com/laurentpoirierfr/scaffold-me
cd scaffold-me
go install
```

During your work, position yourself in a working directory.

```bash
mkdir exemple
cd exemple
```
Run the scaffold-me tool

```bash
scaffold-me -r=<REPO_GIT_HTTPS> -v=<VERSION>
```

Default value

```bash
scaffold-me 
```
Correspond to

```bash
scaffold-me -r=https://github.com/laurentpoirierfr/default-scaffold -v=main
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

** Note **: *.tpl files will be renamed. So it's not mandatory to use it, unless you want to escape some files. Example of file escaped: .gitlab-ci.yml by taking .gitlab-ci.yml.tpl this avoids the activation of the file.


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




