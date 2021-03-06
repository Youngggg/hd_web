package util

import (
	"context"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func GetToken() string {
	accessKey, _ := beego.AppConfig.String("accessKey")
	secretKey, _ := beego.AppConfig.String("secretKey")
	bucket, _ := beego.AppConfig.String("bucket")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// putPolicy.Expires = 7200 //示例2小时有效期，默认1小时
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func UploadFile(filePath, key string) (storage.PutRet, error) {
	upToken := GetToken()
	cfg := storage.Config{}
	// 空间对应的机房
	// cfg.Zone = &storage.ZoneHuadong
	// // 是否使用https域名
	// cfg.UseHTTPS = false
	// // 上传是否使用CDN上传加速
	// cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		// Params: map[string]string{
		//     "x:name": "github logo",
		// },
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
