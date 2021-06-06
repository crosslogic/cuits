package cuits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarCuit(t *testing.T) {
	if ValidarCUIT(20328896479) == false {
		t.Error("Tiró como inválido un cuit que era válido: " + "20328896479")
	}

	if ValidarCUIT(20328896478) == true {
		t.Error("Tiró como válido un cuit que era inválido: " + "20328896478")
	}
	if ValidarCUIT(2032889647) == true {
		t.Error("Tiró como válido un cuit que tenía menos dígitos: " + "2032889647")
	}
	if ValidarCUIT(0) == true {
		t.Error("Tiró como válido un cuit que era cero")
	}
}

func TestToStringSinGuiones(t *testing.T) {
	cuit, err := New(20328896479)
	assert.Nil(t, err)

	assert.Equal(t, "20328896479", cuit.StringSinGuiones())
}

func TestToStringConGuiones(t *testing.T) {
	cuit, err := New(20328896479)
	assert.Nil(t, err)
	assert.Equal(t, "20-32889647-9", cuit.String())
}

func TestEsJuridica(t *testing.T) {

	{
		cuit, err := New(20328896479)
		assert.Nil(t, err)

		es, err := cuit.EsJuridica()
		assert.Nil(t, err)
		assert.False(t, es)
	}

	{
		cuit, err := New(33693450239)
		assert.Nil(t, err)

		es, err := cuit.EsJuridica()
		assert.Nil(t, err)
		assert.True(t, es)
	}
}

func TestExtraerDNI(t *testing.T) {
	{
		cuit, err := New(20328896479)
		assert.Nil(t, err)

		dni, err := cuit.ExtraerDNI()
		assert.Nil(t, err)
		assert.Equal(t, 32889647, dni)
	}
}
