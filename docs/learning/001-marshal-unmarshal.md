# L-001: Marshal and Unmarshal in Go

Date: 2026-01-25
Category: Language
Related: L-002

## 1. Context (Why I Asked)

While implementing the markdown package for gopress, I encountered yaml.Marshal and yaml.Unmarshal in the extractMeta function. I needed to understand the core concept and when to use each approach.

## 2. The Concept

Marshal and Unmarshal are Go's standard patterns for converting between Go data structures and serialized formats (JSON, YAML, XML, etc.).

**Marshal** = Go struct → bytes (encoding/serializing)
**Unmarshal** = bytes → Go struct (decoding/deserializing)

Think of it like packing and unpacking:
- Marshal: Pack your Go data into a suitcase ([]byte) for transport
- Unmarshal: Unpack data from a suitcase ([]byte) into Go structures

## 3. Key Points

- Marshal always returns `[]byte` and `error`
- Unmarshal takes `[]byte` and a pointer to the target struct
- Struct tags control field mapping: `json:"field_name"` or `yaml:"field_name"`
- Tag `-` ignores the field entirely: `json:"-"`
- Tag `omitempty` skips zero values: `json:"count,omitempty"`
- Use Encoder/Decoder for streams (files, HTTP bodies)
- Use Marshal/Unmarshal for in-memory `[]byte` conversion

## 4. Example
```go
import "encoding/json"

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"`
}

// MARSHAL: struct → JSON bytes
user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
jsonBytes, err := json.Marshal(user)
// jsonBytes = []byte(`{"name":"Alice","email":"alice@example.com","age":30}`)

// UNMARSHAL: JSON bytes → struct
var decoded User
err := json.Unmarshal(jsonBytes, &decoded)
// decoded.Name = "Alice"

// ENCODER: write directly to io.Writer (file, http response)
f, _ := os.Create("user.json")
json.NewEncoder(f).Encode(user)

// DECODER: read directly from io.Reader (file, http request body)
json.NewDecoder(r.Body).Decode(&user)
```

## 5. When to Use / When to Avoid

**Use Marshal/Unmarshal when:**
- Converting in-memory data
- Building request bodies for HTTP clients
- Working with []byte directly

**Use Encoder/Decoder when:**
- Reading from or writing to files
- HTTP request/response bodies
- Any io.Reader or io.Writer stream

**Rule of thumb:** If you have an io.Reader or io.Writer, use Encoder/Decoder. If you have []byte, use Marshal/Unmarshal.

## 6. Sources

- AI discussion during gopress development
- Go standard library: encoding/json, encoding/xml
- gopkg.in/yaml.v3