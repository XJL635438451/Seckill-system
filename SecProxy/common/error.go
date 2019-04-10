package common

//product status
const (
    ProductStatusNormal       = 0
    ProductStatusSaleOut      = 1
    ProductStatusForceSaleOut = 2
)

//error code
const (
    ErrInvalidRequest      = 1001
    ErrNotFoundProductId   = 1002
    ErrUserCheckAuthFailed = 1003
    ErrUserServiceBusy     = 1004
    ErrActiveNotStart      = 1005
    ErrActiveAlreadyEnd    = 1006
    ErrActiveSaleOut       = 1007
    ErrProcessTimeout      = 1008
    ErrClientClosed        = 1009
)
