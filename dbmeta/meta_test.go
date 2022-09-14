package dbmeta

import (
	"log"
	"testing"
)

func TestLowerCase(t *testing.T) {
	log.Println(formatFieldName("snake", "l_1_aaa"))
	log.Println(formatFieldName("snake", "l134_aaa"))
	log.Println(formatFieldName("snake", "l1232abc_aaa"))
	log.Println(formatFieldName("snake", "ldef23431_aaa"))
	log.Println(formatFieldName("snake", "l1343_1aaaaa"))
	log.Println(formatFieldName("snake", "l1343_aaa3431"))

}
