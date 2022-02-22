package markdown

import (
	"fmt"
	"strings"
)

type Plain struct {
	Text string
}

func (t *Plain) String() string {
	return t.Text
}

type Italic struct {
	Text string
}

func (t *Italic) String() string {
	return fmt.Sprintf("*%s*", t.Text)
}

type Bold struct {
	Text string
}

func (t *Bold) String() string {
	return fmt.Sprintf("**%s**", t.Text)
}

type BoldItalic struct {
	Text string
}

func (t *BoldItalic) String() string {
	return fmt.Sprintf("_%s_**", t.Text)
}

type StrikeThrough struct {
	Text string
}

func (t *StrikeThrough) String() string {
	return fmt.Sprintf("~~%s~~", t.Text)
}

type InlineCode struct {
	Text string
}

func (t *InlineCode) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

type Link struct {
	Text string
	Url  string
}

func (t *Link) String() string {
	return fmt.Sprintf("[%s](%s)", t.Text, t.Url)
}

type Image struct {
	Text string
	Url  string
}

func (t *Image) String() string {
	return fmt.Sprintf("![%s](%s)", t.Text, t.Url)
}

type Phrase interface {
	String() string
}

type Text = []Phrase

func String(t Text) string {
	var phrase []string
	for _, p := range t {
		phrase = append(phrase, p.String())
	}
	return strings.Join(phrase, " ")
}
