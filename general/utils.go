package General

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func Keys(m interface{}) (keys []interface{}) {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		fmt.Printf("input type not a map: %v", v)
	}

	for _, k := range v.MapKeys() {
		keys = append(keys, k.Interface())
	}
	return keys

}
func JsonViewInterface(data any) string {
	teste, _ := json.MarshalIndent(data, "", "")
	return string(teste)
}

func FindFirstMatchByID(data []map[string]interface{}, key string, value interface{}) (map[string]interface{}, bool) {
	for _, item := range data {
		// fmt.Println(key, item[key], value)
		if v, ok := item[key]; ok && ToInt(v) == ToInt(value) {

			return item, true
		}
	}
	return nil, false
}
func LogMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("🩺 💾 Memória: %.2fMB\n", float64(m.Alloc)/1024.0/1024.0)
}
func FindFirstMatchByToken(data []map[string]interface{}, key string, value interface{}) (map[string]interface{}, bool) {
	for _, item := range data {
		// fmt.Println(key, item[key], value)
		if v, ok := item[key]; ok && strings.TrimSpace(ToString(v)) == strings.TrimSpace(ToString(value)) {
			return item, true
		}
	}
	return nil, false
}
func UpdateCounter(counters *sync.Map, key string) {
	val, _ := counters.LoadOrStore(key, new(int64))
	ptr := val.(*int64)
	atomic.AddInt64(ptr, 1)
}
func AllEmpty(arr []string) bool {
	for _, v := range arr {
		if v != "" {
			return false
		}
	}
	return true
}

func sanitizeUTF8(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte{0x00}, []byte{})
}
func ToInt(v interface{}) int {
	switch n := v.(type) {
	case json.Number:
		i, err := n.Int64()
		if err != nil {
			fmt.Println("❌ Error converting to int json.Number: ", err, "value:", n)
			return 0
		}
		return int(i)
	case string:
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("❌ Error converting string:", err, "value:", n)
			return 0
		}
		return i
	case float64:
		return int(n)
	case int:
		return n
	case int64:
		return int(n)
	case int32:
		return int(n)
	case int16:
		return int(n)
	case int8:
		return int(n)
	case uint:
		return int(n)
	case uint64:
		return int(n)
	case uint32:
		return int(n)
	case uint16:
		return int(n)
	case uint8:
		return int(n)
	default:
		// fmt.Printf("❌ Unknown type (%T): %v\n", v, v)
		return 0
	}
}

// função para converter qualquer tipo para float64
func ToFloat(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case json.Number:
		i, err := v.Float64()
		if err != nil {
			fmt.Println("❌ Error converting json.Number:", err, "value:", v)
			return 0
		}
		return float64(i)
	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return f
	default:
		fmt.Printf("❌ Unknown type (%T): %v to value :\n", v, v)
		fmt.Println(val)
		return 0
	}
}
func RemoveAspasExtremas(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func ParseJSON(jsonString string) ([]map[string]interface{}, error) {
	var produtos []map[string]interface{}
	var intermediate string
	jsonString = strings.TrimSpace(jsonString)
	// Tenta primeiro deserializar diretamente em []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &produtos); err == nil {
		return produtos, nil
	}
	// Tenta deserializar como string (caso seja uma string JSON que contém um array JSON)
	if err := json.Unmarshal([]byte(jsonString), &intermediate); err != nil {
		jsonString = RemoveAspasExtremas(jsonString)
		if err := json.Unmarshal([]byte(jsonString), &produtos); err != nil {
			return nil, fmt.Errorf("erro ao desserializar para string intermediária: %w | conteúdo: %s", err, jsonString)
		} else {
			return produtos, nil
		}
	}
	if err := json.Unmarshal([]byte(intermediate), &produtos); err != nil {
		return nil, fmt.Errorf("erro ao desserializar conteúdo interno: %w | conteúdo interno: %s", err, intermediate)
	}
	return produtos, nil
}

type ProdutoID string

type Produto struct {
	ProdutoID     ProdutoID `json:"produto_id"`
	Quantidade    int       `json:"quantidade"`
	ValorUnitario float64   `json:"valor_unitario"`
}

