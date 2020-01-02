package decotail

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
	"os/signal"
	"strings"
)

const DATETIMEFORMAT = "2006-01-02 15:04:05.000"

type TailWrapper struct {
	path      string
	timestamp bool
	keyword   string
}

func New(path string, timestamp bool, keyword string) *TailWrapper {
	return &TailWrapper{
		path:      path,
		timestamp: timestamp,
		keyword:   keyword,
	}
}

func (t *TailWrapper) Execute() error {
	tf, err := tail.TailFile(t.path, tail.Config{Follow: true})
	if err != nil {
		return err
	}

	go func() {
		for line := range tf.Lines {
			fmt.Print(t.decision(line))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	return nil
}

func (t *TailWrapper) convertToColor(text string) string {
	keywords := strings.Split(t.keyword, " ")
	for index, kw := range keywords {
		text = strings.ReplaceAll(text, kw, DefineColor(index, kw))
	}
	return text

}

func (t *TailWrapper) decision(line *tail.Line) string {
	if t.timestamp {
		return fmt.Sprintln(line.Time.Format(DATETIMEFORMAT),
			line.Time.Unix(),
			t.convertToColor(line.Text))
	}
	return fmt.Sprintln(t.convertToColor(line.Text))
}
