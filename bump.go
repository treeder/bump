package bump

import (
	"fmt"
	"regexp"

	"github.com/coreos/go-semver/semver"
)

// BumpInContent takes finds the first semver string in the content, bumps it, then returns the same content with the new version
func BumpInContent(vbytes []byte, part string) (old, new string, loc []int, newcontents []byte, err error) {
	re := regexp.MustCompile(`(\d+\.)?(\d+\.)?(\*|\d+)`)
	loc = re.FindIndex(vbytes)
	// fmt.Println(loc)
	if loc == nil {
		return "", "", nil, nil, fmt.Errorf("Did not find semantic version")
	}
	vs := string(vbytes[loc[0]:loc[1]])
	// fmt.Printf("vs: '%v'", vs)

	v := semver.New(vs)
	switch part {
	case "major":
		v.BumpMajor()
	case "minor":
		v.BumpMinor()
	default:
		v.BumpPatch()
	}

	len1 := loc[1] - loc[0]
	additionalBytes := len(v.String()) - len1
	// Create and fill an extended buffer
	b := make([]byte, len(vbytes)+additionalBytes)
	copy(b[:loc[0]], vbytes[:loc[0]])
	copy(b[loc[0]:loc[1]+additionalBytes], v.String())
	copy(b[loc[1]+additionalBytes:], vbytes[loc[1]:])
	// fmt.Printf("writing: '%v'", string(b))

	return vs, v.String(), loc, b, nil
}
