// Package bcrypt es un algoritmo de enciptación para Go/Golang.
// Ideal para ser utilizado para generar, almacenar y validar contraseñas de
// usuarios.
// Leer:
// 	https://www.usenix.org/legacy/event/usenix99/provos/provos.pdf
// El resultado de la encriptación es un hash (un texto en base64) que puede
// ser utilizado para ser guardado directamente en una base de datos.
// Para verificar, se compara siempre el valor (la contraseña) que se ha
// recibido con el hash almacenado.
package bcrypt

import "golang.org/x/crypto/bcrypt"

// Encriptar encripta el texto ingresado. Se devuelve un hash y el error
// en caso de suceder.
func Encriptar(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(b), err
}

// Validar valida que el texto ingresado sea idéntico que el contenido del
// hash almacenado.
func Validar(s, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)); err != nil {
		return false
	}

	return true
}
