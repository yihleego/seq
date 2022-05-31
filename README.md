# Seq

A globally, 64 bits, thread-safe identifier for Go. It can generate `4,194,303` numbers per second.

## Introduction

```
┌--------┬--------┬--------┬--------┬--------┬--------┬--------┬--------┐
|11111111|11111111|11111111|11111111|11111111|11111111|11111111|11111111| FORMAT: 64 bits
├--------┼--------┼--------┼--------┼--------┼--------┼--------┼--------┤
|XXXXXXXX|XXXXXXXX|XXXXXXXX|XXXXXXXX|        |        |        |        | TIMESTAMP: 32 bits
├--------┼--------┼--------┼--------┼--------┼--------┼--------┼--------┤
|        |        |        |        |XXXXXXXX|XX      |        |        | WORKER ID: 10 bits
├--------┼--------┼--------┼--------┼--------┼--------┼--------┼--------┤
|        |        |        |        |        |  XXXXXX|XXXXXXXX|XXXXXXXX| SEQUENCE: 22 bits
└--------┴--------┴--------┴--------┴--------┴--------┴--------┴--------┘
```

## Usage

### New Seq

Create a new Seq with the worker identifier, and the worker identifier should be between `0` and `1023`.

```go
seq := NewSeq(0)
v := seq.Next()
```

### Random Seq

Create a new Seq with a random worker identifier.

```go
seq := RandomSeq()
v := seq.Next()
```

## License

This project is under the MIT license. See the [LICENSE](LICENSE) file for details.
