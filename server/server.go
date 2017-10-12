package server

import (
	"io/ioutil"
	"log"
	"net/http"

	macaron "gopkg.in/macaron.v1"
)

type Server struct {
	port int

	m *macaron.Macaron
}

func New(opts ...optFunc) (srv *Server, err error) {
	srv = new(Server)
	for _, fn := range opts {
		err = fn(srv)
		if err != nil {
			return
		}
	}

	m := macaron.New()
	m.Any("*", func(c *macaron.Context) {
		c.Query("")
		log.Println(c.Req.Request.URL)

		copyURL := c.Req.URL
		copyURL.Host = "api.telegram.org"
		copyURL.Scheme = "https"
		resp, err := http.DefaultClient.PostForm(copyURL.String(), c.Req.PostForm)

		if err != nil {
			log.Println(err)
		}
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(bts))
		c.Resp.WriteHeader(200)
		// _, err = io.Copy(c.Resp, resp.Body)
		// if err != nil {
		// 	log.Println(err)
		// }
		c.Resp.Write(bts)
	})
	srv.m = m

	return
}

func (s *Server) Run() {
	if s.port == 0 {
		s.port = 4000
	}
	s.m.Run(s.port)
}
