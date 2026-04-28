package parser

import (
	"fmt"
	"strconv"
)

func SkipTrash(raw []byte, pos int) int {
	for pos < len(raw) && (raw[pos] == '\n' ||
		raw[pos] == '\r' || raw[pos] == '\t' || raw[pos] == ' ') {
		pos++
	}
	return pos
}

func ParseSession(raw []byte) ([]SessionEntry, error) {
	pos := 0
	var entries []SessionEntry
	for pos < len(raw) {
		pos = SkipTrash(raw, pos)
		if pos == len(raw) {
			break
		}

	}
	return entries, nil
}

func ParseEntry(data []byte, pos int) (SessionEntry, int, error) {
	var entry SessionEntry
	for pos < len(data) {
		if data[pos] == '|' {
			return entry, pos, nil
		}
		entry.Key += string(data[pos])
	}
	return entry, pos, fmt.Errorf("Uncorrect format (| not found)")
}

func ParseValue(data []byte, pos int) (SessionValue, int, error) {
	str_len_value := ""
	var DELETE SessionValue // for compiling
	switch data[pos] {
	case 's':
		var value StringValue
		pos++
		if data[pos] == ':' {
			pos++
			for data[pos] != ':' && pos < len(data) {
				str_len_value += string(data[pos])
				pos++
			}
		}
		len_value, err := strconv.Atoi(str_len_value)
		if err != nil {
			return value, pos, fmt.Errorf("Uncorrect format with len_value")
		}
		pos++
		value.value = string(data[pos : pos+len_value-2])
		pos += len_value - 2
		if data[pos] != '"' && pos < len(data)-1 && data[pos+1] == ';' {
			return value, pos, fmt.Errorf("Uncorrect format with ;")
		}
		return value, pos, nil
	}
	return DELETE, pos, nil // FC
}
