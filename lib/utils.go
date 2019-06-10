package samplify

import "github.com/leebenson/conform"

// RemoveWhiteSpace function is trimming the object
func RemoveWhiteSpace(obj struct{}) {
     conform.Strings(&obj)
}
