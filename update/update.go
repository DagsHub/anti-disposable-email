// Copyright 2020-24 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.
// Copyright 2024 DagsHub Inc. All rights reserved.

package update

import (
	"bufio"
	"context"
	"errors"
	"net/http"
)

// Update can be used to update the list of disposable email domains.
// It uses the regularly updated list found here: https://github.com/martenson/disposable-email-domains.
func Update(ctx context.Context, urls []string) (map[string]struct{}, error) {
	newList := make(map[string]struct{}, 3500)
	for _, url := range urls {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err := http.DefaultTransport.RoundTrip(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			_ = resp.Body.Close()
			return nil, errors.New("unable to fetch disposable email domains: " + resp.Status)
		}

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			newList[scanner.Text()] = struct{}{}
		}

		_ = resp.Body.Close()
		err = scanner.Err()
		if err != nil {
			return nil, err
		}
	}

	return newList, nil
}

func fetchList(ctx context.Context, url string, domains chan<- string, errs chan<- error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		errs <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errs <- errors.New("unable to fetch disposable email domains: " + resp.Status)
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		domains <- scanner.Text()
	}

	err = scanner.Err()
	if err != nil {
		errs <- err
	}
	return
}
