package main

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func handlerIp(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to parse client IP", err)
	}

	xForwardedFor := parseClientFromXFF(r.Header.Get("X-Forwarded-For"))
	if xForwardedFor != "" {
		ip = xForwardedFor
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ip))
}

func parseClientFromXFF(xForwardedFor string) string {
	if len(xForwardedFor) == 0 {
		log.Printf("XFF header is empty")
		return ""
	}

	ips := strings.Split(xForwardedFor, ",")

	clientIp := ips[0]
	if net.ParseIP(clientIp) == nil {
		log.Printf("XFF does not contain a valid IP")
		return ""
	}

	return clientIp
}
