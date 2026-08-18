//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetOkayResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastOkayResponseDtoFromCli(c emigo.CliCastable) OkayResponseDto {
	data := OkayResponseDto{}
	return data
}
