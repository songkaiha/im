package interfaces

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/songkaiha/im/model"
	"runtime"
)

func PrintMemStats() {
	// 定义一个 runtime.MemStats对象
	var ms runtime.MemStats
	logger := models.GetLogger("MemStats")
	// 1 将内存中的数据加载到 ms对象中
	runtime.ReadMemStats(&ms)
	//2 将ms对象信息打印出来
	logger.WithFields(logrus.Fields{
		"Alloc(MB)":        fmt.Sprintf("%.2f", float64(ms.Alloc)/float64(1024*1024)),
		"TotalAlloc(MB)":   fmt.Sprintf("%.2f", float64(ms.TotalAlloc)/float64(1024*1024)),
		"Mallocs":          ms.Mallocs,
		"Frees":            ms.Frees,
		"HeapAlloc(MB)":    fmt.Sprintf("%.2f", float64(ms.HeapAlloc)/float64(1024*1024)),
		"HeapIdle(MB)":     fmt.Sprintf("%.2f", float64(ms.HeapIdle)/float64(1024*1024)),
		"HeapReleased(MB)": fmt.Sprintf("%.2f", float64(ms.HeapReleased)/float64(1024*1024)),
		"NextGC(MB)":       fmt.Sprintf("%.2f", float64(ms.NextGC)/float64(1024*1024)),
		"NumGC":            ms.NumGC,
		"NumForcedGC":      ms.NumForcedGC,
		"NumGoroutine":     runtime.NumGoroutine(),
	}).Info("sys-info")
}
