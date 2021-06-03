package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(1, 1)

	if total != 2 {
		t.Errorf("Resultado da soma eh invalido: Resultado %d. Esperado %d", total, 2)
	}
}
