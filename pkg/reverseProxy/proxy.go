package reverseproxy

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Conf struct {
	ProxyUrl string
}

// 做简单的转发操作
func Proxy(c *gin.Context, conf *Conf) {
	err := setTokenToUrl(c.Request.URL, conf)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("填写的地址有误: %s", err.Error()))
		c.Abort()
		return
	}

	req, err := http.NewRequestWithContext(c, c.Request.Method, c.Request.URL.String(), c.Request.Body)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	defer req.Body.Close()
	req.Header = c.Request.Header

	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	// header 也带过来
	for k := range resp.Header {
		for j := range resp.Header[k] {
			c.Header(k, resp.Header[k][j])
		}
	}
	extraHeaders := make(map[string]string)
	extraHeaders["direct"] = "lab"
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, extraHeaders)
	c.Abort()
}

func setTokenToUrl(rawUrl *url.URL, conf *Conf) error {
	u, err := url.Parse(conf.GetProxyUrl())
	if err != nil {
		return err
	}

	rawUrl.Scheme = u.Scheme
	rawUrl.Host = u.Host
	ruq := rawUrl.Query()
	rawUrl.RawQuery = ruq.Encode()
	return nil
}

func (conf *Conf) GetProxyUrl() string {
	if conf.ProxyUrl == "" {
		conf.ProxyUrl = "https://poto.alphatech.mobi"
	}
	return conf.ProxyUrl
}
