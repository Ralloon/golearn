package strings

import (
	"regexp"
	"strings"
)

func DelExtaSpace(s string)string{
	s1 := strings.Replace(s,"	"," ",-1)
	reg,_ := regexp.Compile("\\s{2,}")
	s2 := make([]byte,len(s1))
	copy(s2,s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index)>0{
		s2 = append(s2[:spc_index[0]+1],s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}