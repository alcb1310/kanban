# Kanban App

## Preview

https://github.com/alcb1310/kanban/assets/97707578/ccd44f5f-be72-44d2-816f-c5f817c5599e

## Description

This is an application that will allow you to create and manage a kanban board.

## Installation

### Using Homebrew

To install the project using *Homebrew*:

```bash
brew tap alcb1310/alcb1310
brew install kanban
```

### Upgrade

To upgrade from a previous version of the project using *Homebrew*:

```bash
brew update && brew upgrade kanban
```


### Build Steps

#### Prerequisites

- Go - [https://golang.org](https://golang.org)
- Git - [https://git-scm.com](https://git-scm.com)

#### Installation

To install the project, first clone the repository.

```bash
cd ~
git clone https://github.com/alcb1310/kanban
cd kanban
```

To build the project:

```bash
go build main.go
```

To install the project:

```bash
go install main.go
```

## Usage

To run the application:

```bash
kanban
```

- To select an item within a list:
    - Press the down arrow key or j to select the next item
    - Press the up arrow key or k to select the previous item

- To move an item to the next list:
    - Press the right arrow key or l to move the item to the next list
    - Press the left arrow key or h to move the item to the previous list

- To change an item from one list to another press the space bar

- To create a new todo item press n

- To exit the application press ctrl+c

- To display the help message in the command prompt run `kanban -h` or `kanban --help`

## Tech Stack

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)

## License

MIT
