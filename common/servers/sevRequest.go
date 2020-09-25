package servers

import (
	"admin/common/common"
	"admin/util"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"time"
)
var keyPair = util.NewKeyPairAESEncrypt([]string{
	"i8MY71cS4o6b3vPK0s9TIjQr",
	"vCRUYj41Ga6xnHP52p3V0qf8",
	"PcFvVI8ikHw096UgW2M4fa13",
	"RP5vjMWX1G804Bci36uslAr2",
	"SyL568kl4FIzw1C27jn3dQJA",
	"7nf2GCW9z1abQ6803qjwJITO",
	"9KnwG8sizLhM7Ql40jO63US2",
	"lZVp3ygoKI64rSJqF15289cA",
	"ju3sU7l2BV86rRJi0cT15aNY",
	"MF6rsQ8j2IPfKX3am49Gd01n",
	"oh986aKCTOpLUHn4fF3x21c5",
	"3ec9LoiZ81zNWCJ7j2Ig0a6V",
	"AoDdy7S1w4xn93UIk8BT65El",
	"P802r1ADEWLb7seg6k43jnUR",
	"792K3r6GMnPDudo18xJ0ChvS",
	"HwyN20lK5n7tZp4W1B6Re9Uz",
	"tM34kbY6UrOfNIw875j9xFX2",
	"1B4pTXqF9d6tYGQmyk02u78J",
	"2j04y6OenZ953Vq8cXHzBmKU",
	"q96ncI2V4P5BEd8hz71GkoUF",
})

var AuthSignature = "TKkDpJ9jzZ6fu7Y8F1i2"
func SerResquest(path string,request interface{},needToken bool){

	var headers map[string]string
	//need cookie
	if headers == nil{
		headers = make(map[string]string)
	}
	if needToken{

		if access_token,isok:= util.Instance().Get("access_token");isok{
			headers["access_token"] = access_token
		}
	}
	headers["Version"] = common.GVA_VP.GetString("app.version")
	WriteSecurityHeaders(headers)



}

func WriteSecurityHeaders(headers map[string]string){
	keySeed := rand.Int()
	token := fmt.Sprintf("%s%d",AuthSignature,time.Now().Unix())
    encryptText,timestamp,err := keyPair.Encrypt(keySeed,token)
    if err != nil{
    	common.GVA_LOG.Error("write log server auth info error",zap.Any("err",err))
		return
	}
    headers["Security-Timestamp"] = fmt.Sprintf("%d",timestamp)
    headers["Security-Seed"] = fmt.Sprintf("%d",keySeed)
    headers["Security-Verify"] = encryptText
<<<<<<< HEAD
=======
    //4564654646465464
	headers["Security-Verifyaaa"] = encryptText
>>>>>>> 95ce6ba... first commit
}