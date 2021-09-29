package aop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/common"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/common/pojo"
	"ququ.im/common/utils"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var accessChannel = make(chan string, 100)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetRequestLogger() gin.HandlerFunc {

	go handleAccessChannel()

	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		var result common.Result

		// 日志格式
		if strings.Contains(c.Request.RequestURI, "/docs") {
			return
		}

		if responseBody != "" && responseBody[0:1] == "{" {
			err := json.Unmarshal([]byte(responseBody), &result)
			if err != nil {
				result = *common.Error(-1, "解析异常:"+err.Error())
			}
		}

		// 结束时间
		endTime := time.Now()

		if c.Request.Method == "POST" {
			_ = c.Request.ParseForm()
		}

		// 日志格式
		params := c.GetParamMap()
		postLog := new(pojo.PostLog)
		postLog.ID = bson.NewObjectId()
		postLog.Time = startTime.Format("2006-01-02 15:04:05")
		if strings.Contains(c.Request.RequestURI, "?") {
			postLog.Controller = c.Request.RequestURI[0:strings.Index(c.Request.RequestURI, "?")]
		} else {
			postLog.Controller = c.Request.RequestURI
		}
		if utils.Exists(params, "cvv2") {
			params["cvv2"] = utils.Base64Encode(params["cvv2"])
		}
		if utils.Exists(params, "expirydate") {
			params["expirydate"] = utils.Base64Encode(params["expirydate"])
		}
		postLog.Tradenum = params["tradenum"]
		postLog.Accountid = params["accountid"]
		postLog.Requestparam = params
		postLog.Responsetime = endTime.Format("2006-01-02 15:04:05")
		postLog.Responseparam = result
		postLog.TTL = int(endTime.UnixNano()/1e6 - startTime.UnixNano()/1e6)

		accessLog := "|" + c.Request.Method + "|" + postLog.Controller + "|" + c.ClientIP() + "|" + endTime.Format("2006-01-02 15:04:05.012") + "|" + fmt.Sprintf("%vms", endTime.UnixNano()/1e6-startTime.UnixNano()/1e6)
		logs.Debug(accessLog)
		if result.Status != 1 {
			logs.Info("请求参数:{}", params)
			logs.Info("接口返回:{}", result)
		} else {
			logs.Debug("请求参数:{}", params)
			logs.Debug("接口返回:{}", result)
		}
		accessChannel <- utils.ToJSON(postLog)
	}
}

func handleAccessChannel() {
	for accessLog := range accessChannel {
		var postLog pojo.PostLog
		json.Unmarshal([]byte(accessLog), &postLog)
		err := config.Mgo.C("JxcRequestLog").Insert(postLog)
		if err != nil {
			logs.Error("MongoDB写入错误:" + err.Error())
		}
	}
	return
}
