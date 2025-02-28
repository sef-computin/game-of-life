# Conway's Game of Life

This is an implementation of **Conway's Game of Life**, written in **Go** and utilizing **OpenGL** for rendering the grid and cell updates.

![Alt text]("game of life.gif") / ![]("game of life.gif")

## Overview

Conway's Game of Life is a cellular automaton devised by mathematician **John Conway** in 1970. The game consists of a grid of cells that evolve over time based on a few simple rules. Cells can be alive or dead, and their state is determined by the number of live neighbors they have.

### Rules of Conway's Game of Life:

1.  **Any live cell with fewer than two live neighbors dies (underpopulation).**
2.  **Any live cell with two or three live neighbors continues to live (survival).**
3.  **Any live cell with more than three live neighbors dies (overpopulation).**
4.  **Any dead cell with exactly three live neighbors becomes a live cell (reproduction).**

## Installation

1.  **Clone the Repository**:
    
    ```bash
    git clone https://github.com/sef-comp/game-of-life.git
    cd game-of-life
    ```
    
2.  **Install Dependencies**:
    
    Make sure you have Go and the required OpenGL libraries installed. If you haven't already installed OpenGL bindings for Go, run:
    
    ```bash
    go get -u github.com/go-gl/gl/v4.6-compatibility/gl
    go get -u github.com/vbsw/glut
    ```
    
3.  **Build the Project**:
    
    ```bash
    go build
    
    ```
    
4.  **Run the Game**:
    
    After building the project, you can run the game with:
    
    ```bash
    ./game-of-life
    
    ```
## Controls

-   **Mouse**: Left-click to toggle cell living state.
-   **Enter**: Pause/Unpause
-   **Esc**: Exit the game.
