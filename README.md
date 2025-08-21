# 8 Puzzle en Go con Fyne

Este proyecto implementa el clásico juego 8 Puzzle con una interfaz gráfica en Go usando la librería Fyne. Permite jugar, desordenar, resolver automáticamente y paso a paso usando el algoritmo BFS (Breadth-First Search).

## Estructura del código principal

### Variables globales

- `tamano`: Tamaño de la cuadrícula (3x3).
- `estadoFinal`: Estado objetivo del puzzle.
- `estadoActual`: Estado actual de la cuadrícula.

### Funciones clave

#### copiarEstado

Copia el estado de la cuadrícula para evitar referencias compartidas.

#### reiniciarCuadricula

Reinicia la cuadrícula al estado final y actualiza la interfaz y contadores.

#### actualizarCuadricula

Actualiza la cuadrícula visual en la interfaz, mostrando los valores actuales.

#### desordenarCuadricula

Desordena la cuadrícula en n movimientos válidos y calcula el mínimo de pasos para resolver.

#### buscarVacio

Busca la posición del espacio vacío (0) en la cuadrícula.

#### movimientosValidos

Devuelve los movimientos válidos para el espacio vacío (arriba, abajo, izquierda, derecha).

#### resolverDesdeEstado

Resuelve el puzzle usando el algoritmo BFS (Breadth-First Search), garantizando la solución más corta en términos de número de movimientos.

#### serializar

Convierte el estado de la cuadrícula en una cadena para usar como clave en mapas durante la búsqueda BFS.

#### puedeMover

Verifica si una ficha se puede mover (si está adyacente al espacio vacío).

#### abs

Devuelve el valor absoluto de un número.

#### estaResuelto

Verifica si el puzzle está resuelto (estado actual igual al estado final).

#### actualizarEtiquetas

Actualiza las etiquetas de pasos y mínimos en la interfaz, cambiando el color si se excede el mínimo.

### Interfaz gráfica (main)

- Inicializa la aplicación y la ventana.
- Crea los botones de la cuadrícula y los envuelve en contenedores para permitir resaltar el movimiento.
- Utiliza el algoritmo BFS para encontrar la solución óptima.
- Incluye botones para reiniciar, desordenar, resolver automáticamente y paso a paso.
- Organiza los componentes con márgenes y separaciones para una mejor visualización.

## Ejecución

1. Instala la librería Fyne:
   ```bash
   go get fyne.io/fyne/v2
   ```
2. Ejecuta el programa:
   ```bash
   go run main.go
   ```

## Créditos

Desarrollado por Diego, 2025.

## Licencia

Este proyecto está licenciado bajo la MIT License - ver el archivo [LICENSE](LICENSE) para más detalles.
