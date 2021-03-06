package main

import "fmt"

type Road struct {
	Departure string
	Arrival   string
	Price     int
}

type Route struct {
	Roads []Road
	Price int
}

var best_route Route

func main() {
	fmt.Println(FindRoute("A", "B", []Road{
		{"A", "B", 3},
		{"B", "C", 4},
		{"A", "E", 2},
		{"E", "B", 5},
		{"D", "B", 5},
	}, Route{}))
}

func FindRoute(departure string, arrival string, roads []Road, route Route) Route {
	for i, road := range roads {
		if road.Departure == departure {
			if road.Arrival != arrival {
				FindRoute(road.Arrival, arrival, append(roads[:i], roads[i + 1:]...), Route{append(route.Roads, road), route.Price + road.Price})
			} else if route.Price += road.Price; best_route.Price == 0 || best_route.Price > route.Price {
				best_route = Route{append(route.Roads, road), route.Price}
			}
		}
	}
	return best_route
}
