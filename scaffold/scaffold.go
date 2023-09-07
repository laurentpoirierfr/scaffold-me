package scaffold

import (
	"errors"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rivo/tview"
	"gopkg.in/yaml.v3"
)

const (
	QUIT          = "Quit"
	SAVE          = "Save"
	SCAFFOLD_YAML = "scaffold.yml"
)

type Scaffold struct {
	Version     string  `yaml:"version"`
	Description string  `yaml:"description"`
	Fields      []Field `yaml:"fields"`
	Exclude     Exclude `yaml:"exclude"`
}

type Field struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Default     string `yaml:"default"`
}

type Exclude struct {
	Files   []string
	Folders []string
}

type Scaffolder struct {
	Scaffold Scaffold
	Source   string
	Target   string
	Folders  []string
	Files    []string
	Values   map[string]string
	Response string
}

func NewScaffolder(source, target string) (Scaffolder, error) {
	var scaffold Scaffold

	yamlFile, err := os.ReadFile(source + "/" + SCAFFOLD_YAML)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return Scaffolder{}, err
	}
	err = yaml.Unmarshal(yamlFile, &scaffold)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return Scaffolder{}, err
	}

	return Scaffolder{
		Source:   source,
		Target:   target,
		Scaffold: scaffold,
	}, nil
}

func (s *Scaffolder) Execute() error {
	s.getTemplateValues()
	err := filepath.WalkDir(s.Source, s.walk)
	if err != nil {
		return err
	}
	err = s.createTargetFolders()
	if err != nil {
		return err
	}
	err = s.copyTargetFiles()
	if err != nil {
		return err
	}

	return nil
}

func (s *Scaffolder) getTemplateValues() {
	values := make(map[string]string)
	s.Response = QUIT

	for _, field := range s.Scaffold.Fields {
		values[field.Name] = field.Default
	}

	app := tview.NewApplication()

	form := tview.NewForm()
	form.AddButton(SAVE, func() {
		app.Stop()
		s.Response = SAVE
	})
	form.AddButton(QUIT, func() {
		app.Stop()
		s.Response = QUIT
	})

	for _, field := range s.Scaffold.Fields {
		form.AddInputField(field.Description, field.Default, 100, nil, func(value string) {
			values[field.Name] = value
		})
	}

	form.SetBorder(true).SetTitle(s.Scaffold.Description).SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	s.Values = values
}

func (s *Scaffolder) walk(source string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		if !strings.Contains(source, SCAFFOLD_YAML) {
			s.Files = append(s.Files, strings.ReplaceAll(source, s.Source, ""))
		}
	} else {
		s.Folders = append(s.Folders, strings.ReplaceAll(source, s.Source, ""))
	}
	return nil
}

func (s *Scaffolder) createTargetFolders() error {
	for i := 0; i < len(s.Folders); i++ {
		path := s.Target + s.Folders[i]
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scaffolder) isExcluded(filename string) bool {

	for i := 0; i < len(s.Scaffold.Exclude.Folders); i++ {
		if strings.Contains(filename, s.Scaffold.Exclude.Folders[i]) {
			return true
		}
	}

	for i := 0; i < len(s.Scaffold.Exclude.Files); i++ {
		if strings.Contains(filename, s.Scaffold.Exclude.Files[i]) {
			return true
		}
	}

	return false
}

func (s *Scaffolder) copyTargetFiles() error {
	for i := 0; i < len(s.Files); i++ {
		path := s.Target + s.changeFileName(s.Files[i])

		fileContent, err := os.ReadFile(s.Source + s.Files[i])
		if err != nil {
			log.Fatal(err)
		}

		if !s.isExcluded(path) {
			text := string(fileContent)
			t, err := template.New("scaffolder").Parse(text)
			if err != nil {
				return err
			}
			// Template if not
			f, err := os.Create(path)
			if err != nil {
				return err
			}
			err = t.Execute(f, s.Values)
			if err != nil {
				return err
			}
		} else {
			// Copy if exclude only
			err = os.WriteFile(path, fileContent, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scaffolder) changeFileName(name string) string {
	filename := strings.ReplaceAll(name, ".tpl", "")
	for key, value := range s.Values {
		filename = strings.ReplaceAll(filename, "%"+key+"%", value)
	}
	return filename
}
