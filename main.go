package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const tamano = 3

var (
	finalState = [tamano][tamano]int{ {1, 2, 3}, {4, 5, 6}, {7, 8, 0} }
	currentBoard [tamano][tamano]int
)

func copyState(src [tamano][tamano]int) [tamano][tamano]int {
	var nuevoBoard [tamano][tamano]int

	for fila := 0; fila < tamano; fila++ {
		for columna := 0; columna < tamano; columna++ {
			nuevoBoard[fila][columna] = src[fila][columna]
		}
	}
	return nuevoBoard
}

func resetearBoard(updateFunc func(), botones [][]*widget.Button, contadorPasos *canvas.Text, labelMinimo *canvas.Text, pasos *int, minPasos *int, mensajeLabel *canvas.Text) {

	currentBoard = copyState(finalState)
	*pasos = 0
	*minPasos = 0
	updateFunc()
	actualizarLabels(contadorPasos, labelMinimo, *pasos, *minPasos)
	mensajeLabel.Text = ""
	mensajeLabel.Refresh()
}

func actualizarGrid(botones [][]*widget.Button) {
	
	for i := 0; i < tamano; i++ {
		for j := 0; j < tamano; j++ {
			numero := currentBoard[i][j]
			if numero == 0 {
				botones[i][j].SetText("")
				botones[i][j].Disable()
			} else {
				// conversion a string para mostrar en boton
				texto := fmt.Sprintf("%d", numero)
				botones[i][j].SetText(texto)
				botones[i][j].Enable()
			}
		}
	}
}

func shuffleBoard(updateFunc func(), botones [][]*widget.Button, n int, contadorPasos *canvas.Text, labelMinimo *canvas.Text, pasos *int, minPasos *int, mensajeLabel *canvas.Text) {
	currentBoard = copyState(finalState)
	
	for contador := 0; contador < n; contador++ {
		emptyRow, emptyCol := findEmptySpace(currentBoard)
		posibleMoves := getValidMoves(emptyRow, emptyCol)
		
		randomIndex := rand.Intn(len(posibleMoves))
		selectedMove := posibleMoves[randomIndex]
		
		temp := currentBoard[emptyRow][emptyCol]
		currentBoard[emptyRow][emptyCol] = currentBoard[selectedMove[0]][selectedMove[1]]
		currentBoard[selectedMove[0]][selectedMove[1]] = temp
	}
	*pasos = 0
	
	solucionCompleta := solve(currentBoard)
	*minPasos = len(solucionCompleta) - 1
	updateFunc()
	actualizarLabels(contadorPasos, labelMinimo, *pasos, *minPasos)
	mensajeLabel.Text = ""
	mensajeLabel.Refresh()
}

func findEmptySpace(board [tamano][tamano]int) (int, int) {
	
	for fila := 0; fila < tamano; fila++ {
		for columna := 0; columna < tamano; columna++ {
			if board[fila][columna] == 0 {
				return fila, columna
			}
		}
	}
	return -1, -1 // valor por defecto si no se encuentra
}

func getValidMoves(fila, columna int) [][2]int {
	movimientos := [][2]int{}
	// verificar arriba
	if fila > 0 {
		movimientos = append(movimientos, [2]int{fila - 1, columna})
	}
	// verificar abajo
	if fila < tamano-1 {
		movimientos = append(movimientos, [2]int{fila + 1, columna})
	}
	// verificar izquierda
	if columna > 0 {
		movimientos = append(movimientos, [2]int{fila, columna - 1})
	}
	// verificar derecha
	if columna < tamano-1 {
		movimientos = append(movimientos, [2]int{fila, columna + 1})
	}
	return movimientos
}

func getPuzzleSolution() [][tamano][tamano]int {
	return solve(currentBoard)
}


