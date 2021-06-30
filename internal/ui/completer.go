package ui

import "github.com/c-bata/go-prompt"

func HostCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "build", Description: "Build a payload"},
		{Text: "listen", Description: "Start a server to listen for connections"},
		{Text: "host=", Description: "Set a host"},
		{Text: "port=", Description: "Set a port"},
		{Text: "output=", Description: "Specify a filename to output"},
		{Text: "--windows", Description: "For Windows"},
		{Text: "--macos", Description: "For Mac OS"},
		{Text: "--linux", Description: "For Linux"},
		{Text: "--hidden", Description: "Run client in hidden mode"},
		{Text: "exit", Description: "Exit this program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
