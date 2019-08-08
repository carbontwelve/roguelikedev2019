package ui

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"io/ioutil"
)

var LoadedThemeRepository *ThemeRepository

type Theme struct {
	Name       string   `json:"name"`
	Author     string   `json:"author"`
	Color      []string `json:"color"`
	Foreground string   `json:"foreground"`
	Background string   `json:"background"`
}

func (t Theme) AsRaylibColor() []rl.Color {
	ret := make([]rl.Color, 258)

	for i, c := range t.Color {
		hex, err := t.parseHex(c)
		if err != nil {
			panic(err)
		}
		ret[i] = hex
	}

	ret[256], _ = t.parseHex(t.Background)
	ret[257], _ = t.parseHex(t.Foreground)

	return ret
}

// From https://stackoverflow.com/a/54200713
func (t Theme) parseHex(s string) (c rl.Color, err error) {
	var errInvalidFormat = errors.New("invalid format")

	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}

type ThemeRepository struct {
	path   string // theme directory path
	themes map[string]*Theme
	order  []string
	idx    uint
}

func (t ThemeRepository) GetThemeList() []string {
	return t.order
}

func (t ThemeRepository) GetCurrentTheme() Theme {
	fmt.Println(t.idx)

	return *t.themes[t.order[t.idx]]
}

func (t *ThemeRepository) SetCurrentTheme(name string) error {
	if _, ok := t.themes[name]; ok {
		for i, n := range t.order {
			if n == name {
				t.idx = uint(i)
				return nil
			}
		}
	}

	return errors.New(fmt.Sprintf("unable to set current theme to [%s]", name))
}

// Maybe Not Loop?
func (t *ThemeRepository) Next() {
	if int(t.idx) < len(t.order)-1 {
		t.idx++
	} else {
		t.idx = 0
	}
}

func (t *ThemeRepository) loadThemes() error {
	files, readDirError := ioutil.ReadDir(t.path)
	if readDirError != nil {
		return readDirError
	}

	for _, f := range files {
		theme, readError := t.readTheme(fmt.Sprintf("%s%s", t.path, f.Name()))

		if readError != nil {
			return readError
		}

		t.themes[theme.Name] = theme
		t.order = append(t.order, theme.Name)
	}

	return nil
}

func (t *ThemeRepository) readTheme(filename string) (*Theme, error) {
	file, readErr := ioutil.ReadFile(filename)

	if readErr != nil {
		return nil, readErr
	}

	data := &Theme{}

	marshalErr := json.Unmarshal([]byte(file), &data)

	if marshalErr != nil {
		return nil, marshalErr
	}

	return data, nil
}

func NewThemeRepository(path string) (*ThemeRepository, error) {
	repo := &ThemeRepository{
		path:   path,
		themes: make(map[string]*Theme),
		order:  make([]string, 0),
		idx:    0,
	}

	err := repo.loadThemes()

	if err != nil {
		return nil, err
	}

	return repo, nil
}
