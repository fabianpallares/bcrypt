package bcrypt

import (
	"fmt"
	"testing"
)

func TestEncriptar(t *testing.T) {
	clave := "EstaEsUnaClave"
	fmt.Println("Clave a encriptar:", clave)
	fmt.Println()
	for i := 0; i < 10; i++ {
		hash, _ := Encriptar(clave)
		fmt.Printf("La clave encriptada es: %v. Es válida: %v\n", hash, Validar(clave, hash))
	}
}
