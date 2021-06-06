// Package cuits contiene todo lo relacionado con los informes fiscales.
package cuits

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// CUIT representa un CUIT. Tiene como valor subyacente un int.
type CUIT int64

// Valid devuelve TRUE si el CUIT es válido. Si es cero devuelve false.
func (c CUIT) Valid() bool {
	return ValidarCUIT(int64(c))
}

// StringSinGuiones devuelve "20328896479"
func (c CUIT) StringSinGuiones() string {
	return fmt.Sprint(int(c))
}

// String evuelve "20-32889647-9"
func (c CUIT) String() string {
	if c == 0 {
		return "-"
	}
	if !c.Valid() {
		return fmt.Sprintf("CUIT INVÁLIDO: '%v'", int(c))
	}
	cs := fmt.Sprint(int(c))
	principio := cs[0:2]
	medio := cs[2:10]
	fin := cs[10:11]
	return strings.Join([]string{principio, medio, fin}, "-")
}

// New devuelve un nuevo CUIT con el numero ingresado y lo valida.
func New(cuit int64) (c CUIT, err error) {
	if ValidarCUIT(cuit) == true {
		return CUIT(cuit), nil
	}
	return CUIT(cuit), errors.New(fmt.Sprint("El CUIT ", cuit, " no es válido."))
}

// EsJuridica devuelve true si el cuit arranca con mas de 30.
func (c CUIT) EsJuridica() (es bool, err error) {
	if !c.Valid() {
		return es, errors.Errorf("cuit invalido '%v'", c)
	}
	enString := c.StringSinGuiones()
	primerosDos := enString[:2]
	enInt, err := strconv.Atoi(primerosDos)
	if err != nil {
		return false, err
	}
	if enInt >= 30 {
		return true, nil
	}
	return false, nil
}

// ExtraerDNI devuelve el DNI en el caso que se trate de una persona física
func (c CUIT) ExtraerDNI() (dni int, err error) {
	if !c.Valid() {
		return dni, errors.Errorf("el CUIT '%v' no es valido", c)
	}
	esJuridica, err := c.EsJuridica()
	if err != nil {
		return dni, err
	}
	if esJuridica {
		return dni, errors.Errorf("el CUIT %v pertenece a una persona juridica", c)
	}

	str := fmt.Sprint(int(c))[2:10]
	dni, err = strconv.Atoi(str)
	return
}

// ValidarCUIT es una función que determina si la validez del CUIT.
func ValidarCUIT(c int64) bool {
	cStr := fmt.Sprint(c)

	if len(cStr) != 11 {
		return false
	}

	d1, _ := strconv.Atoi(cStr[0:1])
	d2, _ := strconv.Atoi(cStr[1:2])
	d3, _ := strconv.Atoi(cStr[2:3])
	d4, _ := strconv.Atoi(cStr[3:4])
	d5, _ := strconv.Atoi(cStr[4:5])
	d6, _ := strconv.Atoi(cStr[5:6])
	d7, _ := strconv.Atoi(cStr[6:7])
	d8, _ := strconv.Atoi(cStr[7:8])
	d9, _ := strconv.Atoi(cStr[8:9])
	d10, _ := strconv.Atoi(cStr[9:10])
	d11, _ := strconv.Atoi(cStr[10:11])
	sumaInt := 0
	sumaInt += d1 * 5
	sumaInt += d2 * 4
	sumaInt += d3 * 3
	sumaInt += d4 * 2
	sumaInt += d5 * 7
	sumaInt += d6 * 6
	sumaInt += d7 * 5
	sumaInt += d8 * 4
	sumaInt += d9 * 3
	sumaInt += d10 * 2
	sumaInt += d11 * 1

	suma := float64(sumaInt)
	divisor := float64(11)
	if math.Mod(suma, divisor) != 0 {
		return false
	}
	return true
}
