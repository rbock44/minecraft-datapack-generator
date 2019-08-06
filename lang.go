package main

import (
	"fmt"
)

type namespace string
type functionName string
type functionRef string
type objective string
type target string
type selector string

type message struct {
	Word  string
	Text  string
	Color string
}

type messageList []message

func newMessageList(m ...message) *[]message {
	return &m
}

func (c *functionContent) tellRaw(ml *[]message) {
	rootContentFormat := "tellraw @s [%s]"
	content := ""
	for _, m := range *ml {
		content += fmt.Sprintf(`{"text": "%s", "color": "%s"},`, m.Text, m.Color)
	}
	// remove trailing comma
	content = content[:len(content)-1]
	c.Content.WriteString(fmt.Sprintf(rootContentFormat, content))
}

func (c *functionContent) addHomePlayer() {
	c.Content.WriteString("tag @s add home_player\n")
}

func (c *functionContent) removeHomePlayer() {
	c.Content.WriteString("tag @s add home_player\n")
}

func (c *functionContent) addLine(line string) {
	c.Content.WriteString(line + "\n")
}

func (c *functionContent) addFunctionRef(ref functionRef) {
	c.Content.WriteString(fmt.Sprintf("function %s\n", ref))
}

type criteria string

const (
	criteriaDummy   criteria = "dummy"
	criteriaTrigger criteria = "trigger"
)

func (c *functionContent) addObjective(o objective, cr criteria, displayName string) {
	text := ""
	if displayName == "" {
		text = fmt.Sprintf("scoreboard objectives add %s %s", o, cr)
	} else {
		text = fmt.Sprintf("scoreboard objectives add %s %s [%s]", o, cr, displayName)
	}
	c.Content.WriteString(text)
}

func (c *functionContent) scoreReset(t target, o objective) {
	c.addLine(fmt.Sprintf("scoreboard players reset %s %s", t, o))
}

func (c *functionContent) executeAs(t target, s string) {
	c.addLine(fmt.Sprintf("execute as %s", s))
}

func (c *functionContent) executeAt(t target, s string) {
	c.addLine(fmt.Sprintf("execute at %s %s", t, s))
}

func (c *functionContent) executeIf(s string) {
	c.addLine(fmt.Sprintf("execute if %s", s))
}

func (c *functionContent) executeUnless(s string) {
	c.addLine(fmt.Sprintf("execute unless %s", s))
}

func (c *functionContent) executeStore(s string) {
	c.addLine(fmt.Sprintf("execute store %s", s))
}
