package publicsuffix

import (
	"context"
	"fmt"
	"io/fs"
	"strings"
	"time"

	"github.com/mjdusa/go-ex/pkg/fileex"
	"github.com/mjdusa/go-ex/pkg/netex/httpex"
	"github.com/mjdusa/go-ex/pkg/punycode"
)

const (
	// Main list.
	MainPublicSuffixListFile = "https://publicsuffix.org/list/public_suffix_list.dat"
	// Fallback list.
	BackupPublicSuffixListFile = "https://raw.githubusercontent.com/publicsuffix/list/master/public_suffix_list.dat"
)

func WrapError(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

func ToASCII(src string) string {
	ascii, err := punycode.ToASCII(src)
	if err == nil && ascii != nil && len(*ascii) > 0 {
		return *ascii
	}

	return src
}

func ToUnicode(src string) string {
	unicode, err := punycode.ToUnicode(src)
	if err == nil && unicode != nil && len(*unicode) > 0 {
		return *unicode
	}

	return src
}

func LoadTLDs(ctx context.Context, urls []string, timeout time.Duration,
	cacheFileName string, useASCII bool) (map[string]string, error) {
	var data []byte
	var err error
	for _, url := range urls {
		data, err = httpex.HTTPGet(ctx, url, timeout)
		if err != nil {
			continue // try next URL
		}

		if len(data) > 0 { // if we have data, exit loop to the next step
			break
		}
	}

	// If data from URLs is empty, load Cache file.
	if len(data) == 0 {
		cacheData, cacheErr := ReadCacheFile(cacheFileName)
		if cacheErr != nil {
			return nil, WrapError("Unable to load TLDs", cacheErr)
		}
		data = cacheData
	}

	lines := BytesToLines(data)

	list := Lines2List(lines, useASCII)

	return list, nil
}

func ReadCacheFile(fileName string) ([]byte, error) {
	data, err := fileex.ReadAllFile(fileName)
	if err != nil {
		return nil, WrapError("ReadCacheFile->ReadAllFile", err)
	}

	return data, nil
}

func WriteCacheFile(fileName string, list map[string]string, perm fs.FileMode) error {
	values := GetListValues(list)

	if len(values) > 0 {
		data := strings.Join(values, "\n")

		err := fileex.WriteAllFile(fileName, []byte(data), perm)
		if err != nil {
			return WrapError("WriteCacheFile>WriteAllFile", err)
		}
	}

	return nil
}

// Lines2List - Add normalized lines to list with optional PunyCode useASCII.
func Lines2List(lines []string, useASCII bool) map[string]string {
	list := make(map[string]string)

	for _, line := range lines {
		key := line

		if useASCII {
			key = ToASCII(key)
		}

		list[key] = line
	}

	return list
}

// GetListKeys - get list of keys from generic map as string array.
func GetListKeys(list map[string]string) []string {
	keys := make([]string, 0, len(list))
	for key := range list {
		keys = append(keys, key)
	}
	return keys
}

// GetListValues - get list of keys from generic map as string array.
func GetListValues(list map[string]string) []string {
	values := make([]string, 0, len(list))
	for _, value := range list {
		values = append(values, value)
	}
	return values
}

// RemoveNoiseLines - remove blank lines and comments, converting each kept line to lowercase.
func RemoveNoiseLines(srcLines []string) []string {
	dstLines := []string{}
	for _, line := range srcLines {
		// Remove leading and tailing white space
		line = strings.TrimSpace(line)
		// append lines if they aren't empty or comments lines
		if line != "" && !strings.HasPrefix(line, "//") {
			dstLines = append(dstLines, line)
		}
	}
	return dstLines
}

// StringBufferToLines - split string into array of strings and remove noise lines.
func StringBufferToLines(buffer string) []string {
	lines := strings.Split(buffer, "\n")
	return RemoveNoiseLines(lines)
}

// BytesToLines - Convert byte buffer to array of strings and remove noise lines.
func BytesToLines(bytes []byte) []string {
	str := string(bytes)
	return StringBufferToLines(str)
}
