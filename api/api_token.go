package api

import (
	"encoding/base64"
	"net"
	"net/http"
	"strings"
	"time"
)

func (api *API) getToken(user string, r *http.Request) (token string, err error) {
	timestamp := time.Now().Format("20060102150405")
	phrase := user + "|" + getIPAddress(r) + "|" + timestamp
	val, err := api.secret.Encrypt([]byte(phrase))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(val), nil
}

// Pour récupérer la vraie IP du user, il faut tenir compte du fait que l'ip
// au niveau de la connexion réseau sera faussée par l'utilisation d'un
// reverse-proxy comme Traefik. Heureusement, ces derniers surchargent le
// header http avec des entrées comme X-Fowarded-For et/ou X-Real-IP pour garder
// trace de l'IP réelle.
// Donc le principe c'est de vérifier si on trouve l'une de ces entrées,
// en prenant la dernière utilisée (c'est censé être l'IP avant de rentrer dans
// notrereverse-proxy)
// Et si on n'en trouve pas, on retourne l'adresse réseau.
func getIPAddress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addr := strings.Split(r.Header.Get(h), ",")
		for i := len(addr) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addr[i])
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() {
				continue
			}
			return ip
		}
	}
	// par défaut, on renvoi l'IP réseau après avoir enlever le port
	ip := strings.Split(r.RemoteAddr, ":")
	return ip[0]
}
