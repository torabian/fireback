package workspaces

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Generic autocomplete model
type cliQueryModel struct {
	input      string
	cursor     int
	options    []string
	selected   map[string]bool // Track selected items
	loading    bool
	onComplete func([]string)
	selection  []string
	title      string
	queryFunc  func(string, int) ([]string, *QueryResultMeta, error) // Generic query func
	page       int                                                   // Keep track of pagination page
}

func (m cliQueryModel) Init() tea.Cmd {
	m.selected = make(map[string]bool) // Initialize selected items map
	m.page = 1
	return m.querySuggestions("")
}

func CaptureIds(items []string) []string {
	result := []string{}
	for _, item := range items {
		value := item
		if strings.Contains(value, ">>>") {
			value = strings.TrimSpace(strings.Split(item, ">>>")[0])
		}
		result = append(result, value)
	}
	return result
}

func (m cliQueryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+@": // Toggle selection
			if len(m.options) > 0 {
				item := m.options[m.cursor]
				if m.selected[item] {
					delete(m.selected, item)
				} else {
					m.selected[item] = true
				}
			}
		case "enter": // Finalize selection
			selectedItems := []string{}
			for item := range m.selected {
				selectedItems = append(selectedItems, item)
			}
			return m, tea.Batch(
				func() tea.Msg {
					m.onComplete(CaptureIds(selectedItems))
					return nil
				},
				tea.ClearScreen,
				tea.Quit,
			)
		case "up": // Pagination: go to the previous page
			if m.cursor > 0 {
				m.cursor--
			} else if m.page > 1 {
				m.page--
				m.loading = true
				return m, m.querySuggestions(m.input)
			}
		case "down": // Pagination: go to the next page
			if m.cursor < len(m.options)-1 {
				m.cursor++
			} else {
				m.page++
				m.loading = true
				return m, m.querySuggestions(m.input)
			}
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
				m.loading = true
				return m, m.querySuggestions(m.input)
			}
		case "ctrl+c":
			return m, tea.Quit
		default:
			m.input += msg.String()
			m.loading = true
			return m, m.querySuggestions(m.input)
		}

	case []string: // Handle new suggestions
		newOptions := msg
		newSelected := make(map[string]bool)

		// Persist selections if items still exist in the new options
		for _, option := range newOptions {
			if m.selected[option] {
				newSelected[option] = true
			}
		}

		for selectedItem := range m.selected {
			if _, found := newSelected[selectedItem]; !found {
				// The item is no longer in the options list, but we still want to keep it in memory
				newSelected[selectedItem] = true
			}
		}

		m.options = newOptions
		m.selected = newSelected
		m.loading = false
		return m, nil
	}
	return m, nil
}

func (m cliQueryModel) View() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s:\n", m.title))                                                                        // Display the title
	b.WriteString(("Use arrows to navigate, CTRL+Space toggle selection, and Enter to close, or type in for search\n")) // Display the title
	b.WriteString("Keyword: " + m.input + "\n\n")
	if m.loading {
		b.WriteString("Loading...\n")
	} else if len(m.options) == 0 {
		b.WriteString("No results found.\n")
	} else {
		for i, option := range m.options {
			cursor := " "
			if i == m.cursor {
				cursor = ">"
			}

			// Check if the option is selected by its content
			selected := "[ ]"
			if m.selected[option] {
				selected = "[x]"
			}

			b.WriteString(fmt.Sprintf("%s %s %s\n", cursor, selected, option))
		}
	}
	return b.String()
}

func (m cliQueryModel) querySuggestions(input string) tea.Cmd {
	return func() tea.Msg {

		results, _, err := m.queryFunc(input, m.page)
		if err != nil {
			return []string{"Error fetching results"}
		}
		var options []string
		for _, item := range results {
			options = append(options, item)
		}
		return options
	}
}

func CliInteractiveSearchAndSelect(
	title string,
	fn func(keyword string, page int) ([]string, *QueryResultMeta, error),
) []string {
	selection := []string{}
	model := cliQueryModel{
		onComplete: func(s []string) {
			selection = s
		},
		title:     title,
		queryFunc: fn,
	}
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return selection
}

func QueryStringCastCli(searchFields []string, keyword string, page int) QueryDSL {
	searchFieldsNew := []string{}
	for _, item := range searchFields {
		searchFieldsNew = append(searchFieldsNew, strings.ReplaceAll(item, "{keyword}", keyword))
	}

	query2 := strings.Join(searchFieldsNew, " or ")

	query := QueryDSL{
		ItemsPerPage: 10,
		Query:        query2,
		StartIndex:   page,
	}

	return query
}
