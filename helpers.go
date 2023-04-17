package main

// Read ip's from txt and convert to slice, so that it can be looped
func getIpData(file string) ([]string, error) {
	ipAddresses := []string{}

	// If env is "test" use fake ip address data
	if appEnv == "test" {
		ipAddresses = []string{"127.0.0.1", "127.0.0.2", "127.0.0."}
	} else {
		// Read ip's from txt file
		ipData, err := readFileContent(file)
		if err != nil {
			return []string{}, err
		}

		// Convert string to slice, so we can loop
		ipAddresses = splitLines(ipData)
	}

	return ipAddresses, nil
}
