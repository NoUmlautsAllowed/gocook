package utils

//go:generate go run github.com/golang/mock/mockgen -package=utils -destination=mock_httphandler.go net/http Handler
//go:generate go run github.com/golang/mock/mockgen -package=utils -destination=mock_responsewriter.go github.com/gin-gonic/gin ResponseWriter
