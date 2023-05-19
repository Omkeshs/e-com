package svcparams

const (
	ProductSVCPort = "8000"
	ProductSVCName = "product"
)

const (
	PremiumType = "Premium"
	RegularType = "Regular"
	BugetType   = "Buget"
)

var CategoryMap = map[int32]string{
	1: PremiumType,
	2: RegularType,
	3: BugetType,
}
