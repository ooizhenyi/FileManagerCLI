package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	name, desc string
	path       string
	isDir      bool
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.name }

type model struct {
	list        list.Model
	currentPath string
	quitting    bool
	err         error
}

func InitialModel(path string) model {
	return model{
		list:        list.New(getItems(path), list.NewDefaultDelegate(), 0, 0),
		currentPath: path,
	}
}

func getItems(path string) []list.Item {
	files, err := os.ReadDir(path)
	if err != nil {
		return []list.Item{}
	}

	var items []list.Item
	if path != "/" && path != "." {
		items = append(items, item{name: "..", desc: "Go Up", path: filepath.Dir(path), isDir: true})
	}

	for _, file := range files {
		desc := "File"
		if file.IsDir() {
			desc = "Directory"
		} else {
			info, _ := file.Info()
			desc = fmt.Sprintf("%d bytes", info.Size())
		}
		items = append(items, item{
			name:  file.Name(),
			desc:  desc,
			path:  filepath.Join(path, file.Name()),
			isDir: file.IsDir(),
		})
	}
	return items
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok && i.isDir {
				m.currentPath = i.path
				m.list.SetItems(getItems(m.currentPath))
				m.list.ResetSelected()
				return m, nil
			}
		case "backspace", "left":
			parent := filepath.Dir(m.currentPath)
			if parent != m.currentPath {
				m.currentPath = parent
				m.list.SetItems(getItems(m.currentPath))
				m.list.ResetSelected()
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	return docStyle.Render(m.list.View())
}
