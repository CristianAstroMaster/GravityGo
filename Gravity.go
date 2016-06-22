package main

import (
	"fmt"
	"math"
)

func aceleraciones(gamma ,M ,x,y float64)(float64,float64){
	var C float64 = -gamma*M
	var r float64 = math.Pow(x,2) + math.Pow(y,2)
	var ax float64 = (C*x)/math.Pow(r,1.5)
	var ay float64 = (C*y)/math.Pow(r,1.5)

	return ax,ay
}

func velocidades(gamma,M, vx, vy, x, y,dt float64)(float64,float64){
	var ax,ay float64
	var vxf,vyf float64

	ax,ay = aceleraciones(gamma,M,x,y)

	vxf = vx +dt*ax 
	vyf = vy +dt*ay

	return vxf,vyf
}

func posiciones(gamma,M, vx, vy, x, y,dt float64)(float64,float64,float64,float64){
	var ax,ay float64
	var vxf,vyf float64
	var xf,yf float64

	ax,ay = aceleraciones(gamma,M,x,y)

	vxf = vx +dt*ax 
	vyf = vy +dt*ay

	// Proposed but seemingly wrong
	xf = x + dt*vx
	yf = y + dt*vy

	// An URL suggest this
	//xf = x + dt*vxf
	//yf = y + dt*vyf

	// Midpoint rule
	//xf = x + dt*(vx+vxf)/2.0
	//yf = y + dt*(vy+vyf)/2.0

	return xf,yf,vxf,vyf
}

func main() {
	var x,y float64
	var vx,vy float64
	var t,dt,tf float64

	x=0
	//y=15
	y=5

	vx=-0.7
	vy=-1

	fmt.Println("#t\tx\ty\tvx\tvy")

	dt=0.005
	t=0.0
	tf=2500


	for t<tf{
		fmt.Println(t,x,y,vx,vy)
		x,y,vx,vy = posiciones(0.5,9, vx, vy, x, y,dt)
		t+=dt
	}



}
