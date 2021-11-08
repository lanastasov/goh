package main

import (
    "fmt"
    "os"
	 "os/exec"
	 "runtime"
	 "log"
	 "context"

	 "github.com/google/go-github/v39/github"
    tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	repo_name []string
	repo_url []string
	cursor int
	selected map[int]struct{}
}



func open_browser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}



func initial_model(r []*github.Repository) (*model) {
	max := len(r)
	rs := new(model)
	rs.repo_name = make([]string, max)
	rs.repo_url = make([]string, max)
	
	for i := 0; i < max; i++ {
		rs.repo_name[i] = *r[i].Name
		rs.repo_url[i] = *r[i].HTMLURL
	}
	rs.selected = make(map[int]struct{})
	
	return rs
}


func clear_screen() {

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	// no need for default or error checking
	// if we're calling this function, one of these
	// options will work
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.repo_name)-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
					 open_browser(m.repo_url[m.cursor])
					 clear_screen()
					 return m, tea.Quit
            }
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    //s := "repos\n\n"
	 var s string

    // Iterate over our choices
    for i, choice := range m.repo_name {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = "\033[31m>\033[0m" // cursor!
        }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func check_args(args []string) string {
	if len(args[1:]) != 1 {
		return ""
	} else {
		return args[1]
	}
}

func print_usage() {
	fmt.Println("# goh")
	fmt.Println("	navigate github repos in a TUI")
	fmt.Println()
	fmt.Println("usage: goh <github username>")
	fmt.Println()
}


func main() {
	
	user_name := check_args(os.Args)
	if user_name == "" {
		print_usage()
		os.Exit(1)
	}
	// get github info
	client := github.NewClient(nil)
	repos, _, err := client.Repositories.List(context.Background(), user_name, nil)
	if err != nil {
		fmt.Println("error: problem getting repo data for user");
		os.Exit(1)
	}


    p := tea.NewProgram(initial_model(repos))
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

