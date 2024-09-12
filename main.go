package main

import (
	"MVC_PROJ/Controller2" //pentru aflarea rezultatelor
	"MVC_PROJ/View2"       //pentru citirea si afisarea datelor
)

func main() {
	// input
	rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx := View2.Read()

	// obtinerea rezultatelor
	v, gamma, Vmax, Gammamax := Controller2.DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax)

	// output
	View2.DispRes(nil, rho, S, P, G, Cx0, k, eta, Cz, Cx, v, gamma, Vmax, Gammamax)
}
