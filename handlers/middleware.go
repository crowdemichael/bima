package handlers

import (
	"net/http"
	"sort"

	configs "github.com/crowdeco/bima/configs"
	events "github.com/crowdeco/bima/events"
)

type Middleware struct {
	Dispatcher  *events.Dispatcher
	Middlewares []configs.Middleware
}

func (m *Middleware) Register(middlewares []configs.Middleware) {
	sort.Slice(middlewares, func(i, j int) bool {
		return middlewares[i].Priority() > middlewares[j].Priority()
	})
	m.Middlewares = middlewares
}

func (m *Middleware) Attach(handler http.Handler) http.Handler {
	sort.Slice(m.Middlewares, func(i, j int) bool {
		return m.Middlewares[i].Priority() > m.Middlewares[j].Priority()
	})

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		for _, middleware := range m.Middlewares {
			stop := middleware.Attach(request)
			if stop {
				return
			}
		}

		m.Dispatcher.Dispatch(events.BEFORE_REQUEST, &events.Request{
			HttpRequest: request,
		})

		handler.ServeHTTP(response, request)

		m.Dispatcher.Dispatch(events.AFTER_REQUEST, &events.Response{
			ResponseWriter: response,
		})
	})
}
