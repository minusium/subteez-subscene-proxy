package router

import (
	"net/http"
	"subteez/subteez"

	"github.com/gin-gonic/gin"
)

func handleDownload(c *gin.Context) {
	var request subteez.SubtitleDownloadRequest
	if c.ShouldBindJSON(&request) != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "bad request",
			},
		)
		return
	}

	url, err := subteezApi.GetDownloadLink(request)
	if err != nil {
		c.Error(err)

		if _, ok := err.(*subteez.NotFoundError); ok {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status": "not found",
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status": "server error",
				},
			)
		}

		return
	}

	subteez.ProxyFile(url, c)
}