package main

import "time"

type ResponseSuccess struct {
	Kind             string        `json:"kind"`
	Etag             string        `json:"etag"`
	Summary          string        `json:"summary"`
	Updated          time.Time     `json:"updated"`
	TimeZone         string        `json:"timeZone"`
	AccessRole       string        `json:"accessRole"`
	DefaultReminders []interface{} `json:"defaultReminders"`
	NextSyncToken    string        `json:"nextSyncToken"`
	Items            []struct {
		Kind        string    `json:"kind"`
		Etag        string    `json:"etag"`
		ID          string    `json:"id"`
		Status      string    `json:"status"`
		HTMLLink    string    `json:"htmlLink"`
		Created     time.Time `json:"created"`
		Updated     time.Time `json:"updated"`
		Summary     string    `json:"summary"`
		Description string    `json:"description"`
		Creator     struct {
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
			Self        bool   `json:"self"`
		} `json:"creator"`
		Organizer struct {
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
			Self        bool   `json:"self"`
		} `json:"organizer"`
		Start struct {
			Date string `json:"date"`
		} `json:"start"`
		End struct {
			Date string `json:"date"`
		} `json:"end"`
		Transparency string `json:"transparency"`
		Visibility   string `json:"visibility"`
		ICalUID      string `json:"iCalUID"`
		Sequence     int    `json:"sequence"`
		EventType    string `json:"eventType"`
	} `json:"items"`
}

type ResponseError struct {
	Error struct {
		Errors []struct {
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
			Message string `json:"message"`
		} `json:"errors"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Holiday struct {
	Title string `json:"title"`
	Date  string `json:"date"`
	// Day   string `json:"day"`
}