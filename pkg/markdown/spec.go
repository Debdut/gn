package markdown

import (
	"fmt"
	"strings"
)

type Heading struct {
	Text  Text
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
	Text  Text
	List  []ListItem
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
	if len(l.List) > 0 {
		text += "\n"
		for i, li := range l.List {
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
	List  []ListItem
	Order bool
}

func (l *List) Render() string {
	text := ""
	for i, li := range l.List {
		prefix := "*"
		if l.Order {
			prefix = fmt.Sprint(i + 1)
		}
		text += li.Render(prefix, 1) + "\n"
	}
	return text
}

type Table struct {
	Headers []Text
	Rows    [][]Text
	Align   []string
}

func CellString(cell string, width int) string {
	spaces := strings.Repeat(" ", width-len(cell))
	return cell + spaces
}

func RowString(row []string, maxColLen []int) string {
	text := "| "
	for i, cell := range row {
		text += CellString(cell, maxColLen[i]) + " |"
	}
	return text
}

func SeparatorString(width int, align string) string {
	if align == "right" {
		return strings.Repeat("-", width-1) + ":"
	} else if align == "center" {
		return ":" + strings.Repeat("-", width-2) + ":"
	}
	return strings.Repeat("-", width)
}

func SeparatorRowString(maxColLen []int, align []string) string {
	text := "| "
	if len(align) != len(maxColLen) {
		for _, colLen := range maxColLen {
			text += SeparatorString(colLen, "left") + " |"
		}
	} else {
		for i, colLen := range maxColLen {
			text += SeparatorString(colLen, align[i]) + " |"
		}
	}
	return text
}

func (t *Table) Render() string {
	rows := len(t.Rows)
	cols := len(t.Headers)

	// get the maximum length of strings in each column
	// also fill a table of rendered strings from Text
	maxColLen := make([]int, cols)
	table := make([][]string, rows, cols)
	headers := make([]string, cols)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			cell := String(t.Rows[row][col])
			table[row][col] = cell
			if maxColLen[col] < len(cell) {
				maxColLen[col] = len(cell)
			}
		}
		header := String(t.Headers[col])
		headers[col] = header
		if maxColLen[col] < len(header) {
			maxColLen[col] = len(header)
		}
		maxColLen[col] += 1
	}

	text := ""
	text += fmt.Sprintln(RowString(headers, maxColLen))
	text += fmt.Sprintln(SeparatorRowString(maxColLen, t.Align))
	for _, row := range table {
		text += fmt.Sprintln(RowString(row, maxColLen))
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
	Text Text
}

func (p *Paragraph) Render() string {
	return String(p.Text)
}

type BlockQuote struct {
	Text Text
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

type MarkdownNode interface {
	Render() string
}
