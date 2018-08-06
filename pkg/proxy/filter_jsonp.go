package proxy

import (
	"github.com/fagongzi/gateway/pkg/filter"
	"github.com/fagongzi/gateway/pkg/log"
	"strings"
)

// JsonpFiliter cache api result
type JsonpFiliter struct {
	filter.BaseFilter
}

func newJsonpFiliter() filter.Filter {
	return &JsonpFiliter{

	}
}

// Name return name of this filter
func (f *JsonpFiliter) Name() string {
	return FilterJsonp
}

// Pre execute before proxy
func (f *JsonpFiliter) Pre(c filter.Context) (statusCode int, err error) {
	if c.OriginRequest().IsGet() {
		callbackKey := "callback"
		callback := c.OriginRequest().QueryArgs().Peek(callbackKey)
		if len(callback) > 0 {
			c.OriginRequest().SetUserValue(filter.UsingJsonpValue, callback)
			c.ForwardRequest().SetRequestURI(strings.Replace(string(c.ForwardRequest().RequestURI()), callbackKey, "____", -1))
			//c.SetAttr(filter.UsingJsonpValue, string(callback))
			log.Infof("request %s callback %s", string(c.OriginRequest().RequestURI()), string(callback))
		}
	}
	return f.BaseFilter.Post(c)
}

// Post execute after proxy
func (f *JsonpFiliter) Post(c filter.Context) (statusCode int, err error) {
	//var buffer bytes.Buffer
	//buffer.Write(b1)
	//
	//buffer.Write(b2)
	//
	//b3 :=buffer.Bytes()
	//
	//
	//jsonpBody := c.GetAttr(filter.UsingJsonpValue) + "(" + string(c.Response().Body()) + ")"
	//c.Response().SetBody(jsonpBody)
	//if c.DispatchNode().Cache == nil {
	//	return f.BaseFilter.Post(c)
	//}
	//
	//matches, id := getCachingID(c)
	//if !matches {
	//	return f.BaseFilter.Post(c)
	//}
	//
	//f.cache.Add(id, genCachedValue(c))
	//f.tw.Schedule(time.Second*time.Duration(c.DispatchNode().Cache.Deadline),
	//	f.removeCache, id)
	return f.BaseFilter.Post(c)
}
