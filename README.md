# Scaffold-me

### Usage

```bash
git clone https://github.com/laurentpoirierfr/scaffold-me
cd scaffold-me
go install
```

Lors de vos travaux, positionnez-vous dans un répertoire de travail.

```bash
mkdir exemple
cd exemple
```
Lancer l'outil scaffold-me

```bash
scaffold-me -r=<REPO_GIT_HTTPS> -v=<VERSION>
```

Valeur par défaut

```bash
scaffold-me 
```
Correspond à 

```bash
scaffold-me -r=https://github.com/laurentpoirierfr/default-scaffold -v=main
```


### Création d'un scaffolder

Exemple d'organisation de répertoire 

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

Remarque : les fichiers *.tpl seront renommés. Donc ce n'est pas obligatoire de l'utilisé, sauf si vous voulez échappé certains  fichiers. Exemple de fichier a échapé : .gitlab-ci.yml en prenant .gitlab-ci.yml.tpl cela évite l'activation du fichier.



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

Exemple d'intégraztion d'un field dans un fichier

* https://github.com/laurentpoirierfr/default-scaffold


