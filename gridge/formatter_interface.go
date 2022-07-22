package gridge

import "fmt"

type FormatterInterface interface {
	format(s string) string
}

type Refined struct {
	handel FormatterInterface
}

func New(formatterInterface FormatterInterface) *Refined {
	return &Refined{handel: formatterInterface}
}
func (r *Refined) setFormatter(formatterInterface FormatterInterface) {
	r.handel = formatterInterface
}
func (r Refined) get() string {
	return r.handel.format("hello world")
}

type HtmlFormatter struct {
}

func (h HtmlFormatter) format(s string) string {
	return fmt.Sprintf("<b>%s</b>", s)
}

type PlainTextFormatter struct {
}

func (p PlainTextFormatter) format(s string) string {
	return s
}
