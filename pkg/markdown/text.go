package markdown

import (
	"fmt"
	"strings"
)

type Plain string

func (t Plain) Repr() string {
	return fmt.Sprint(t)
}

type Italic string

func (t Italic) Repr() string {
	return fmt.Sprintf("*%s*", t)
}

type Bold string

func (t Bold) Repr() string {
	return fmt.Sprintf("**%s**", t)
}

type BoldItalic string

func (t BoldItalic) Repr() string {
	return fmt.Sprintf("_%s_**", t)
}

type StrikeThrough string

func (t StrikeThrough) Repr() string {
	return fmt.Sprintf("~~%s~~", t)
}

type InlineCode string

func (t InlineCode) Repr() string {
	return fmt.Sprintf("`%s`", t)
}

type Link struct {
	Text string
	Url  string
}

func (t *Link) Repr() string {
	return fmt.Sprintf("[%s](%s)", t.Text, t.Url)
}

type Image struct {
	Caption string
	Url     string
}

func (t *Image) Repr() string {
	return fmt.Sprintf("![%s](%s)", t.Caption, t.Url)
}

type Phrase interface {
	Repr() string
}

type Text = []Phrase

func String(t Text) string {
	var phrase []string
	for _, p := range t {
		phrase = append(phrase, p.Repr())
	}
	return strings.Join(phrase, " ")
}