type ProdutoInShark struct {
	ProdutoID     int     `json:"produto_id"`
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

func (p *ProdutoID) UnmarshalJSON(data []byte) error {
	// Se for "null"
	if string(data) == "null" {
		*p = ""
		return nil
	}

	// Se for número → converte pra string
	if data[0] >= '0' && data[0] <= '9' {
		*p = ProdutoID(string(data))
		return nil
	}

	// Se for string → tira aspas e guarda
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*p = ProdutoID(s)
	return nil
}
func ParseProdutosJSONStructed(jsonString string) ([]Produto, error) {
	var produtos []Produto
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	decoder.UseNumber() // <--- preserves large numbers
	if err := decoder.Decode(&produtos); err != nil {
		return nil, err
	}
	var intermediate string
	jsonString = strings.TrimSpace(jsonString)
	// Tenta primeiro deserializar diretamente em []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &produtos); err == nil {
		return produtos, nil
	}
	// Tenta deserializar como string (caso seja uma string JSON que contém um array JSON)
	if err := json.Unmarshal([]byte(jsonString), &intermediate); err != nil {
		jsonString = RemoveAspasExtremas(jsonString)
		if err := json.Unmarshal([]byte(jsonString), &produtos); err != nil {
			return nil, fmt.Errorf("erro ao desserializar para string intermediária: %w | conteúdo: %s", err, jsonString)
		} else {
			return produtos, nil
		}
	}
	if err := json.Unmarshal([]byte(intermediate), &produtos); err != nil {
		return nil, fmt.Errorf("erro ao desserializar conteúdo interno: %w | conteúdo interno: %s", err, intermediate)
	}
	return produtos, nil
}

// função para converter qualquer tipo para string
func ToString(val interface{}) string {
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%v", val)
}

func CompareMaps(m1, m2 map[string]interface{}) bool {
	for key, val1 := range m1 {
		val2, ok := m2[key]
		if !ok {
			fmt.Println("Chave ausente no segundo mapa:", key)
			return false
		}

		if !compareValues(val1, val2) {
			if key == "duracao" {
				fmt.Printf("Valores diferentes na chave '%s':\n\tmap1: %v\n\tmap2: %v\n", key, val1, val2)

			}
			return false
		}
	}
	return true
}

func compareValues(v1, v2 interface{}) bool {
	// Normalização de tipo
	v1Norm := normalize(v1)
	v2Norm := normalize(v2)
	f1, ok1 := toFloat64(v1Norm)
	f2, ok2 := toFloat64(v2Norm)
	if ok1 && ok2 {
		return almostEqual(f1, f2)
	}
	return reflect.DeepEqual(v1Norm, v2Norm)
}
func toFloat64(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case float32:
		return float64(val), true
	case float64:
		return val, true
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return f, true
		}
	}
	return 0, false
}
func almostEqual(a, b float64) bool {
	const epsilon = 0.00001
	return math.Abs(a-b) < epsilon
}
func normalize(v interface{}) interface{} {
	switch val := v.(type) {
	case string:
		// Try parse date
		if t, err := parseDate(val); err == nil {
			return t.UTC().Format("2006-01-02T15:04:05Z")
		}
		// Try parse as float
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			// Keep decimals as float
			if strings.Contains(val, ".") {
				return f
			}
			// Keep big integers as string
			return val
		}
		return strings.TrimSpace(val)

	case json.Number:
		// Try int first
		if i, err := val.Int64(); err == nil {
			return i
		}
		// Try float
		if f, err := val.Float64(); err == nil {
			return f
		}
		// Fallback keep as string
		return val.String()

	case float32, float64:
		return val // keep as float, don’t convert to string

	case int, int32, int64:
		return reflect.ValueOf(val).Int() // normalize all ints to int64

	case []interface{}:
		normalized := make([]interface{}, len(val))
		for i, item := range val {
			normalized[i] = normalize(item)
		}
		return normalized

	case map[string]interface{}:
		normalized := map[string]interface{}{}
		for k, v := range val {
			normalized[k] = normalize(v)
		}
		return normalized

	default:
		return val
	}
}

func parseDate(s string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, strings.TrimSpace(s)); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("formato inválido: %s", s)
}
