package markdown

import (
	"fmt"
	"strings"
)

type Heading struct {
	Text  *Text
	Level uint8
}

func (h *Heading) Render() string {
	return fmt.Sprintf(
		"%s %s",
		strings.Repeat("#", int(h.Level)),
		String(h.Text),
	)
}

type ListItem struct {
	Text  *Text
	List  *[]ListItem
	Order bool
}

func (l *ListItem) Render(prefix string, level uint8) string {
	text := ""
	if l.Text != nil {
		text = fmt.Sprintf(
			"%s %s %s",
			strings.Repeat("  ", int(level)),
			prefix,
			String(l.Text),
		)
	}
	if len(*l.List) > 0 {
		text += "\n"
		for i, li := range *l.List {
			childPrefix := "*"
			if l.Order {
				childPrefix = fmt.Sprint(i + 1)
			}
			text += li.Render(childPrefix, level+1)
		}
	}
	return text
}

type List struct {
	List  *[]ListItem
	Order bool
}

func (l *List) Render() string {
	text := ""
	for i, li := range *l.List {
		prefix := "*"
		if l.Order {
			prefix = fmt.Sprint(i + 1)
		}
		text += li.Render(prefix, 1) + "\n"
	}
	return text
}

type Code struct {
	Text string
	Lang string
}

func (c *Code) Render() string {
	return fmt.Sprintf("```%s\n%s\n```", c.Lang, c.Text)
}

type Paragraph struct {
	Text *Text
}

func (p *Paragraph) Render() string {
	return String(p.Text)
}

type BlockQuote struct {
	Text *Text
}

func (q *BlockQuote) Render() string {
	return fmt.Sprintf("> %s", q.Text)
}

type Line struct{}

func (l *Line) Render() string {
	return "---"
}

type NewLine struct{}

func (l *NewLine) Render() string {
	return ""
}
