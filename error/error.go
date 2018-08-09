package error

import (
	"strings"
)

func IsNeedWaitError(err error) bool {
	return checkError(err,"Need waited.")
}

func IsConnectionRefuseError(err error) bool{
	return checkError(err,"connection refuse")
}

func IsUnexceptedEOFError(err error) bool {
	return checkError(err,"unexpected error")
}

func checkError(error error,str string) bool {
	erstr := error.Error()
	return strings.Contains(erstr,str)
}