func solve(startBoard [tamano][tamano]int) [][tamano][tamano]int {
	// lista de todos los estados visitados
	var todosLosEstados [][tamano][tamano]int
	var padresDeEstados []int
	
	// agregar estado inicial
	todosLosEstados = append(todosLosEstados, startBoard)
	padresDeEstados = append(padresDeEstados, -1)
	
	// mapa para verificar estados visitados
	visitados := make(map[string]bool)
	visitados[boardToString(startBoard)] = true
	
	indiceSolucion := -1
	
	// busqueda iterativa
	for i := 0; i < len(todosLosEstados); i++ {
		estadoActual := todosLosEstados[i]
		
		// verificar si es el estado final
		if isSolved(estadoActual) {
			indiceSolucion = i
			break
		}
		
		// encontrar espacio vacio
		emptyFila, emptyCol := findEmptySpace(estadoActual)
		
		// generar movimientos posibles
		movesPosibles := getValidMoves(emptyFila, emptyCol)
		
		// procesar cada movimiento
		for _, move := range movesPosibles {
			// crear nueva configuracion del board
			nuevoBoard := copyState(estadoActual)
			
			// intercambiar posiciones
			nuevoBoard[emptyFila][emptyCol] = nuevoBoard[move[0]][move[1]]
			nuevoBoard[move[0]][move[1]] = 0
			
			// verificar si ya fue visitado
			boardString := boardToString(nuevoBoard)
			if !visitados[boardString] {
				visitados[boardString] = true
				todosLosEstados = append(todosLosEstados, nuevoBoard)
				padresDeEstados = append(padresDeEstados, i)
			}
		}
	}
	
	// retornar lista vacia si no hay solucion
	if indiceSolucion == -1 {
		return [][tamano][tamano]int{}
	}
	
	// reconstruir secuencia de solucion
	var pathSolucion [][tamano][tamano]int
	indiceActual := indiceSolucion
	
	for indiceActual != -1 {
		pathSolucion = append([][tamano][tamano]int{todosLosEstados[indiceActual]}, pathSolucion...)
		indiceActual = padresDeEstados[indiceActual]
	}
	
	return pathSolucion
}

// conversion de board a string para comparaciones
func boardToString(board [tamano][tamano]int) string {
	resultado := ""
	for i := 0; i < tamano; i++ {
		for j := 0; j < tamano; j++ {
			if board[i][j] < 10 {
				resultado += fmt.Sprintf("%d", board[i][j])
			}
		}
	}
	return resultado
}

func canMoveHere(fila, col int) bool {
	emptyFila, emptyCol := findEmptySpace(currentBoard)
	// verificar si esta adyacente al espacio vacio
	if fila == emptyFila && (col == emptyCol-1 || col == emptyCol+1) {
		return true
	}
	if col == emptyCol && (fila == emptyFila-1 || fila == emptyFila+1) {
		return true
	}
	return false
}

// verificar si un board especifico esta resuelto
func isSolved(board [tamano][tamano]int) bool {
	for i := 0; i < tamano; i++ {
		for j := 0; j < tamano; j++ {
			if board[i][j] != finalState[i][j] {
				return false
			}
		}
	}
	return true
}

// verificar si el board actual esta resuelto
func isCurrentBoardSolved() bool {
	// comparacion manual
	for i := 0; i < tamano; i++ {
		for j := 0; j < tamano; j++ {
			if currentBoard[i][j] != finalState[i][j] {
				return false
			}
		}
	}
	return true
}

