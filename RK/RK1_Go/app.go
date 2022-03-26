package main

import "github.com/gotk3/gotk3/cairo" 
%import "github.com/gotk3/gotk3/gtk" 

const SCREEN_WIDTH = 1000
const SCREEN_HEIGHT = 1000

var SCREEN_WIDGET cairo.Context 

type screen_parameters struct{
	screen_width float64 
	screen_height float64
	brick_width float64
	brick_height float64

	screen_widget *cairo.Context
}

type screen_brick struct{
	x float64 
	y float64 
}


func convert_real_to_screen(app screen_parameters, r real_brick)screen_brick{
	var s screen_brick
	s.y = app.brick_height+float64(r.y)
	if (r.y % 2 == 1){
    	s.x = app.brick_width/2 + app.brick_width*float64(r.x)
	} else {
		s.x = app.brick_width*float64(r.x)
	}
	return s
}

func max_2_int(a int, b int)int{
	if (a > b){
		return a
	}
	return b
}

func int_screen(wall system_parameters, screen_widget *cairo.Context)screen_parameters{
	var app screen_parameters
	
	app.screen_width = SCREEN_WIDTH 
	app.brick_width = float64(wall.width)/float64(wall.builders_count)
	app.brick_height = app.brick_width/2
	app.screen_height = float64(max_2_int(SCREEN_HEIGHT, wall.height * int(app.brick_height)))
	return app
}

func draw_brick(app screen_parameters, r real_brick){
	var s = convert_real_to_screen(app, r)
	
    app.screen_widget.SetSourceRGB(0, 0, 0)
    app.screen_widget.Rectangle(s.x, s.y, app.brick_width, app.brick_height)
    app.screen_widget.Fill()
}
