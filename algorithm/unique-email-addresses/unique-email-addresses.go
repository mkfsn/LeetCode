package unique_email_addresses

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	seen := make(map[string]bool)

	for _, email := range emails {
		e := trimEmail(email)
		seen[e] = true
	}

	return len(seen)
}

func trimEmail(email string) string {
	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]

	parts = strings.Split(local, "+")

	localName := strings.ReplaceAll(parts[0], ".", "")

	return localName + "@" + domain
}
