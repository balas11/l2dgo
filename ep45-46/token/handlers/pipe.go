package handlers

import "net/http"

type HandlerPipeStage func(http.HandlerFunc) http.HandlerFunc

func Pipeline(h http.HandlerFunc, stages ...HandlerPipeStage) http.HandlerFunc {
	if len(stages) < 1 {
		return h
	}

	staged := h

	for i := len(stages) - 1; i >= 0; i-- {
		staged = stages[i](staged)
	}
	return staged
}

var ProtectedStages = []HandlerPipeStage{
	CheckAuth,
}

var AdminStages = []HandlerPipeStage{
	CheckAuth,
	CheckAdmin,
}
