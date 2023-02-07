package cmd

import "fmt"

func promptForMissingFlags(fd []*flagDetails) {
	for _, details := range fd {
		if *details.content.(*string) != "" {
			continue
		}
		for {
			fmt.Printf("%s: ", details.help)
			_, err := fmt.Scanf("%s", details.content)
			if err == nil {
				break
			}
		}
	}
}
