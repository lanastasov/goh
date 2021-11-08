# goh
Navigate github repos in a tui

### Why
I am constantly refering to my github repos and repos from others for code snippets that are relevant to what I'm working on in the moment. I hope this will be easier than opening a browser, going to GitHub, clicking on the repos tab, and scrolling through the list. We'll see!

### Usage
```
goh <username>
```
- returns a list of the repos from `<username>`.
- move the cursor up and down the list with arrow keys.
- press enter to open the webpage of the selected repo.
- quit the program with `q` or `ctrl-c`

### Dependencies
`goh` is built using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) Go tui framework and [Go GitHub](https://github.com/google/go-github) GitHub API Go library.  

### To Do
- fix issue that if the list is longer than the term size, it goes haywire moving through the list
- add descriptions and languages to the tui menu
- maybe an option to look through files in the repos?
- setup installation. add to homebrew maybe?

