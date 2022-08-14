package pkg

import (
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func GetenvWithFallback(key string, fallback string) string {
	temp := os.Getenv(key)
	if len(temp) == 0 {
		return fallback
	}
	return temp
}

func GetIP(r *http.Request) string {
	var userIP string
	if len(r.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = r.Header.Get("CF-Connecting-IP")
		log.Debug("CF-Connecting-IP ", net.ParseIP(userIP))
	} else if len(r.Header.Get("X-Forwarded-For")) > 1 {
		userIP = r.Header.Get("X-Forwarded-For")
		log.Debug("X-Forwarded-For ", net.ParseIP(userIP))
	} else if len(r.Header.Get("X-Real-IP")) > 1 {
		userIP = r.Header.Get("X-Real-IP")
		log.Debug("X-Real-IP ", net.ParseIP(userIP))
	} else {
		userIP = r.RemoteAddr
		if strings.Contains(userIP, ":") {
			log.Debug("RemoteAddr Host ", net.ParseIP(strings.Split(userIP, ":")[0]))
		} else {
			log.Debug(net.ParseIP(userIP))
		}
	}
	return userIP
}

// randomly sleep
func RandomSleep() {
	rand.Seed(time.Now().UnixNano())
	delta := rand.Intn(2000)
	time.Sleep((500 + time.Duration(delta)) * time.Millisecond)
	log.Debugf("Sleep %d millisecond", delta)
}
