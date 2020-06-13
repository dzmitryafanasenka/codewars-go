package kata

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var townRew *regexp.Regexp
var monthReg *regexp.Regexp
var AllReg *regexp.Regexp

func init() {
	townRew, _ = regexp.Compile(`(?i)[a-z]*[:]`)
	monthReg, _ = regexp.Compile(`(?i)[a-z]* [0-9]*\.?[0-9]*`)
	AllReg, _ = regexp.Compile(`(?i)([a-z]*[:])|([a-z]* [0-9]*\.?[0-9]*)`)
}

func New() Rainfall {
	return Rainfall{
		Towns: map[string]Town{},
	}
}

type Rainfall struct {
	Towns map[string]Town
}

type Town map[string]float64

func (r *Rainfall) Parse(strng string) {
	values := AllReg.FindAllStringSubmatch(strng, -1)
	townName := ""
	for _, v := range values {
		value := v[0]
		if townRew.MatchString(value) {
			townName = strings.ReplaceAll(value, ":", "")
			r.Towns[townName] = Town{}
		}

		if monthReg.MatchString(value) {
			m := strings.Split(value, " ")
			floatValue, err := strconv.ParseFloat(m[1], 64)
			r.Towns[townName][m[0]] = floatValue
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (r Rainfall) GetTown(townName string) (Town, bool) {
	t, ok := r.Towns[townName]

	return t, ok
}

func (t Town) Average() float64 {
	var sum float64

	for _, v := range t {
		sum += v
	}

	return sum / float64(len(t))
}

func Mean(town string, strng string) float64 {
	r := New()

	r.Parse(strng)

	t, ok := r.GetTown(town)

	if !ok {
		return -1
	}

	return t.Average()
}

func Variance(town string, strng string) float64 {
	r := New()

	r.Parse(strng)

	t, ok := r.GetTown(town)

	if !ok {
		return -1
	}

	avg := t.Average()

	var variance float64

	for _, v := range t {
		variance += (v - avg) * (v - avg)
	}

	return variance / float64(len(t))
}
