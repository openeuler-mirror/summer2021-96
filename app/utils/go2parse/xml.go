package go2parse

import (
	"io/ioutil"
	"encoding/xml"
	"bytes"
	"regexp"
	"errors"
	"io"
)

type Node struct {
	dup   bool
	key   string
	value string
	nodes []*Node
}

var charsetReader func(charset string, input io.Reader)(io.Reader, error)

func NewXml(fileName string, v interface{}) *Config {
	cfg := config()
	cfg.loadXml(fileName, v)

	return cfg
}

func (cfg *Config) loadXml (fileName string, v interface{}) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		panic(err)
	}

	if v != nil {
		err = xml.Unmarshal([]byte(data), &v)
		if err != nil {
			panic(err)
		}
	} else {
		xmlMap, err := parseXML([]byte(data))
		if err != nil{
			panic(err)
		}

		for key, value := range xmlMap {
			cfg.Values[key] = value
		}
	}
}

func XmlByteToTree(data []byte) (*Node, error) {

	re, _ := regexp.Compile("[ \t\n\r]*<")
	data   = re.ReplaceAll(data, []byte("<"))

	b := bytes.NewBuffer(data)
	p := xml.NewDecoder(b)
	p.CharsetReader = charsetReader

	node, err := xmlToTree("", nil, p)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func parseXML(data []byte) (map[string]interface{}, error) {
	var r bool

	node, err := XmlByteToTree(data)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m[node.key] = node.treeToMap(r)

	return m, nil
}

func xmlToTree(key string, a []xml.Attr, p *xml.Decoder) (*Node, error) {

	node := new(Node)
	node.nodes = make([]*Node, 0)

	if key != "" {
		node.key = key

		if len(a) > 0 {
			for _, v := range a {
				nn := new(Node)
				nn.key = `-` + v.Name.Local
				nn.value = v.Value
				node.nodes = append(node.nodes, nn)
			}
		}
	}

	for {

		token, err := p.Token()
		if err != nil {
			return nil, err
		}

		switch token.(type) {
		case xml.StartElement:
			tt := token.(xml.StartElement)
			if node.key == "" {
				node.key = tt.Name.Local
				if len(tt.Attr) > 0 {
					for _, v := range tt.Attr {
						nn := new(Node)
						nn.key = v.Name.Local
						nn.value = v.Value
						node.nodes = append(node.nodes, nn)
					}
				}
			} else {
				nn, err := xmlToTree(tt.Name.Local, tt.Attr, p)
				if err != nil {
					return nil, err
				}

				node.nodes = append(node.nodes, nn)
			}

		case xml.EndElement:
			return node, nil

		case xml.CharData:
			value := string(token.(xml.CharData))
			if len(node.nodes) > 0 {
				nn := new(Node)
				nn.value = value
				node.nodes = append(node.nodes, nn)
			} else {
				node.value = value
			}

		default:
		}
	}

	return nil, errors.New("xml parse error")
}

func (n *Node) treeToMap(r bool) interface{} {

	if len(n.nodes) == 0 {
		return n.value
	}

	m := make(map[string]interface{}, 0)
	for _, v := range n.nodes {

		if !v.dup && len(v.nodes) == 0 {
			m[v.key] = v.value
			continue
		}

		if v.dup {
			var a []interface{}
			if vv, ok := m[v.key]; ok {
				a = vv.([]interface{})
			} else {
				a = make([]interface{}, 0)
			}

			a = append(a, v.treeToMap(r))
			m[v.key] = interface{}(a)
			continue
		}

		m[v.key] = v.treeToMap(r)
	}

	return interface{}(m)
}