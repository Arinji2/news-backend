package cronjobs

import "fmt"

func LiveNewsController() {

	fmt.Println("Running Live News Controller")
	news := getAllLiveNews()
	setLiveNews(news)
	fmt.Println("Finished Running Live News Controller")

}

func CountryNewsController() {

	fmt.Println("Running Country News Controller")
	news := getAllCountryNews()
	setCountryNews(news)
	fmt.Println("Finished Running Country News Controller")

}

func CategoryNewsController() {

	fmt.Println("Running Category News Controller")
	news := getAllCategoryNews()
	setCategoryNews(news)
	fmt.Println("Finished Running Category News Controller")

}
