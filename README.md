# README

## Scaffold-me

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
