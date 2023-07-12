package paymentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/Johnfor93/go-payment/config"

	"crypto/sha256"
	"encoding/base64"
	"time"
	"net/http"
	"fmt"
)

func Create(ctx *gin.Context){
	timestamp := time.Now()
	merchantid := config.envVariable("MERCHANT_ID")
	merchantkey := config.envVariable("MERCHANT_KEY")
	
	convertTime := timestamp.Format("20060102150405")

	h := sha256.New()
	h.Write([]byte(convertTime+merchantid+merchantkey))
	fmt.Printf("Time now is %v", convertTime)
	
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Printf("The SHA Encoding is %v", sha)

	ctx.JSON(http.StatusOK, gin.H{"hash": sha})
}