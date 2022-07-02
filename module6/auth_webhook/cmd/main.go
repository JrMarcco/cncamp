package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	authentication "k8s.io/api/authentication/v1beta1"
)

func main() {
	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var tr authentication.TokenReview

		if err := decoder.Decode(&tr); err != nil {
			log.Println("[Error]", err.Error())
			w.WriteHeader(http.StatusBadRequest)

			if err := json.NewEncoder(w).Encode(map[string]interface{}{
				"apiVersion": "authentication.k8s.io/v1beta1",
				"kind":       "TokenReview",
				"status": authentication.TokenReviewStatus{
					Authenticated: false,
				},
			}); err != nil {
				return
			}
			return
		}
		log.Print("receving request")
		// Check User
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: tr.Spec.Token},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client := github.NewClient(tc)
		user, _, err := client.Users.Get(context.Background(), "")
		if err != nil {
			log.Println("[Error]", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(map[string]interface{}{
				"apiVersion": "authentication.k8s.io/v1beta1",
				"kind":       "TokenReview",
				"status": authentication.TokenReviewStatus{
					Authenticated: false,
				},
			}); err != nil {
				return
			}
			return
		}
		log.Printf("[Success] login as %s", *user.Login)
		w.WriteHeader(http.StatusOK)
		trs := authentication.TokenReviewStatus{
			Authenticated: true,
			User: authentication.UserInfo{
				Username: *user.Login,
				UID:      *user.Login,
			},
		}
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind":       "TokenReview",
			"status":     trs,
		}); err != nil {
			return
		}
	})
	log.Println(http.ListenAndServe(":8080", nil))
}
