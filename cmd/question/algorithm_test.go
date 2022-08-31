package question_test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func Test_BigNumber(t *testing.T) {
	aa := NewBigNumber("-11112222333344445555")
	print(aa)

	a := NewBigNumber("12340000000000000000")
	b := NewBigNumber("12340000000000000000")
	c := NewBigNumber("24680000000000000000")
	print(add(a, b))                      // 24680000000000000000
	print(subtract(c, b))                 // 12340000000000000000
	print(multiply(a, NewBigNumber("2"))) // 24680000000000000000
	print(divide(c, NewBigNumber("2")))   // 1234,0000,0000,0000,0000
}

type BigNumber struct {
	Sign  bool   // false: positive, true: negative
	Value [5]int // 5 group, each 0~9999
}

func NewBigNumber(s string) BigNumber {
	neg := string(s[0]) == "-"
	bn := BigNumber{}
	if neg {
		s = s[1:]
	}
	for i := len(s); i > 0; i -= 4 {
		start := i - 4
		if start < 0 {
			start = 0
		}
		strnum := s[start:i]
		n, err := strconv.ParseInt(strnum, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		bn.Value[i/5] = int(n)
	}

	if neg {
		bn = comp(bn)
	}
	return bn
}

func add(a, b BigNumber) BigNumber {
	ans := BigNumber{}
	if b.Sign {
		b = comp(b)
		return subtract(a, b)
	}

	carry := 0
	for i := len(a.Value) - 1; i >= 0; i-- {
		ans.Value[i] = a.Value[i] + b.Value[i] + carry
		if ans.Value[i] < 10000 {
			carry = 0
		} else { // 進位
			ans.Value[i] = ans.Value[i] - 10000
			carry = 1
		}
	}
	return ans
}

func subtract(a, b BigNumber) BigNumber {
	if b.Sign {
		b = comp(b)
		return add(a, b)
	}

	ans := BigNumber{}
	borrow := 0
	for i := len(a.Value) - 1; i >= 0; i-- {
		ans.Value[i] = a.Value[i] - b.Value[i] - borrow
		if ans.Value[i] >= 0 {
			borrow = 0
		} else { // 借位
			ans.Value[i] = ans.Value[i] + 10000
			borrow = 1
		}
	}
	return ans
}

func multiply(a, b BigNumber) BigNumber {
	return BigNumber{}
}

func divide(a, b BigNumber) BigNumber {
	return BigNumber{}
}

func comp(a BigNumber) BigNumber {
	for i := 0; i < len(a.Value); i++ {
		a.Value[i] = 9999 - a.Value[i]
	}

	a.Sign = !a.Sign
	a.Value[len(a.Value)-1] += 1
	return a
}

// 印出十進制大數
func print(a BigNumber) {
	out := ""
	if a.Sign {
		out += "-"
		a = comp(a)
	}
	for i := range a.Value {
		out += fmt.Sprintf("%04d", a.Value[i])
	}

	fmt.Println(out)
}

// 遞迴處理物件屬性(轉換 key name)
func Test_Tran(t *testing.T) {
	for _, val := range getTestData() {
		fmt.Println(string(camel2snake(val)))
	}
}

func camel2snake(raw []byte) []byte {
	// write your code ...
	var obj interface{}
	json.Unmarshal(raw, &obj)
	ans := recursive(obj)
	bytes, _ := json.MarshalIndent(ans, "", "  ")

	return bytes
}

func recursive(ori interface{}) interface{} {
	switch src := ori.(type) {
	case map[string]interface{}:
		nmap := make(map[string]interface{})
		for k, v := range src {
			nmap[transKey(k)] = recursive(v)
		}
		return nmap

	case []interface{}:
		nlist := make([]interface{}, 0)
		for i := range src {
			nlist = append(nlist, recursive(src[i]))
		}
		return nlist

	default:
		return ori
	}
}

func transKey(key string) string {
	reg, _ := regexp.Compile("[A-Z]")
	nkey := string(key[0])
	for i := 1; i < len(key); i++ {
		if reg.Match([]byte(string(key[i]))) {
			nkey += "_" + string(key[i])
		} else {
			nkey += string(key[i])
		}
	}
	return strings.ToLower(nkey)
}

func getTestData() [][]byte {
	return [][]byte{
		[]byte(`{
			"version": "1.0",
			"rules": [
				{
					"resource": {
						"path": "/api/data/documents"
				},
				"allowOrigins": [ "http://this.example.com", "http://that.example.com" ],
				"allowMethods": [ "GET" ],
				"allowCredentials": true
				}
			]
		}`),
		[]byte(`{
			"stringField": "ShouldNotBeChanged",
			"numberField": 0,
			"stringArray": ["ShouldNotBeChanged", "ShouldNotBeChanged"],
			"numberArray": [0, 1],
			"nestedObj1": [
				{
					"fieldName11": 0,
					"fieldName12": "ShouldNotBeChanged"
				}
			],
			"nestedObj2": {
				"fieldName21": 0,
				"fieldName22": "ShouldNotBeChanged",
				"nestedObj23": [
					{
						"fieldName231": 0,
						"fieldName232": "ShouldNotBeChanged"
					}
				]
			},
			"nestedObj3": {
				"nestedObj31": {
					"nestedObj32": {
						"fieldName33": "ShouldNotBeChanged"
					}
				}
			}
		}`),
	}
}
