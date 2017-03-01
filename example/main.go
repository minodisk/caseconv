package main

import (
	"fmt"

	"github.com/minodisk/caseconv"
)

func main() {
	fmt.Println(caseconv.LowerCamelCase(`Obi-Wan "Ben" Kenobi`)) // -> "obiWanBenKenobi"
	fmt.Println(caseconv.UpperCamelCase(`Obi-Wan "Ben" Kenobi`)) // -> "ObiWanBenKenobi"
	fmt.Println(caseconv.SnakeCase(`Obi-Wan "Ben" Kenobi`))      // -> "Obi_Wan_Ben_Kenobi"
	fmt.Println(caseconv.LowerSnakeCase(`Obi-Wan "Ben" Kenobi`)) // -> "obi_wan_ben_kenobi"
	fmt.Println(caseconv.UpperSnakeCase(`Obi-Wan "Ben" Kenobi`)) // -> "OBI_WAN_BEN_KENOBI"
	fmt.Println(caseconv.Hyphens(`Obi-Wan "Ben" Kenobi`))        // -> "Obi-Wan-Ben-Kenobi"
	fmt.Println(caseconv.LowerHyphens(`Obi-Wan "Ben" Kenobi`))   // -> "obi-wan-ben-kenobi"
	fmt.Println(caseconv.UpperHyphens(`Obi-Wan "Ben" Kenobi`))   // -> "OBI-WAN-BEN-KENOBI"
}
