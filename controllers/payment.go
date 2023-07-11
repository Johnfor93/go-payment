package paymentcontroller

import (
	"github.com/gin-gonic/gin"

	"crypto/sha256"
	"encoding/base64"
	"time"
	"net/http"
	"fmt"
)

func Create(ctx *gin.Context){
	timestamp := time.Now()

	convertTime := timestamp.Format("20060102150405")

	h := sha256.New()
	h.Write([]byte(convertTime))
	fmt.Printf("Time now is %v", convertTime)

	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))

	ctx.JSON(http.StatusOK, gin.H{"hash": sha})
}