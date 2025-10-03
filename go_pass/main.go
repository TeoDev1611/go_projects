package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type Conjuntos struct {
	Minusculas string
	Mayusculas string
	Digitos    string
	Simbolos   string
	Ambiguos   string // Caracteres a evitar
}

type Opcion struct {
	Nombre       string
	Seleccionado bool
}

func NuevosConjuntos() Conjuntos {
	return Conjuntos{
		Minusculas: "abcdefghijklmnopqrstuvwxyz",
		Mayusculas: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Digitos:    "0123456789",
		Simbolos:   "!@#$%^&*()-_=+[]{}|;:,.<>?",
		Ambiguos:   "1lI0Oo8B", // Evitar confusiones
	}
}

func CrearPool(usarMinus, usarMayus, usarDigitos, usarSimbolos bool, excluirAmbiguos bool) string {
	conjuntos := NuevosConjuntos()

	var pool string

	if usarMinus {
		if excluirAmbiguos {
			pool += eliminarAmbiguos(conjuntos.Minusculas, conjuntos.Ambiguos)
		} else {
			pool += conjuntos.Minusculas
		}
	}

	if usarMayus {
		if excluirAmbiguos {
			pool += eliminarAmbiguos(conjuntos.Mayusculas, conjuntos.Ambiguos)
		} else {
			pool += conjuntos.Mayusculas
		}
	}

	if usarDigitos {
		if excluirAmbiguos {
			pool += eliminarAmbiguos(conjuntos.Digitos, conjuntos.Ambiguos)
		} else {
			pool += conjuntos.Digitos
		}
	}

	if usarSimbolos {
		pool += conjuntos.Simbolos
	}

	return pool
}

func eliminarAmbiguos(conjunto, ambiguos string) string {
	resultado := ""
	for _, char := range conjunto {
		if !strings.ContainsRune(ambiguos, char) {
			resultado += string(char)
		}
	}
	return resultado
}

func indiceAleatorioSeguro(longitud int) int {
	if longitud <= 1 {
		return 0
	}

	maxValor := 256 - (256 % longitud)

	for {
		b := make([]byte, 1)
		rand.Read(b)
		valor := int(b[0])

		if valor < maxValor {
			return valor % longitud
		}
		// Si no, continuamos el loop y generamos otro
	}
}

func CrearPwsd(pool string, longitud int) string {
	caracteres := []rune(pool)
	resultado := make([]rune, longitud)
	for i := 0; i < longitud; i++ {
		indice := indiceAleatorioSeguro(len(caracteres))
		resultado[i] = caracteres[indice]
	}
	return string(resultado)
}

func menuPool() (usarMinus, usarMayus, usarDigitos, usarSimbolos, excluirAmbiguos bool) {
	opciones := []string{
		"Usar minúsculas (a-z)",
		"Usar mayúsculas (A-Z)",
		"Usar dígitos (0-9)",
		"Usar símbolos (!@#$%...)",
		"Excluir caracteres ambiguos (l,1,O,0...)",
	}

	// Opciones seleccionadas por defecto
	defaults := []string{
		"Usar minúsculas (a-z)",
		"Usar mayúsculas (A-Z)",
		"Usar dígitos (0-9)",
	}

	var seleccionadas []string

	prompt := &survey.MultiSelect{
		Message: "Selecciona las opciones para el pool de caracteres:",
		Options: opciones,
		Default: defaults,
		Help:    "Usa ESPACIO para seleccionar/deseleccionar, ENTER para confirmar",
	}

	err := survey.AskOne(prompt, &seleccionadas)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false, false, false, false, false
	}

	// Convertir selecciones a booleanos
	usarMinus = contiene(seleccionadas, opciones[0])
	usarMayus = contiene(seleccionadas, opciones[1])
	usarDigitos = contiene(seleccionadas, opciones[2])
	usarSimbolos = contiene(seleccionadas, opciones[3])
	excluirAmbiguos = contiene(seleccionadas, opciones[4])

	return
}

func contiene(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func pedirNumero() int {
	var respuesta string

	prompt := &survey.Input{
		Message: "Ingresa la longitud de tu contraseña:",
	}

	// Validar que sea un número
	err := survey.AskOne(prompt, &respuesta, survey.WithValidator(survey.Required),
		survey.WithValidator(func(ans interface{}) error {
			str := ans.(string)
			_, err := strconv.Atoi(str)
			if err != nil {
				return fmt.Errorf("debes ingresar un número válido")
			}
			return nil
		}))

	if err != nil {
		fmt.Println(err)
		return 0
	}

	numero, _ := strconv.Atoi(respuesta)
	return numero
}

func main() {
	fmt.Println(`
 ::::::::  :::::::: :::::::::    :::     ::::::::  ::::::::  
:+:    :+::+:    :+::+:    :+: :+: :+:  :+:    :+::+:    :+: 
+:+       +:+    +:++:+    +:++:+   +:+ +:+       +:+        
:#:       +#+    +:++#++:++#++#++:++#++:+#++:++#+++#++:++#++ 
+#+   +#+#+#+    +#++#+      +#+     +#+       +#+       +#+ 
#+#    #+##+#    #+##+#      #+#     #+##+#    #+##+#    #+# 
 ########  ######## ###      ###     ### ########  ########  
	`)
	usarMinus, usarMayus, usarDigitos, usarSimbolos, excluirAmbiguos := menuPool()
	pool := CrearPool(usarMinus, usarMayus, usarDigitos, usarSimbolos, excluirAmbiguos)

	// Mostrar resultados
	fmt.Println("\n📊 Resumen de selección:")
	fmt.Printf("  Minúsculas: %v\n", usarMinus)
	fmt.Printf("  Mayúsculas: %v\n", usarMayus)
	fmt.Printf("  Dígitos: %v\n", usarDigitos)
	fmt.Printf("  Símbolos: %v\n", usarSimbolos)
	fmt.Printf("  Excluir ambiguos: %v\n", excluirAmbiguos)
	fmt.Printf("\n🎯 Pool generado (%d caracteres):\n%s\n", len(pool), pool)

	longitud := pedirNumero()

	fmt.Printf("Contraseña generada: %s\n", CrearPwsd(pool, longitud))

}
