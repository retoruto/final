package agency

func Naive() {

	var Ins [10]Info
	customers := GetCustomerDetails()
	destinations := GetRecommendedDestinations(customers)

	for index, dest := range destinations {
		q := GetQuote(dest)
		w := GetWeatherForecast(dest)
		Ins[index] = Info{Destinations:dest, Quote:q, Weather:w}
	}

}
