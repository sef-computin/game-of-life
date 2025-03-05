package main

import (
	"os"
	"runtime"

	"github.com/go-gl/gl/v4.5-compatibility/gl"
	"github.com/vbsw/glut"
)

const (
	WINWIDTH   = 1500
	WINHEIGHT  = 1500
	FRAMESPEED = 15

	GRIDY = 100
	GRIDX = 100
	GRIDS = min(WINWIDTH/GRIDX, WINHEIGHT/GRIDY)
)

const (
	DEAD = iota
	ALIVE
)

type cell struct {
	x int
	y int
}

var alive_cells map[cell]interface{} = make(map[cell]interface{})

var grid []int

var run_state bool

var framecount uint32 = 0

func main() {
	runtime.LockOSThread()

	grid = make([]int, GRIDX*GRIDY)
	run_state = false
	// alive_cells = append(alive_cells, []cell{{50, 49}, {50, 50}, {51, 50}, {50, 51}, {49, 51}}...)

	_ = initOpenGL()
	init_w()
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func init_w() {
	glut.Init()
	glut.InitDisplayMode(glut.DOUBLE | glut.RGBA)
	glut.InitWindowSize(WINWIDTH, WINHEIGHT)
	glut.CreateWindow("Conway's Game of Life")

	gl.ClearColor(0.3, 0.33, 0.33, 0)
	gl.Ortho(0, WINWIDTH, WINHEIGHT, 0, -1, 1)
	glut.TimerFunc(1000/FRAMESPEED, timer, 0)

	glut.DisplayFunc(display)
	// glut.
	glut.KeyboardUpFunc(keyboardUp)
	glut.MouseFunc(mouseClicked)
	glut.MainLoop()
}

func mouseClicked(btn, state, x, y int) {
	if state == 1 {
		return
	}
	xc, yc := x/GRIDS, y/GRIDS
	c := cell{xc, yc}
	if _, ok := alive_cells[c]; ok {
		delete(alive_cells, c)
	} else {
		alive_cells[c] = true
	}
}

func keyboardUp(key uint8, x, y int) {
	if key == 13 {
		run_state = !run_state
	}
	if key == 27 {
		glut.DestroyWindow(glut.GetWindow())
		os.Exit(0)
	}
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	repopulate_grid()
	draw_grid()

	if run_state {
		next_generation()
	} else {
		draw_pause()
	}

	glut.SwapBuffers()
}

func repopulate_grid() {
	clear_grid()

	for cell := range alive_cells {
		grid[cell.y*GRIDX+cell.x] = 1
	}
}

func clear_grid() {
	for i := 0; i < len(grid); i++ {
		grid[i] = 0
	}
}

func next_generation() {
	check_area := map[cell]int{}

	for cl := range alive_cells {
		xleft, xright := cl.x-1, cl.x+1
		yleft, yright := cl.y-1, cl.y+1

		for i := xleft; i <= xright; i++ {
			for j := yleft; j <= yright; j++ {
				if !(i == cl.x && j == cl.y) {
					x, y := i, j
					if x < 0 {
						x = GRIDX - 1
					} else if x > GRIDX-1 {
						x = 0
					}
					if y < 0 {
						y = GRIDY - 1
					} else if y > GRIDY-1 {
						y = 0
					}

					check_area[cell{x, y}]++
				}
			}
		}
	}

	for k := range alive_cells {
		delete(alive_cells, k)
	}

	for cl, n := range check_area {
		if grid[cl.y*GRIDX+cl.x] == 0 {
			if n == 3 {
				alive_cells[cl] = true
			}
		} else if n >= 2 && n <= 3 {
			alive_cells[cl] = true
		}
	}

	// alive_cells = next_gen
}

func draw_grid() {

	for y := 0; y < GRIDY; y++ {
		for x := 0; x < GRIDX; x++ {
			if grid[y*GRIDX+x] > 0 {
				gl.Color3f(1, 1, 1)
			} else {
				gl.Color3f(0, 0, 0)
			}
			var x0, y0 int32 = int32(x * GRIDS), int32(y * GRIDS)
			gl.Begin(gl.QUADS)
			gl.Vertex2i(x0+1, y0+1)
			gl.Vertex2i(x0+1, y0+int32(GRIDS)-1)
			gl.Vertex2i(x0+int32(GRIDS)-1, y0+int32(GRIDS)-1)
			gl.Vertex2i(x0+int32(GRIDS)-1, y0+1)
			gl.End()
		}
	}

}

func draw_pause() {
	pause_msg := "PAUSE"
	gl.Color3f(1.0, 0.2, 0.2)

	gl.RasterPos2i(32, 32)
	for _, c := range pause_msg {
		glut.BitmapCharacter(glut.BITMAP_TIMES_ROMAN_24, c)
	}

}

func timer(i int) {
	glut.PostRedisplay()
	framecount++
	glut.TimerFunc(1000/FRAMESPEED, timer, 0)
}
