package decotail

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"github.com/hpcloud/tail"
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

	for line := range tf.Lines {
		fmt.Print(t.decision(line))
		//TODO シグナル受信時の挙動
		//シグナル受信したらループを抜けて、そのままnil返却
	}

	return nil
}

func (t *TailWrapper) convertToColor(text string) string {
	//TODO 複数キーワードに対応する
	if t.keyword != "" && strings.Contains(text, t.keyword) {
		return strings.ReplaceAll(text, t.keyword, fmt.Sprint(brush.Red(t.keyword)))
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
