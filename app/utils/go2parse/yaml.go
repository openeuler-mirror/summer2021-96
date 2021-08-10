package go2parse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	MAP_KEY_ONLY = iota
)

type yamlReader struct {
	br       *bufio.Reader
	nodes    []interface{}
	lineNum  int
	lastLine string
}

func NewYaml(fileName string) *Config {
	cfg := config()
	cfg.loadYaml(fileName)

	return cfg
}

func (cfg *Config) loadYaml (fileName string) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		panic(err)
	}

	yamlMap, err := parseYAML([]byte(data))
	if err != nil{
		panic(err)
	}

	for key, value := range yamlMap {
		cfg.Values[key] = value
	}
}

func parseYAML(buffer []byte) (cnf map[string]interface{}, err error) {

	if len(buffer) < 3 {
		return
	}

	if string(buffer[0:1]) == "{" {
		err = json.Unmarshal(buffer, &cnf)
		if err == nil {
			return
		}
	}

	data, err := readBuffer(bytes.NewBuffer(buffer))
	if data == nil || err != nil {
		return
	}

	cnf, ok := data.(map[string]interface{})
	if !ok {
		cnf = nil
	}

	return
}

func readBuffer(r io.Reader) (interface{}, error) {

	y := &yamlReader{}
	y.br = bufio.NewReader(r)

	o, err := y.readObject(0)
	if err == io.EOF {
		err = nil
	}

	return o, err
}

func (y *yamlReader) readObject(minIndent int) (interface{}, error) {

	line, err := y.nextLine()
	if err != nil {
		if err == io.EOF && line != "" {
		} else {
			return nil, err
		}
	}

	y.lastLine = line
	indent, str := getIndent(line)
	if indent < minIndent {
		return nil, y.error("Unexpect indent", nil)
	}

	if indent > minIndent {
		minIndent = indent
	}

	switch str[0] {
	case '-':
		return y.readList(minIndent)
	case '[':
		fallthrough
	case '{':
		y.lastLine = ""

		_, value, err := y.mapKeyValue("tmp:" + str)
		if err != nil {
			return nil, y.error("err inline map or list", nil)
		}

		return value, nil
	}

	return y.readMap(minIndent)
}

func (y *yamlReader) readList(minIndent int) ([]interface{}, error) {

	list := []interface{}{}

	for {
		line, err := y.nextLine()
		if err != nil {
			return list, err
		}

		indent, str := getIndent(line)
		switch {
		case indent < minIndent:

			y.lastLine = line
			if len(list) == 0 {
				return nil, nil
			}

			return list, nil
		case indent == minIndent:

			if str[0] != '-' {
				y.lastLine = line

				return list, nil
			}

			if len(str) < 2 {
				return nil, y.error("list item is nil", nil)
			}

			key, value, err := y.mapKeyValue(str[1:])
			if err != nil {
				return nil, err
			}

			switch value {
			case nil:
				list = append(list, key)
			case MAP_KEY_ONLY:
				return nil, y.error("not support last map", nil)
			default:
				m := map[string]interface{}{key.(string): value}
				list = append(list, m)

				l, err := y.nextLine()
				if err != nil && err != io.EOF {
					return nil, err
				}

				if l == "" {
					return list, nil
				}

				y.lastLine = l

				idt, its := getIndent(line)
				if idt >= minIndent + 2 {
					switch its[0] {
					case '-':
						return nil, y.error("Unexpect", nil)
					case '[':
						return nil, y.error("Unexpect", nil)
					case '{':
						return nil, y.error("Unexpect", nil)
					}

					m, err := y.readMap(idt)
					if m != nil {
						m[key.(string)] = value
					}

					if err != nil {
						return list, err
					}
				}
			}

			continue
		default:
			return nil, y.error("bad indent " + line, nil)
		}
	}

	return nil, y.error("Impossible", nil)
}

func (y *yamlReader) readMap(minIndent int) (map[string]interface{}, error) {
	_map := map[string]interface{}{}

OUT:
	for {

		line, err := y.nextLine()
		if err != nil {
			return _map, err
		}

		indent, str := getIndent(line)
		switch {
		case indent < minIndent:

			y.lastLine = line
			if len(_map) == 0 {
				return nil, nil
			}

			return _map, nil
		case indent == minIndent:

			key, value, err := y.mapKeyValue(str)
			if err != nil {
				return nil, err
			}

			switch value {
			case nil:
				return nil, y.error("parse error", nil)
			case MAP_KEY_ONLY:

				le, err := y.nextLine()
				if err != nil {
					if err == io.EOF {
						if le == "" {
							_map[key.(string)] = nil
							return _map, err
						}
					} else {
						return nil, y.error("parse error", err)
					}
				}

				y.lastLine = le

				idt, its := getIndent(le)
				if idt < minIndent {
					return _map, nil
				}

				if idt == minIndent {
					if its[0] == '-' {
						l, err := y.readList(minIndent)
						if l != nil {
							_map[key.(string)] = l
						}

						if err != nil {
							return _map, nil
						}

						continue OUT
					} else {
						_map[key.(string)] = nil
						continue OUT
					}
				}

				obj, err := y.readObject(idt)
				if obj != nil {
					_map[key.(string)] = obj
				}

				if err != nil {
					return _map, err
				}
			default:
				_map[key.(string)] = value
			}
		default:
			return nil, y.error("bad indent", nil)
		}
	}

	return nil, y.error("parse error", nil)
}

