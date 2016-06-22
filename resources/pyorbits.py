#!/usr/bin/env python3

import math


# Esta función calcula las aceleraciones en 'x' y 'y'
def aceleraciones(gamma ,M ,x,y):
    C  = -gamma*M
    r  = math.pow(x,2) + math.pow(y,2)
    ax = (C*x)/math.pow(r,1.5)
    ay = (C*y)/math.pow(r,1.5)
    
    return ax,ay


# Definí esta función pero no la usé jamás
def velocidades(gamma,M, vx, vy, x, y,dt):
    acc = aceleraciones(gamma,M,x,y)

    vxf = vx +dt*acc[0] 
    vyf = vy +dt*acc[1]

    return vxf,vyf


def posiciones(gamma,M, vx, vy, x, y,dt):

    acc = aceleraciones(gamma,M,x,y)

    vxf = vx +dt*acc[0] 
    vyf = vy +dt*acc[1]

    # Asi propuse la solución
    xf = x + dt*vx
    yf = y + dt*vy

    # En internet encontré esta sugerencia
    # que usa vxf y vxf que son las "velocidades futuras"
    # es decir que calcula la posición usando la velocidad futura
    #xf = x + dt*vxf
    #yf = y + dt*vyf

    # También encontré esta regla que saca el promedio
    # entre la velocidad actual y la velocidad futura
    # y así calcula las posiciones
    #xf = x + dt*(vx+vxf)/2.0
    #yf = y + dt*(vy+vyf)/2.0

    return xf,yf,vxf,vyf


x=0
y=5

vx=-0.7
vy=-1

print("#t\tx\ty\tvx\tvy")

dt=0.005
t=0.0
tf=2500


while (t<tf):
    print(t,x,y,vx,vy)
    # El resultado es de la forma
    #x,y,vx,vy = posiciones(0.5,9, vx, vy, x, y,dt)
    resultado = posiciones(0.5,9, vx, vy, x, y,dt)
    x = resultado[0]
    y = resultado[1]
    vx = resultado[2]
    vy = resultado[3]
    t+=dt

