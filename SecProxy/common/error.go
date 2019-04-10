package common

import (
    "fmt"
)

var ErrorCodeMsg = make(map[int]string)

//error code
const (
    ErrInvalidRequest    = 1001
    ErrNotFoundProductId = 1002
)

func InitErrorMsg() {
    ErrorCodeMsg[ErrInvalidRequest] = "Invalid request."
    ErrorCodeMsg[ErrNotFoundProductId] = "Not find product id."
}

//err message
func ErrMsg(errCode int, err interface{}) (errMsg string) {
    if msg, ok := ErrorCodeMsg[errCode]; ok {
        errMsg = fmt.Sprintf("%s Error: %v.", msg, err)
    } else {
        errMsg = fmt.Sprintf("unknown error code [%d]. Error: %v.", errCode, err)
    }

    return 
}
    

