package samplify

import "github.com/leebenson/conform"

// Remove whitespace characters
func RemoveWhiteSpace(obj struct{}) {
     conform.Strings(&obj)
}
