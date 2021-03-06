package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type byteStream []byte

// exit code
const (
	exitCodeOK int = iota
	exitCodeError
)

const (
	maxSize = 65536
	endByte = 0xff
)

var (
	cmd5LengthList = [][2]int{}
	cmd6LengthList = [][2]int{}
)

var (
	srcIndex  = 0
	destIndex = 0
)

var (
	commandLog = []byte{}
)

func main() {
	os.Exit(Run())
}

// Run program
func Run() int {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "parse error: at least input must be specified\n")
		return exitCodeError
	}

	input := flag.Arg(0)

	src, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file error: fail to load file\n")
		return exitCodeError
	}

	result := Decompress(src)
	if len(args) > 1 {
		output := flag.Arg(1)
		err := ioutil.WriteFile(output, result, os.FileMode(os.ModePerm))
		if err != nil {
			fmt.Fprintf(os.Stderr, "file error: fail to write file\n")
			return exitCodeError
		}
	} else {
		// fmt.Printf("Result: %s\n\n", byteStream(result))
	}
	fmt.Printf("Compression: %d Bytes => %d Bytes (%d%%)\n\n", len(result), len(src), 100*len(src)/len(result))
	return exitCodeOK
}

// Decompress HAL compress format
// ref: https://github.com/devinacker/exhal/blob/76b12eebd3fcd90fd3f21d427b62d9de8f5fb797/compress.c#L652
func Decompress(src []byte) []byte {
	srcIndex, destIndex = 0, 0
	decompressed := []byte{}

	for {
		insize := maxSize - srcIndex
		if insize < 1 {
			return decompressed
		}

		// read command byte from input
		input := src[srcIndex]
		srcIndex++

		// command 0xff(endByte) = end of data
		if input == endByte {
			break
		}

		// check if it is a long or regular command, get the command no. and size
		command := input >> 5
		length := (int(input&0x1f) + 1)
		if input&0xe0 == 0xe0 {
			command = (input >> 2) & 0x07
			length = ((int(input&0x03) << 8) | int(src[srcIndex])) + 1
			srcIndex++
		}

		// don't try to decompress > 64kb
		if (command == 2 && (destIndex+2*length > maxSize)) || (destIndex+length > maxSize) {
			panic("decompress size exceeds max limit: 64KB")
		}

		switch command {
		case 0: // write uncompressed bytes
			decompressed = append(decompressed, src[srcIndex:srcIndex+length]...)
			srcIndex += length
		case 1: // 8-bit RLE
			for i := 0; i < length; i++ {
				decompressed = append(decompressed, src[srcIndex])
			}
			srcIndex++
		case 2: // 16-bit RLE
			for i := 0; i < length; i++ {
				decompressed = append(decompressed, src[srcIndex])
				decompressed = append(decompressed, src[srcIndex+1])
			}
			srcIndex += 2
		case 3: // 8-bit increasing sequence
			for i := 0; i < length; i++ {
				decompressed = append(decompressed, src[srcIndex]+byte(i))
			}
			srcIndex++
		case 4, 7: // regular backref (NOTE: offset is big-endian)
			command = 4
			offset := int(uint16(src[srcIndex])<<8 | uint16(src[srcIndex+1]))

			if offset+length > maxSize {
				panic("decompress size exceeds max limit: 64KB")
			}

			for i := 0; i < length; i++ {
				decompressed = append(decompressed, decompressed[offset+i])
			}
			srcIndex += 2
		case 5: // backref with bit rotation (NOTE: offset is big-endian)
			offset := int(uint16(src[srcIndex])<<8 | uint16(src[srcIndex+1]))

			if offset+length > maxSize {
				panic("decompress size exceeds max limit: 64KB")
			}

			cmd5LengthList = append(cmd5LengthList, [2]int{offset - len(decompressed), length})
			for i := 0; i < length; i++ {
				decompressed = append(decompressed, rotate(decompressed[offset+i])) // [0x80, 0x2a, rotate(0x80), rotate(0x54)]
			}
			srcIndex += 2
		case 6: // backwards backref (NOTE: offset is big-endian)
			offset := int(uint16(src[srcIndex])<<8 | uint16(src[srcIndex+1]))

			if offset < length-1 {
				panic("offset < length-1")
			}

			cmd6LengthList = append(cmd6LengthList, [2]int{offset - len(decompressed), length})
			for i := 0; i < length; i++ {
				decompressed = append(decompressed, decompressed[offset-i]) // [1, 2, 3, 4, 4, 3, 2, 1]
			}
			srcIndex += 2
		}
		commandLog = append(commandLog, command)
	}

	printCommandLog()
	// fmt.Printf("Cmd5 length list: %v\n", cmd5LengthList)
	// fmt.Printf("Cmd6 length list: %v\n", cmd6LengthList)
	return decompressed
}

// 1000_0000 -> 0000_0001
// 0010_1010 -> 0101_0100
func rotate(b byte) byte {
	result := byte(0)
	for i := 0; i < 8; i++ {
		if b&(1<<i) != 0 {
			result |= (1 << (7 - i))
		}
	}
	return result
}

func (bs byteStream) String() string {
	builder := &strings.Builder{}
	builder.WriteString("[")
	for i, b := range bs {
		builder.WriteString(fmt.Sprintf("%d", b))
		if i < len(bs)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func printCommandLog() {
	cmdName := [7]string{
		"Trash", "RLE8", "RLE16", "Inc", "LZ", "LZR", "LZ-",
	}
	cmdCounter := [7]int{}
	for _, cmd := range commandLog {
		cmdCounter[cmd]++
	}
	result := "Commands: {"
	for i := 0; i < 7; i++ {
		result += fmt.Sprintf("%s: %d", cmdName[i], cmdCounter[i])
		if i < 6 {
			result += ", "
		} else {
			result += "}"
		}
	}
	fmt.Println(result)
}
