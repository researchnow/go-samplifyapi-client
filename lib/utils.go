package samplify

import "github.com/leebenson/conform"

func RemoveWhiteSpace(obj struct{}) {
     conform.Strings(&obj)
}
