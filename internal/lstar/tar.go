package lstar

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"
)

type Tar struct {
	permission string
	owner      string
	group      string
	size       int64
	date       string
	path       string
	md5        string
	Elements   []*Element
}

type Element struct {
	permission string
	owner      string
	group      string
	size       int64
	date       string
	path       string
	data       []byte
	md5        string
}

func Setup(path string) *Tar {
	// Open FIle
	file, _ := os.Open(path)

	// Open Tar
	var tarData *tar.Reader
	e := filepath.Ext(path)
	if e == ".gz" {
		gz, _ := gzip.NewReader(file)
		tarData = tar.NewReader(gz)
	} else {
		tarData = tar.NewReader(file)
	}

	var collection []*Element

	for {
		header, err := tarData.Next()
		if err == io.EOF {
			break
		}

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, tarData)

		e := &Element{
			permission: header.FileInfo().Mode().String(),
			owner:      header.Uname,
			group:      header.Gname,
			size:       header.Size,
			date:       header.ModTime.Format("2006-01-02 15:04:05"),
			path:       header.Name,
			data:       buf.Bytes(),
		}

		e.md5 = CalcCheckSum(e.data)
		collection = append(collection, e)
	}

	// Set Struct
	info, _ := file.Stat()
	var u string
	var g string
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		u = ResolveUser(int(stat.Uid))
		g = ResolveGroup(int(stat.Gid))
	}

	t := &Tar{
		permission: info.Mode().String(),
		owner:      u,
		group:      g,
		size:       info.Size(),
		date:       info.ModTime().Format("2006-01-02 15:04:05"),
		path:       path,
		md5:        CalcCheckSumForFile(path),
		Elements:   collection,
	}
	return t
}

func (t *Tar) getTarName() string {
	return filepath.Base(t.path)
}

func (t *Tar) getCheckSum() string {
	return t.md5
}

func (t *Tar) getArchivedFileName() []string {
	var result []string
	for _, v := range t.Elements {
		result = append(result, v.path)
	}
	return result
}

func (t *Tar) getArchivedChecksum() []string {
	var result []string
	for _, v := range t.Elements {
		result = append(result, v.md5)
	}
	return result
}

func (t *Tar) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', 0)
	_, _ = w.Write([]byte(strings.Join([]string{
		"Permission",
		"Owner",
		"Group",
		"Size",
		"Date",
		"Path",
		"Checksum"},
		"\t") + "\n"))

	_, _ = w.Write([]byte(strings.Join([]string{
		t.permission,
		t.owner,
		t.group,
		strconv.Itoa(int(t.size)),
		t.date,
		t.path,
		t.md5},
		"\t") + "\n"))

	for _, v := range t.Elements {
		_, _ = w.Write([]byte(strings.Join([]string{
			v.permission,
			v.owner,
			v.group,
			strconv.Itoa(int(v.size)),
			v.date,
			v.path,
			v.md5}, "\t") + "\n"))
	}
	_ = w.Flush()

	fmt.Println("")
}
