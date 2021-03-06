package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shawn-ogg/machineid"
)

func main() {

	id, err := machineid.ProtectedID("Meeter")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}

	at, _ := strconv.Atoi(os.Getenv("CHANNEL_CHANGE_AT_HOUR"))
	atTime := time.Duration(at)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.String()
		if len(u) > 1 {
			u = strings.TrimSuffix(u, "/")
		}
		t := time.Now()
		t = t.Add(-time.Hour * atTime)
		day := t.Format("01-02-2006 Monday")
		hash := sha1.New()
		hash.Write([]byte(id + day))
		sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))

		fmt.Println(r.Host)
		http.Redirect(w, r, "https://meet.jit.si"+u+"_"+sha[0:7], http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
