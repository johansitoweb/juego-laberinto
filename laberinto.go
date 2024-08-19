package main

import (
    "fmt"
    "os"
)

const (
    wall       = '#'
    path       = ' '
    player     = 'P'
    exit       = 'E'
)

var maze = [][]rune{
    {wall, wall, wall, wall, wall, wall, wall, wall},
    {wall, path, path, wall, path, path, path, wall},
    {wall, path, wall, wall, wall, path, wall, wall},
    {wall, path, path, path, wall, path, exit, wall},
    {wall, wall, wall, path, wall, wall, wall, wall},
    {wall, wall, wall, path, path, path, path, wall},
    {wall, wall, wall, wall, wall, wall, wall, wall},
}

var playerPos = [2]int{1, 1} // posición inicial del jugador

func printMaze() {
    for i, row := range maze {
        for j := range row {
            if i == playerPos[0] && j == playerPos[1] {
                fmt.Print(string(player))
            } else {
                fmt.Print(string(row[j]))
            }
        }
        fmt.Println()
    }
}

func movePlayer(direction string) {
    newPos := playerPos

    switch direction {
    case "w": // arriba
        newPos[0]--
    case "s": // abajo
        newPos[0]++
    case "a": // izquierda
        newPos[1]--
    case "d": // derecha
        newPos[1]++
    default:
        fmt.Println("Movimiento no válido. Usa 'w', 'a', 's', 'd'.")
        return
    }

    if maze[newPos[0]][newPos[1]] != wall {
        playerPos = newPos
    } else {
        fmt.Println("¡Chocaste con una pared!")
    }

    if maze[playerPos[0]][playerPos[1]] == exit {
        fmt.Println("¡Felicidades! Has encontrado la salida.")
        os.Exit(0)
    }
}

func main() {
    fmt.Println("Bienvenido al juego de Laberinto.")
    fmt.Println("Usa 'w' para arriba, 's' para abajo, 'a' para izquierda, 'd' para derecha.")
    fmt.Println("Intenta encontrar la salida (E).")

    for {
        printMaze()
        var input string
        fmt.Print("Tu movimiento: ")
        fmt.Scan(&input)
        movePlayer(input)
    }
}
