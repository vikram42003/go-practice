// Package twofer consists of response functions for when you get a cookie from the bakery.
package twofer

// ShareWith returns a string response based on whether name was provided or not.
func ShareWith(name string) string {
	if name != "" {
        return "One for " + name + ", one for me."
    } else {
        return "One for you, one for me."
    }
}
