package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
    // _ "github.com/charmbracelet/lipgloss"
	_ "gopkg.in/yaml.v3"
)

var (
    argc int = len(os.Args)
    argv []string = os.Args
    YamlPath string = os.Getenv("HOME")+"/.local/share/golem"
)

// add new item 
func NewItem(filepath string) {
    var title        string
    var descriptiion string
    var check        bool

    tmp := DoubleCheck("testing DoubleCheck function")
    fmt.Println(tmp)

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
            Title("TODO title").
            Value(&title).
            Validate(func(str string) error {
                if str == "" {
                    return errors.New("The title cannot be empty")
                }
                return nil
            }),

            huh.NewText().
            Title("Description").
            CharLimit(400).
            Value(&descriptiion),

            huh.NewConfirm().
            Title("Add it?").
            Value(&check),
        ),
    ).WithTheme(huh.ThemeCharm())

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    // check = DoubleCheck("Are you sure?")

    fmt.Println(title)
    fmt.Println(descriptiion)
    fmt.Println(check)

    // for _, value := range toppings {
    //     fmt.Println(value)
    // }

}

// double check
func DoubleCheck(check_message string) bool {
    var check        bool
    form := huh.NewForm(
        huh.NewGroup(
            huh.NewConfirm().
            Title(check_message).
            Value(&check),
        ),
    ).WithTheme(huh.ThemeCharm())

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    return check
}

// intractive menu 
func IntractiveMenu(filepath string) string {
    var choice string
    form := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[string]().
                Title("Golem TODO app, What's your business here?").
                Options(
                    huh.NewOption("I wanna add a new task", "add"),
                    huh.NewOption("Give me a summary", "show"),
                    huh.NewOption("I need to change smth", "edit"),
                    huh.NewOption("Some works are done, don't need them any more", "delete"),
                ).
            Value(&choice),
        ),
    ).WithTheme(huh.ThemeCharm())

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    return choice
}

// display items 
func DisplayItems(filepath string) {
}


// main func
func main() {
    // fmt.Println(argc)
    // fmt.Println(argv)
    if argc == 1 {
        switch menu := IntractiveMenu("/tmp/test.yml"); menu {
        case "add":
            NewItem(YamlPath+"/db.yml")
        case "see":
            // see
        case "edit":
            // edit
        case "delete":
            // delete
        default:
            log.Fatalln("Unkown behavior")
        }
        // NewItem("test")
    }
}