func (y *yamlReader) nextLine() (line string, err error) {

	if y.lastLine != "" {
		line 		= y.lastLine
		y.lastLine 	= ""

		return
	}

	for {
		y.lineNum++

		line, err = y.br.ReadString('\n')
		if err != nil {
			return
		}

		if strings.HasPrefix(line, "---") || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.TrimRight(line, "\n\t\r ")
		if line == "" {
			continue
		}

		return
	}

	return
}

func (y *yamlReader) mapKeyValue(str string) (key interface{}, val interface{}, err error) {

	tokens := splitToken(str)
	key     = tokens[0]

	if len(tokens) == 1 {
		return key, nil, nil
	}

	if tokens[1] != ":" {
		return "", nil, y.error("Unexpect " + str, nil)
	}

	if len(tokens) == 2 {
		return key, MAP_KEY_ONLY, nil
	}

	if len(tokens) == 3 {
		return key, tokens[2], nil
	}

	switch tokens[2] {
	case "[":
		list := []interface{}{}

		for i := 3; i < len(tokens)-1; i++ {
			list = append(list, tokens[i])
		}

		return key, list, nil
	case "{":
		m := map[string]interface{}{}

		for i := 3; i < len(tokens)-1; i += 4 {
			if i > len(tokens)-2 {
				return "", nil, y.error("Unexpect " + str, nil)
			}

			if tokens[i+1] != ":" {
				return "", nil, y.error("Unexpect " + str, nil)
			}

			m[tokens[i].(string)] = tokens[i+2]
			if (i + 3) < (len(tokens) - 1) {
				if tokens[i+3] != "," {
					return "", nil, y.error("Unexpect " + str, nil)
				}
			} else {
				break
			}
		}

		return key, m, nil
	}

	return "", nil, y.error("Unexpect " + str, nil)
}

func (y *yamlReader) error(message string, err error) error {

	if err != nil {
		message = fmt.Sprintf("message: %s, error: %v", message, err.Error())
	}

	return errors.New(message)
}

func splitToken(str string) (tokens []interface{}) {

	str = strings.Trim(str, "\r\t\n ")
	if str == "" {
		panic("parse error")
		return
	}

	tokens = []interface{}{}
	lastPos := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case ':':
			fallthrough
		case '{':
			fallthrough
		case '[':
			fallthrough
		case '}':
			fallthrough
		case ']':
			fallthrough
		case ',':
			if i > lastPos {
				tokens = append(tokens, str[lastPos:i])
			}

			tokens = append(tokens, str[i:i+1])
			lastPos = i + 1
		case ' ':
			if i > lastPos {
				tokens = append(tokens, str[lastPos:i])
			}

			lastPos = i + 1
		case '\'':
			i++
			start := i
			for ; i < len(str); i++ {
				if str[i] == '\'' {
					break
				}
			}

			tokens = append(tokens, str[start:i])
			lastPos = i + 1
		case '"':
			i++
			start := i
			for ; i < len(str); i++ {
				if str[i] == '"' {
					break
				}
			}

			tokens = append(tokens, str[start:i])
			lastPos = i + 1
		}
	}

	if lastPos < len(str) {
		tokens = append(tokens, str[lastPos:])
	}

	if len(tokens) == 1 {
		tokens[0] = tokens[0].(string)
		return
	}

	if tokens[1] == ":" {
		if len(tokens) == 2 {
			return
		}

		if tokens[2] == "{" || tokens[2] == "[" {
			return
		}

		str = strings.Trim(strings.SplitN(str, ":", 2)[1], "\t ")
		if len(str) > 2 {
			if str[0] == '\'' && str[len(str)-1] == '\'' {
				str = str[1 : len(str)-1]
			} else if str[0] == '"' && str[len(str)-1] == '"' {
				str = str[1 : len(str)-1]
			}
		}

		val := str
		tokens = []interface{}{tokens[0], tokens[1], val}

		return
	}

	if len(str) > 2 {
		if str[0] == '\'' && str[len(str)-1] == '\'' {
			str = str[1 : len(str)-1]
		} else if str[0] == '"' && str[len(str)-1] == '"' {
			str = str[1 : len(str)-1]
		}
	}

	val := str
	tokens = []interface{}{val}

	return
}

func getIndent(str string) (int, string) {

	indent := 0
	for i, s := range str {
		switch s {
		case ' ':
			indent++
		case '\t':
			indent += 4
		default:
			return indent, str[i:]
		}
	}

	return -1, ""
}