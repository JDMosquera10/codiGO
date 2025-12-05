función test() hacer
    repetir i desde 1 hasta 5 hacer
        si i == 3 entonces
            retornar falso
        fin
    fin
    retornar verdadero
fin

mostrar("Llamando función...")
si test() entonces
    mostrar("Retornó verdadero")
sino
    mostrar("Retornó falso")
fin
