package cmp

import (
	"fmt"
	"net/url"
	"strings"
)

type Companies map[string][]Company

type Company struct {
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	CareerPageURL string   `json:"career_page_url"`
	Type          string   `json:"type,omitempty"`
	RemotePolicy  string   `json:"remote_policy,omitempty"`
	HiringPolicy  string   `json:"hiring_policy,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

func (c Company) GetTagsString() string {
	return strings.Join(c.Tags, " - ")
}

func (c Company) Validate() error {
	if len(c.Tags) > 20 {
		return fmt.Errorf("Company tags must be less than 20")
	}

	_, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return fmt.Errorf("Company URL is not valid. ", err)
	}

	_, err = url.ParseRequestURI(c.CareerPageURL)
	if err != nil {
		return fmt.Errorf("Company Career Page URL is not valid. ", err)
	}
	return nil
}
