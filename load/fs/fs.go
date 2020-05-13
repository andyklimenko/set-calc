package fs

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Loader struct {
	fileName string
}

func (l *Loader) Load() ([]int, error) {
	content, err := ioutil.ReadFile(l.fileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	nums := make([]int, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func New(fileName string) *Loader {
	return &Loader{
		fileName: fileName,
	}
}
