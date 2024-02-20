package bump

import (
	"fmt"
	"regexp"

	"github.com/coreos/go-semver/semver"
)

const (
	semverMatcher = `(\d+\.){1}(\d+\.){1}(\*|\d+)`
)

// BumpInContent takes finds the first semver string in the content, bumps it, then returns the same content with the new version
func BumpInContent(vbytes []byte, part string, index int) (old, new string, loc []int, newcontents []byte, err error) {
	return replace(vbytes, "", part, index)
}

// ReplaceInContent takes finds the first semver string in the content and replaces it with replaceWith
func ReplaceInContent(vbytes []byte, replaceWith string, index int) (old, new string, loc []int, newcontents []byte, err error) {
	return replace(vbytes, replaceWith, "", index)
}

func ReplaceInContent2(vbytes []byte, options *Options) (old, new string, loc []int, newcontents []byte, err error) {
	return replace2(vbytes, options)
}

func BumpInContent2(vbytes []byte, options *Options) (old, new string, loc []int, newcontents []byte, err error) {
	return replace2(vbytes, options)
}

func BumpString(input string, options *Options) (string, error) {
	old, new, loc, n2, err := replace2([]byte(input), options)
	if err != nil {
		return "", err
	}
	fmt.Println(old, new, loc, n2)
	return new, nil
}

type Options struct {
	Replace    string
	Part       string
	Index      int
	PreRelease string
	Metadata   string
}

// if index is set, it will find all matches and choose the one at the given index, -1 means last
func replace(vbytes []byte, replace, part string, index int) (old, new string, loc []int, newcontents []byte, err error) {
	options := &Options{
		Replace: replace,
		Part:    part,
		Index:   index,
	}
	return replace2(vbytes, options)
}

func replace2(vbytes []byte, options *Options) (old, new string, loc []int, newcontents []byte, err error) {
	re := regexp.MustCompile(semverMatcher)
	if options.Index == 0 {
		loc = re.FindIndex(vbytes)
	} else {
		locs := re.FindAllIndex(vbytes, -1)
		if locs == nil {
			return "", "", nil, nil, fmt.Errorf("Did not find semantic version")
		}
		locsLen := len(locs)
		if options.Index >= locsLen {
			return "", "", nil, nil, fmt.Errorf("semver index to replace out of range. Found %v, want %v", locsLen, options.Index)
		}
		if options.Index < 0 {
			loc = locs[locsLen+options.Index]
		} else {
			loc = locs[options.Index]
		}
	}
	// fmt.Println(loc)
	if loc == nil {
		return "", "", nil, nil, fmt.Errorf("Did not find semantic version")
	}
	vs := string(vbytes[loc[0]:loc[1]])

	replace := options.Replace
	if replace == "" {
		fmt.Println("bumping", vs, "part", options.Part)
		v := semver.New(vs)
		switch options.Part {
		case "major":
			v.BumpMajor()
		case "minor":
			v.BumpMinor()
		default:
			v.BumpPatch()
		}
		if options.PreRelease != "" {
			v.PreRelease = semver.PreRelease(options.PreRelease)
		}
		if options.Metadata != "" {
			v.Metadata = options.Metadata
		}
		replace = v.String()
	}

	len1 := loc[1] - loc[0]
	additionalBytes := len(replace) - len1
	// Create and fill an extended buffer
	b := make([]byte, len(vbytes)+additionalBytes)
	copy(b[:loc[0]], vbytes[:loc[0]])
	copy(b[loc[0]:loc[1]+additionalBytes], replace)
	copy(b[loc[1]+additionalBytes:], vbytes[loc[1]:])
	// fmt.Printf("writing: '%v'", string(b))

	return vs, replace, loc, b, nil
}
