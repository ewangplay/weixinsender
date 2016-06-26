package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/ewangplay/jzlconfig"
	"github.com/outmana/log4jzl"
	"jzlservice/weixinsender"
	"os"
)

var g_config jzlconfig.JZLConfig
var g_logger *log4jzl.Log4jzl

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [--config path_to_config_file]")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	//parse command line
	var configFile string
	flag.Usage = Usage
	flag.StringVar(&configFile, "config", "weixinsender.conf", "specified config filename")
	flag.Parse()

	fmt.Println("config file: ", configFile)

	//read config file
	if err := g_config.Read(configFile); err == nil {
		fmt.Println(g_config)
	}

	//init logger
	var err error
	g_logger, err = log4jzl.New("weixinsender")
	if err != nil {
		fmt.Println("Open log file fail.", err)
		os.Exit(1)
	}

	//init log level object
	g_logLevel, err = NewLogLevel()
	if err != nil {
		LOG_ERROR("创建SNSDBMgr对象失败，失败原因: %v", err)
		os.Exit(1)
	}

	//format the server listening newwork address
	var networkAddr string
	serviceIp, serviceIPIsSet := g_config.Get("service.addr")
	servicePort, servicePortIsSet := g_config.Get("service.port")
	if serviceIPIsSet && servicePortIsSet {
		networkAddr = fmt.Sprintf("%s:%s", serviceIp, servicePort)
	} else {
		networkAddr = "127.0.0.1:19090"
	}

	//startup weixinsender server
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(networkAddr)
	if err != nil {
		fmt.Println("create socket listening fail.", err)
		os.Exit(1)
	}
	handler := &WeixinSenderImpl{}
	processor := weixinsender.NewWeixinSenderProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("weixinsender server working on", networkAddr)
	g_logger.Info("weixinsender服务启动，监听地址：%v", networkAddr)

	server.Serve()
}
