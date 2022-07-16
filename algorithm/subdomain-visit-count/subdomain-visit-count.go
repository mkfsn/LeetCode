package subdomain_visit_count

import (
	"fmt"
	"strconv"
	"strings"
)

type domainNode struct {
	count      int
	subDomains map[string]*domainNode
}

func newDomainNode() *domainNode {
	return &domainNode{subDomains: make(map[string]*domainNode)}
}

func (d *domainNode) addCount(count int, reverseDomains []string) {
	d.addCountRecursive(count, 0, reverseDomains)
}

func (d *domainNode) addCountRecursive(count int, idx int, reverseDomains []string) {
	subDomain, ok := d.subDomains[reverseDomains[idx]]
	if !ok {
		subDomain = newDomainNode()
		d.subDomains[reverseDomains[idx]] = subDomain
	}

	subDomain.count += count
	if idx < len(reverseDomains)-1 {
		subDomain.addCountRecursive(count, idx+1, reverseDomains)
	}
}

func (d *domainNode) getCountPairDomains() []string {
	res := make([]string, 0)

	for name, subDomain := range d.subDomains {
		res = append(res, fmt.Sprintf("%d %s", subDomain.count, name))
		subDomain.traverseCount(&res, name)
	}

	return res
}

func (d *domainNode) traverseCount(result *[]string, baseDomain string) {
	for name, subdomain := range d.subDomains {
		domainName := fmt.Sprintf("%s.%s", name, baseDomain)
		*result = append(*result, fmt.Sprintf("%d %s", subdomain.count, domainName))
		subdomain.traverseCount(result, domainName)
	}
}

func subdomainVisits(cpdomains []string) []string {
	tree := newDomainNode()

	for _, cpDomain := range cpdomains {
		count, domain := splitCountPairedDomain(cpDomain)
		reverseDomains := reverseDomain(domain)
		tree.addCount(count, reverseDomains)
	}

	return tree.getCountPairDomains()
}

func splitCountPairedDomain(cpDoamin string) (int, string) {
	parts := strings.Split(cpDoamin, " ")
	count, _ := strconv.Atoi(parts[0])
	return count, parts[1]
}

func reverseDomain(domain string) []string {
	parts := strings.Split(domain, ".")
	res := make([]string, 0, len(parts))
	for i := len(parts) - 1; i >= 0; i-- {
		res = append(res, parts[i])
	}
	return res
}