func actualizarLabels(contadorPasos, labelMinimo *canvas.Text, pasos, minPasos int) {
	contadorPasos.Text = fmt.Sprintf("Pasos: %d", pasos)
	// cambiar color si se excede el minimo
	if pasos > minPasos && minPasos > 0 {
		contadorPasos.Color = theme.ErrorColor()
	} else {
		contadorPasos.Color = theme.ForegroundColor()
	}
	contadorPasos.Refresh()
	labelMinimo.Text = fmt.Sprintf("Mínimo para resolver: %d", minPasos)
	labelMinimo.Refresh()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	app := app.New()
	ventana := app.NewWindow("8 Puzzle - Mi App Casera")

	var solutionPath [][tamano][tamano]int
	var stepIndex int
	contadorPasos := 0
	minimoPasos := 0

	// crear labels
	labelPasos := canvas.NewText("Pasos: 0", theme.ForegroundColor())
	labelPasos.TextSize = 18
	labelMinimo := canvas.NewText("Mínimo para resolver: 0", theme.ForegroundColor())
	labelMinimo.TextSize = 18
	mensajeLabel := canvas.NewText("", theme.SuccessColor())
	mensajeLabel.TextSize = 20

	// crear botones para el grid
	botones := make([][]*widget.Button, tamano)
	contenedores := make([][]*fyne.Container, tamano)
	for i := 0; i < tamano; i++ {
		botones[i] = make([]*widget.Button, tamano)
		contenedores[i] = make([]*fyne.Container, tamano)
		for j := 0; j < tamano; j++ {
			// capturar indices para closure
			filaCapturada, colCapturada := i, j
			rectangulo := canvas.NewRectangle(theme.BackgroundColor())
			boton := widget.NewButton("", nil)
			boton.Resize(fyne.NewSize(70, 70))
			boton.OnTapped = func() {
				if canMoveHere(filaCapturada, colCapturada) {
					emptyRow, emptyCol := findEmptySpace(currentBoard)
					// intercambio manual
					temp := currentBoard[emptyRow][emptyCol]
					currentBoard[emptyRow][emptyCol] = currentBoard[filaCapturada][colCapturada]
					currentBoard[filaCapturada][colCapturada] = temp
					contadorPasos++
					actualizarGrid(botones)
					actualizarLabels(labelPasos, labelMinimo, contadorPasos, minimoPasos)
					fmt.Printf("Pasos: %d, Mínimo: %d\n", contadorPasos, minimoPasos)
					if isCurrentBoardSolved() {
						mensajeLabel.Text = "¡Ganaste! Puzzle completado."
						mensajeLabel.Color = theme.SuccessColor()
						mensajeLabel.Refresh()
						fmt.Println("¡Puzzle resuelto!")
					}
					// iluminar boton cuando se mueve
					rectangulo.FillColor = theme.SuccessColor()
					rectangulo.Refresh()
				}
			}
			botones[i][j] = boton
			contenedores[i][j] = container.NewMax(rectangulo, boton)
		}
	}

	// funcion para actualizar grid
	funcionUpdate := func() { actualizarGrid(botones) }
	resetearBoard(funcionUpdate, botones, labelPasos, labelMinimo, &contadorPasos, &minimoPasos, mensajeLabel)

	// crear grid principal
	gridPrincipal := container.NewGridWithColumns(tamano)
	for i := 0; i < tamano; i++ {
		for j := 0; j < tamano; j++ {
			gridPrincipal.Add(contenedores[i][j])
		}
	}

	// botones de control
	botonReset := widget.NewButton("Reset", func() {
		resetearBoard(funcionUpdate, botones, labelPasos, labelMinimo, &contadorPasos, &minimoPasos, mensajeLabel)
		stepIndex = 0
		solutionPath = [][tamano][tamano]int{}
		fmt.Println("Puzzle reseteado.")
	})

	botonShuffle := widget.NewButton("Mezclar", func() {
		shuffleBoard(funcionUpdate, botones, 30, labelPasos, labelMinimo, &contadorPasos, &minimoPasos, mensajeLabel)
		stepIndex = 0
		solutionPath = [][tamano][tamano]int{}
		fmt.Printf("Puzzle mezclado. Mínimo para resolver: %d\n", minimoPasos)
	})

	botonSolve := widget.NewButton("Resolver Todo", func() {
		solucion := getPuzzleSolution()
		if len(solucion) > 0 {
			solutionPath = solucion
			// copiar estado final
			currentBoard = copyState(solutionPath[len(solutionPath)-1])
			actualizarGrid(botones)
			contadorPasos = len(solutionPath) - 1
			actualizarLabels(labelPasos, labelMinimo, contadorPasos, minimoPasos)
			stepIndex = 0
			mensajeLabel.Text = "¡Ganaste! Puzzle completado automáticamente."
			mensajeLabel.Color = theme.SuccessColor()
			mensajeLabel.Refresh()
			fmt.Println("¡Puzzle resuelto de manera automática!")
		}
	})

	botonStepByStep := widget.NewButton("Paso a Paso", func() {
		if len(solutionPath) == 0 {
			solucionCalculada := getPuzzleSolution()
			if len(solucionCalculada) > 0 {
				solutionPath = solucionCalculada
				stepIndex = 0
			}
		}
		if len(solutionPath) > 0 && stepIndex < len(solutionPath) {
			currentBoard = copyState(solutionPath[stepIndex])
			actualizarGrid(botones)
			contadorPasos = stepIndex
			actualizarLabels(labelPasos, labelMinimo, contadorPasos, minimoPasos)
			stepIndex++
			if isCurrentBoardSolved() {
				mensajeLabel.Text = "¡Ganaste! Puzzle completado paso a paso."
				mensajeLabel.Color = theme.SuccessColor()
				mensajeLabel.Refresh()
				fmt.Println("¡Puzzle resuelto paso a paso!")
			}
		}
	})

	// organizar controles
	controlesContainer := container.NewVBox(
		widget.NewLabel(""),
		botonReset,
		botonShuffle,
		botonSolve,
		botonStepByStep,
	)

	// layout principal con espaciadores
	contenedorPrincipal := container.NewHBox(
		container.NewVBox(
			container.NewVBox(
				widget.NewLabelWithStyle("8 Puzzle", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
				labelPasos,
				labelMinimo,
				mensajeLabel,
			),
			container.NewVBox(gridPrincipal),
		),
		widget.NewLabel("     "), // espaciador horizontal
		container.NewVBox(controlesContainer),
	)

	// aplicar margenes
	contenedorConMargenes := container.NewVBox(
		widget.NewLabel(""), // margen superior
		container.NewHBox(
			widget.NewLabel("   "), // margen izquierdo
			contenedorPrincipal,
			widget.NewLabel("   "), // margen derecho
		),
		widget.NewLabel(""), // margen inferior
	)

	ventana.SetContent(contenedorConMargenes)
	ventana.Resize(fyne.NewSize(380, 340)) // tamaño fijo
	ventana.CenterOnScreen()
	ventana.ShowAndRun()
}