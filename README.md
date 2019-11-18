# packet

## Installation
This package is Go Modules compatible. You can just use `go get` to install it
### install v1
`go get github.com/qapquiz/packet`
### install v2
`go get github.com/qapquiz/packet/v2`

I recommend to use packet v2 only. But for compatible I will keep version 1 to older user.

## Usage

### Write Packet
```
w := NewWriter()

idx := 0
idx += w.WriteString(a.Name, idx)
idx += w.WriteInt64(a.BirthDay.Unix(), idx)
idx += w.WriteString(a.Phone, idx)
idx += w.WriteInt32(int32(a.Siblings), idx)
idx += w.WriteBoolean(a.Spouse, idx)
idx += w.WriteFloat64(a.Money, idx)

return w.Bytes()
```

### Read Packet
```
r := NewReader(in)

a.Name = r.ReadString()
a.BirthDay = time.Unix(r.ReadInt64(), 0)
a.Phone = r.ReadString()
a.Siblings = int(r.ReadInt32())
a.Spouse = r.ReadBoolean()
a.Money = r.ReadFloat64()
```

## Benchmark v1
```
pkg: github.com/qapquiz/packet
BenchmarkMarshalA-8     	 2278370	     508 ns/op	     200 B/op	      12 allocs/op
BenchmarkUnmarshalA-8   	 2977406	     401 ns/op	     144 B/op	      12 allocs/op
```

## Benchmark v2
```
pkg: github.com/qapquiz/packet/v2
BenchmarkMarshalA-8     	 5070519	     228 ns/op	     120 B/op	       4 allocs/op
BenchmarkUnmarshalA-8   	17714499	    67.0 ns/op	      24 B/op	       2 allocs/op
```
