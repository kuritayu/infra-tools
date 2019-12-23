package lstar

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

func CalcCheckSumForFile(path string) string {
	b, _ := ioutil.ReadFile(path)
	return CalcCheckSum(b)
}

func CalcCheckSum(b []byte) string {
	md5val := md5.Sum(b)
	return hex.EncodeToString(md5val[:])
}
