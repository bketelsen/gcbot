package bot

import (
	"testing"
)

func TestCommand(t *testing.T) {
	a1 := Arg{
		Name:              "Email",
		Required:          true,
		ValidationMessage: "Email is required.",
	}

	a2 := Arg{
		Name:              "TicketNumber",
		Required:          true,
		ValidationMessage: "Ticket Number is required.",
	}
	c := Command{
		Trigger: "validate",
		Args:    []Arg{a1, a2},
	}
	err := c.Parse("bketelsen@gmail.com 1234567 123456")
	if err == nil {
		t.Fatal("expected 1 error, got 0")
	}

	err = c.Parse("bketelsen@gmail.com 1234567")
	if err != nil {
		t.Fatalf("expected 0 errors, got 1 %s", err.Error())
	}
	if c.Args[0].Value != "bketelsen@gmail.com" {
		t.Fatalf("expected %s error, got %s", "bketelsen@gmail.com", c.Args[0].Value)
	}
}
