package interfaces

import (
	"HoteTestBareksa/config"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

// IsLetter function to check string is aplhanumeric only
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	// Index Route
	r.GET("/", index)
	r.GET("/api/v1", index)

	// News Route
	r.GET("/api/v1/news", getAllNews)
	r.GET("/api/v1/news/:param", getNews)
	r.POST("/api/v1/news", createNews)
	r.DELETE("/api/v1/news/:news_id", removeNews)
	r.PUT("/api/v1/news/:news_id", updateNews)

	// Tags Route
	r.GET("/api/v1/tag", getAllTag)
	r.GET("/api/v1/tag/:tag_id", getTag)
	r.POST("/api/v1/tag", createTags)
	r.DELETE("/api/v1/tag/:tag_id", removeTag)
	r.PUT("/api/v1/tag/:tag_id", updateTag)

	// Topic Route
	r.GET("/api/v1/topic", getAllTopic)
	r.GET("/api/v1/topic/:topic_id", getTopic)
	r.POST("/api/v1/topic", createTopic)
	r.DELETE("/api/v1/topic/:topic_id", removeTopic)
	r.PUT("/api/v1/topic/:topic_id", updateTopic)

	// Migration Route
	r.GET("/api/v1/migrate", migrate)

	return r
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, http.StatusOK, "GO DDD API")
}

// =============================
//    MIGRATE
// =============================

func migrate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := config.DBMigrate()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	msg := "Success Migrate"
	JSON(w, http.StatusOK, msg)
}
