package ip

import (
	"io"
	"log"
	"net/http"
)

func PublicIP() string {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	ip := string(body)
	return ip
}
