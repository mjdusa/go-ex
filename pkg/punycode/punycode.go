package punycode

import (
	"fmt"

	"golang.org/x/net/idna"
)

// Wrapper for idna.Punycode to allow for seamless future updates

func WrapError(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

// ToASCII - To ASCII (decoded Punycode).
func ToASCII(val string) (*string, error) {
	ascii, err := idna.Punycode.ToASCII(val)
	if err != nil {
		return nil, WrapError("ToASCII->idna.Punycode.ToASCII", err)
	}
	return &ascii, nil
}

// ToUnicode - To Unicode (encoded Punycode).
func ToUnicode(val string) (*string, error) {
	unicode, err := idna.Punycode.ToUnicode(val)
	if err != nil {
		return nil, WrapError("ToUnicode->idna.Punycode.ToUnicode", err)
	}
	return &unicode, nil
}

// Unicode2ASCII - Convert Unicode (encoded Punycode) list to ASCII (decoded Punycode) list.
func Unicode2ASCII(unicode map[string]struct{}) map[string]struct{} {
	ascii := make(map[string]struct{})

	for elem := range unicode {
		key, err := ToASCII(elem)
		if err != nil {
			ascii[*key] = struct{}{}
		} else {
			ascii[elem] = struct{}{}
		}
	}

	return ascii
}

// ASCII2Unicode - Convert ASCII (decoded Punycode) list to Unicode (encoded Punycode) list.
func ASCII2Unicode(ascii map[string]struct{}) map[string]struct{} {
	unicode := make(map[string]struct{})

	for elem := range ascii {
		key, err := ToUnicode(elem)
		if err != nil {
			unicode[*key] = struct{}{}
		} else {
			unicode[elem] = struct{}{}
		}
	}

	return unicode
}
