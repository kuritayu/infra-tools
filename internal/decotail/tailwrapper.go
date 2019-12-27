package decotail

import (
	"fmt"
	"github.com/hpcloud/tail"
)

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
		//TODO timestampフラグ対応
		//TODO keywordフラグ対応
		fmt.Println(line.Text)
		//TODO シグナル受信時の挙動
		//シグナル受信したらループを抜けて、そのままnil返却
	}

	return nil
}
