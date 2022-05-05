package search

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	t.Log(os.Args[0]) // 执行文件所在路径
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	t.Log(pwd)
	t.Log(os.Getwd()) //源代码文件所在路径
}

func TestSubstr(t *testing.T) {
	path := "C:\\Users\\hjh\\AppData\\Local\\Temp\\GoLand"
	t.Log(path[0:strings.LastIndex(path, "\\")])
}

func TestLoadData(t *testing.T) {
	feeds, err := loadData()
	if err != nil {
		t.Log(err)
		return
	}
	for i, feed := range feeds {
		t.Log(i, feed)
	}
}

type MyFeed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func loadData() ([]*MyFeed, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	dir = dir[0:strings.LastIndex(dir, "/")]
	dir = dir + "/data/data.json"
	log.Println(dir)
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)
	var feeds []*MyFeed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
