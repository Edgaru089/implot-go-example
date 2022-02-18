package itype

// Dataset describes arbitrary data stored in String-Interface{} pairs.
// This is the major way how complicated data is stored in blocks and entities.
//
// The empty interface should only hold these 4 types:
//   - int
//   - float64
//   - bool
//   - string
type Dataset map[string]interface{}
