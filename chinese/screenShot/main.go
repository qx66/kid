package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"go.uber.org/zap"
	"net/url"
	"os"
	"path"
	"time"
)

var word string

const chineseWordDir = "/Users/qianxing/Downloads/chinese/word"

func init() {
	flag.StringVar(&word, "word", "", "-word")
}

func main() {
	flag.Parse()
	
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	
	//
	wordDir := path.Join(chineseWordDir, word)
	strokeScreenshot := path.Join(wordDir, "stroke.png")
	headerScreenshot := path.Join(wordDir, "header.png")
	fi, err := os.Stat(wordDir)
	
	if err != nil {
		
		if os.IsNotExist(err) {
			err = os.MkdirAll(wordDir, 0775)
			if err != nil {
				logger.Error("mkdir wordDir失败", zap.Error(err), zap.String("wordDir", wordDir))
				return
			}
		} else {
			logger.Error("wordDir状态失败", zap.Error(err), zap.String("wordDir", wordDir))
			return
		}
	} else {
		if !fi.IsDir() {
			err = os.MkdirAll(wordDir, 0775)
			if err != nil {
				logger.Error("mkdir wordDir失败", zap.Error(err), zap.String("wordDir", wordDir))
				return
			}
		}
	}
	
	//
	wordUrl := fmt.Sprintf("https://hanyu.baidu.com/s?wd=%s&cf=rcmd&t=img&ptype=zici", url.QueryEscape(word))
	
	opt := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("enable-logging", true),
	)
	
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opt...)
	defer cancel()
	
	// NewContent
	//ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	
	// 设置超时 -- important: 一定要放在 NewContent 之后
	ctx, cancel = context.WithTimeout(ctx, 1*60*time.Second)
	defer cancel()
	
	// 添加监听目标
	
	// 设置任务
	var stroke []byte
	var header []byte
	
	// Screenshot sel 选择 class 的名字, 需要选择 chromedp.NodeVisible  (需要使用  元素.className )
	// Screenshot sel 选择 id 的名字, 需要选择 chromedp.ById  (需要使用 #id)
	task := chromedp.Tasks{
		network.Enable(),
		chromedp.Navigate(wordUrl),
		
		// 截图 -- 一定要放在 Navigate 之后
		//chromedp.Screenshot("div.word-stroke", &stroke, chromedp.NodeVisible),
		//chromedp.Screenshot("div.word-stroke-unit", &stroke, chromedp.NodeVisible),
		chromedp.Screenshot("div.word-stroke-wrap", &stroke, chromedp.NodeVisible),
		
		chromedp.Screenshot("#word-header", &header, chromedp.ByID),
		
		
		//chromedp.Screenshot("div.header-info", &header, chromedp.NodeVisible),
	}
	
	// 执行任务
	err = chromedp.Run(ctx, task)
	if err != nil {
		logger.Error("运行chromeDp失败", zap.Error(err))
		return
	}
	
	//
	f, err := os.OpenFile(strokeScreenshot, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	
	if err != nil {
		logger.Error("打开 strokeScreenshot 文件句柄失败", zap.Error(err), zap.String("strokeScreenshot", strokeScreenshot))
		return
	}
	
	_, err = f.Write(stroke)
	if err != nil {
		logger.Error("写入 strokeScreenshot 失败", zap.Error(err), zap.String("strokeScreenshot", strokeScreenshot))
		return
	}
	
	//
	headerF, err := os.OpenFile(headerScreenshot, os.O_RDWR|os.O_CREATE, 0755)
	defer headerF.Close()
	
	if err != nil {
		logger.Error("打开 headerScreenshot 文件句柄失败", zap.Error(err), zap.String("headerScreenshot", headerScreenshot))
		return
	}
	
	_, err = headerF.Write(header)
	if err != nil {
		logger.Error("写入 headerScreenshot 失败", zap.Error(err), zap.String("headerScreenshot", headerScreenshot))
		return
	}
	
}
