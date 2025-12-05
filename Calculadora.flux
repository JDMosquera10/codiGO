definir num1 = 30
definir num2 = 20

función Sumar(a, b) hacer
    retornar a + b
fin

función Restar(a, b) hacer
    retornar a - b
fin

función Multiplicar(a, b) hacer
    retornar a * b
fin

función Dividir(a, b) hacer
    si b <= 0 entonces
        retornar "No se puede dividir entre 0"
    fin

    retornar a / b
fin

función Modulo(a, b) hacer
    si b <= 0 entonces
        retornar "No se puede encontrar el modulo entre 0"
    fin
    retornar a % b
fin

mostrar("Resultado de sumar: " + Sumar(num1, num2))
mostrar("Resultado de restar: " + Restar(num1, num2))
mostrar("Resultado de multiplicar: " + Multiplicar(num1, num2))
mostrar("Resultado de dividir: " + Dividir(num1, num2))
mostrar("Resultado del modulo: " + Modulo(num1, num2))