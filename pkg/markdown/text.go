package markdown

import "fmt"

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

type Code struct {
	Text string
}

func (t *Code) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

type Phrase interface {
	String() string
}

type Text = []Phrase

func String(t *Text) string {
	text := ""
	for _, p := range *t {
		text += p.String()
	}
	return text
}
