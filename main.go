package main

import (
	_ "embed"
	"fmt"

	"wx_channel/cmd"
	"wx_channel/internal/application"
)

//go:embed certs/SunnyRoot.cer
var cert_file []byte

//go:embed certs/private.key
var private_key_file []byte

//go:embed lib/FileSaver.min.js
var js_file_saver []byte

//go:embed lib/jszip.min.js
var js_zip []byte

//go:embed inject/pagespy.min.js
var js_pagespy []byte

//go:embed inject/pagespy.js
var js_debug []byte

//go:embed inject/error.js
var js_error []byte

//go:embed inject/utils.js
var js_utils []byte

//go:embed inject/main.js
var js_main []byte

//go:embed inject/live.js
var js_live_main []byte

var RootCertificateName = "SunnyNet"
var AppVer = "251018_04"

func main() {
	files := &application.BizFiles{
		CertFile:       cert_file,
		PrivateKeyFile: private_key_file,
		JSFileSaver:    js_file_saver,
		JSZip:          js_zip,
		JSPageSpy:      js_pagespy,
		JSDebug:        js_debug,
		JSError:        js_error,
		JSUtils:        js_utils,
		JSMain:         js_main,
		JSLiveMain:     js_live_main,
	}
	cmd.Initialize(AppVer, RootCertificateName, files)
	if err := cmd.Execute(); err != nil {
		fmt.Printf("初始化失败 %v", err.Error())
		fmt.Printf("按 Ctrl+C 退出...\n")
		select {}
	}
}
