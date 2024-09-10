package main

import (
	"fmt"
	"log"
    "errors"

    "github.com/charmbracelet/huh"
    _ "gopkg.in/yaml.v3"
)

// add new item {{{
func NewItem(filepath string) {
    var title        string
    var descriptiion string
    var check        bool

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
            Title("Add the item to the todo list?").
            Value(&check),
        ),
    ).WithTheme(huh.ThemeBase())

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(title)
    fmt.Println(descriptiion)
    fmt.Println(check)

    // for _, value := range toppings {
    //     fmt.Println(value)
    // }

}
// }}}

// intractive menu {{{
func IntractiveMenu(filepath string) {
}
// }}}

// dispaly items {{{
func DisplayItems(filepath string) {
}
// }}}

func main() {
    NewItem("test")
}

