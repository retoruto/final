package agency

func Optimized() {

	customer := GetCustomerDetails()
	destinations := GetRecommendedDestinations(customer)
	var Ins [10]Info

	quotes := [10]chan Quoting{}
	weathers := [10]chan Weather{}

	for i := range quotes {
		quotes[i] = make(chan Quoting)
	}

	for i := range weathers {
		weathers[i] = make(chan Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			quotes[i] <- GetQuote(d)
		}()

		go func() {
			weathers[i] <- GetWeatherForecast(d)
		}()
	}

	for index, dest := range destinations {
		Ins[index] = Info{Destinations:dest, Quote:<-quotes[index], Weather:<-weathers[index]}
	}
}
