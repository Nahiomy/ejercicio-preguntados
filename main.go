package main

import "fmt"

func main() {

	puntuacion := 0
	primera_respuesta := 1
	segunda_respuesta := 3
	tercera_respuesta := 3
	cuarta_respuesta := 2
	quinta_respuesta := 2
	sexta_respuesta := 4
	septima_respuesta := 3
	octava_respuesta := 2
	novena_respuesta := 4
	decima_respuesta := 1

	fmt.Println("Bienvenido, nuevo usuario, ¿Cuál es tu nombre?")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("Comencemos!!!")

	primera_pregunta := `¿Cuál es el río más caudaloso del mundo?
		1. Río Amazonas
		2. Río Nilo
		3. Río misisipi
		4. Río Orinoco`

	segunda_pregunta := `¿Cuándo se estrenó la película "Harry Potter y la piedra filosofal"
		1. 2000
		2. 1999
		3. 2001
		4. 2002`

	tercera_pregunta := `¿Cómo queda un maguito después de comer?
		1. Mago conento
		2. Mago relleno
		3. Magordito
		4. Magordote`

	cuarta_pregunta := `¿Cuál fue la primera pelicula de disney?
		1. Cenicienta
		2. Blanca Nieves y los siete enanos
		3. Zootopia
		4. Mickey Mouse`

	quinta_pregunta := `¿qué clavó Pablito?
		1. El clavito que Pablito clavó
		2. Pablito clavó un clavito
		3. Pablito es muy pequeño para clavar un clavito
		4. Pablito no tenía un martillo para clavar un clavito`

	sexta_pregunta := `¿Cuándo salió el primer episodio de Pokémon?
		1. 1998
		2. 1992
		3. 2001
		4. 1997`

	septima_pregunta := `¿De dónde son originarios los juegos olímpicos?
		1. Asia
		2. Europa
		3. Antigua Grecia
		4. Italia`

	octava_pregunta := `¿La ballena qué tipo de animal es?
		1. Mamífero
		2. Mamífero marino
		3. Invertebrado
		4. Vertebrado`

	novena_pregunta := `¿Cómo se llama la capital de Mongolia?
		1. Mongólica
		2. Mongol
		3. Ciudad de Mongolia
		4. Ulan Bator`

	decima_pregunta := `¿En qué país se usó la primera bomba atómica en combate?
		1. Hiroshima
		2. Tokyo
		3. Estocolmo
		4. Nagasaki`

	fmt.Println(primera_pregunta)

	_, err := fmt.Scanf("%d", &primera_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if primera_respuesta == 1 {
		puntuacion++
		_, err := fmt.Println("Excelente trabajo!!! \n Tienes ", puntuacion, "puntos")
		if err != nil {
			panic(err)
		}
	} else {
		puntuacion = 0
		fmt.Println("Respuesta Incorrecta!!! \n Tienes", puntuacion, "puntos")
	}

	fmt.Println(segunda_pregunta)

	_, err = fmt.Scanf("%d", &segunda_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if segunda_respuesta == 3 {
		puntuacion++
		_, err := fmt.Println("Excelente trabajo!!! \n Tienes ", puntuacion, "puntos")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes", puntuacion, "puntos")
	}

	fmt.Println(tercera_pregunta)

	_, err = fmt.Scanf("%d", &tercera_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if tercera_respuesta == 3 {
		puntuacion++
		_, err := fmt.Println("Excelente trabajo!!! \n Tienes ", puntuacion, "puntos")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes", puntuacion, "puntos")
	}

	fmt.Println(cuarta_pregunta)

	_, err = fmt.Scanf("%d", &cuarta_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if cuarta_respuesta == 2 {
		puntuacion++
		_, err := fmt.Println("Excelente trabajo!!! \n Tienes ", puntuacion, "puntos")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes", puntuacion, "puntos")
	}

	fmt.Println(quinta_pregunta)

	_, err = fmt.Scanf("%d", &quinta_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if quinta_respuesta == 2 {
		puntuacion++
		_, err := fmt.Println("Excelente trabajo!!! \n Tienes ", puntuacion, "puntos")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes", puntuacion, "puntos")
	}

	fmt.Println(sexta_pregunta)

	_, err = fmt.Scanf("%d", &sexta_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if sexta_respuesta == 4 {
		puntuacion++
		_, err = fmt.Println("Fantástico!!! \n tienes ", puntuacion, "puntos")
		if err != nil {
			fmt.Println("ERROR DETECTADO")
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes ", puntuacion, "puntos")
	}

	fmt.Println(septima_pregunta)

	_, err = fmt.Scanf("%d", &septima_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if septima_respuesta == 3 {
		puntuacion++
		_, err = fmt.Println("Qué comes que adivinas!!! \n tienes ", puntuacion, "puntos")
		if err != nil {
			fmt.Println("ERROR DETECTADO")
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes ", puntuacion, "puntos")
	}

	fmt.Println(octava_pregunta)

	_, err = fmt.Scanf("%d", &octava_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if octava_respuesta == 2 {
		puntuacion++
		_, err = fmt.Println("Qué comes que adivinas!!! \n tienes ", puntuacion, "puntos")
		if err != nil {
			fmt.Println("ERROR DETECTADO")
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes ", puntuacion, "puntos")
	}

	fmt.Println(novena_pregunta)

	_, err = fmt.Scanf("%d", &novena_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if novena_respuesta == 4 {
		puntuacion++
		_, err = fmt.Println("Respuesta Correcta!!! \n tienes ", puntuacion, "puntos")
		if err != nil {
			fmt.Println("ERROR DETECTADO")
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tienes ", puntuacion, "puntos")
	}

	fmt.Println(decima_pregunta)

	_, err = fmt.Scanf("%d", &decima_respuesta)
	if err != nil {
		fmt.Println("ERROR DETECTADO")
		panic(err)
	}

	if decima_respuesta == 1 {
		puntuacion++
		_, err = fmt.Println("Comiste y adivinaste!!! \n Tu puntiación final es de ", puntuacion, "puntos")
		if err != nil {
			fmt.Println("ERROR DETECTADO")
			panic(err)
		}
	} else {
		fmt.Println("Respuesta Incorrecta!!! \n Tu puntiación final es de ", puntuacion, "puntos")
	}

}
