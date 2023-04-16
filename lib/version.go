package lib

import "io/ioutil"

func Version() string {
	b, _ := ioutil.ReadFile("../version")
	v := string(b)
	return v
}
