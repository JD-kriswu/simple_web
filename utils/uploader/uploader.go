package uploader

import (
	"onbio/logger"
	"path/filepath"

	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"go.uber.org/zap"
)

//先把图片都下载到本机的temp目录下，再上传ucloud
const (
	//uploadFile    = "./a.txt"
	configFile = "conf/ucloud.json"
	//remoteFileKey = "12345678"
)

//更新图片到ucloud上
func UploadIMGToUcloud(path string) (remoteUrl string, err error) {

	config, err := ufsdk.LoadConfig(configFile)
	if err != nil {
		panic(err.Error())
	}

	req, err := ufsdk.NewFileRequest(config, nil)
	if err != nil {
		panic(err.Error())
	}

	remoteFileName := filepath.Base(path)

	logger.Info("begin to syn pic to ucloud....")

	err = req.AsyncMPut(path, remoteFileName, "")
	if err != nil {
		logger.Error("failed reason:", zap.Error(err))
		return
	}
	logger.Info("succ sync pic to ucloud", zap.String("file", remoteFileName))
	remoteUrl = req.GetPublicURL(remoteFileName)
	return
}
