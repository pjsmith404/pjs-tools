package main

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func handlerIp(w http.ResponseWriter, r *http.Request) {
	clientIp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to parse client IP", err)
	}

	xForwardedFor := parseClientFromXFF(r.Header.Get("X-Forwarded-For"))
	if xForwardedFor != "" {
		clientIp = xForwardedFor
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(clientIp))
}

func parseClientFromXFF(xForwardedFor string) string {
	if len(xForwardedFor) == 0 {
		log.Printf("XFF header is empty")
		return ""
	}

	ips := strings.Split(xForwardedFor, ",")
	clientIp := strings.TrimSpace(ips[0])
	for i, ip := range ips {
		log.Printf("Forwarded For IP %v: %v", i, ip)
	}
	if net.ParseIP(clientIp) == nil {
		log.Printf("clientIp is not a valid IP")
		return ""
	}

	return clientIp
}
