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
func Update(ctx context.Context, url string) (map[string]struct{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New("unable to fetch disposable email domains: " + resp.Status)
	}

	newList := make(map[string]struct{}, 3500)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		newList[scanner.Text()] = struct{}{}
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return newList, nil
}
