// Programa de ejemplo: Calculadora de números primos
definir numero = 20
definir contador = 2
definir esPrimo = verdadero

función verificarPrimo(n) hacer
    si n <= 1 entonces
        retornar falso
    fin
    
    repetir i desde 2 hasta n - 1 hacer
        si n % i == 0 entonces
            retornar falso
        fin
    fin
    
    retornar verdadero
fin

mostrar("Verificando si " + numero + " es primo...")

si verificarPrimo(numero) entonces
    mostrar("El número " + numero + " es primo")
sino
    mostrar("El número " + numero + " no es primo")
fin

// Mostrar todos los primos hasta el número
mostrar("Números primos hasta " + numero + ":")
repetir i desde 2 hasta numero hacer
    si verificarPrimo(i) entonces
        mostrar(i)
    fin
fin

