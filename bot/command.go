package bot

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandAction func(Command, *discordgo.Session, *discordgo.User, string) error

type Command struct {
	Bot         *Discord
	Trigger     string
	Args        []Arg
	Action      CommandAction
	Description string
}

func (c *Command) Help() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Command: `%s`\t", c.Trigger)

	for _, a := range c.Args {
		opt := "optional"
		if a.Required {
			opt = "required"
		}
		fmt.Fprintf(&b, "%s (%s) ", a.Name, opt)
	}

	fmt.Fprintf(&b, "\n\t%s\n", c.Description)
	return b.String()
}
func (c *Command) Parse(rest string) error {
	values := strings.Fields(rest)

	if len(values) > len(c.Args) {
		return errors.New("too many arguments provided")
	}
	if len(values) != len(c.Args) {
		var b strings.Builder
		for _, a := range c.Args {
			fmt.Fprintf(&b, "{%s} ", a.ValidationMessage)
		}
		return errors.New(b.String())
	}
	for i, v := range values {
		c.Args[i].Value = v
	}
	return nil
}

type Arg struct {
	Name              string
	Required          bool
	ValidationMessage string
	Value             string
}
