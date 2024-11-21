package api

import "regexp"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func ValidateEmail(email string) bool {
	regexEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return regexEmail.MatchString(email)
}

func ValidateOrg(name string) bool {
	regexName := regexp.MustCompile("^([А-Я]{3})[\"']([А-Я][а-яё]+[\"'])$")
	return regexName.MatchString(name)
}
