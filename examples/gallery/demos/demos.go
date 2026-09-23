// Package demos holds the per-component gallery pages. Each demo_*.templ
// file converts one jughead showcase page (sites/tld/domain1/components)
// into a registry entry: it drops the links.LinkResolver coupling and
// registers a page body under the component's slug. Importing this package
// populates the gallery registry.
package demos
