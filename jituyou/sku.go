package sku

import ()

type SKUCode string

func (c SKUCode) Invalid() bool {
	return false
}

func (c SKUCode) ItemCD() string {
	return string(c[0:5])
}

func (c SKUCode) SizeCD() string {

	return string(c[5:7])
}

func (c SKUCode) ColorCD() string {
	return string(c[7:9])
}
