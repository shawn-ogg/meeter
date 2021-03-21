package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/shawn-ogg/machineid"
)

func main() {

	id, err := machineid.ProtectedID("Meeter")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.String()
		if len(u) > 1 {
			u = strings.TrimSuffix(u, "/")
		}
		t := time.Now()
		t = t.Add(-time.Hour * 6) // We change to a new sha at 3 o'clock in the morning
		day := t.Format("01-02-2006 Monday")
		hash := sha1.New()
		hash.Write([]byte(id + day))
		sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))

		fmt.Println(r.Host)
		http.Redirect(w, r, "https://meet.jit.si"+u+"_"+sha[0:7], http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
