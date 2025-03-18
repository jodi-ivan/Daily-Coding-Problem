package main

import (
	"log"
)

func isPrime(n int) bool {
	/*

		bool isPrime(int n) {
		    // Corner cases
		    if(n <= 1) return false;
		    if(n <= 3) return true;

		    // This is checked so that we can skip
		    // middle five numbers in below loop
		    if(n % 2 == 0 || n % 3 == 0) return false;

		    for(int i = 5; i * i <= n; i = i + 6)
		        if(n % i == 0 || n % (i + 2) == 0) return false;

		    return true;
		}

	*/

	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
func main() {

	n := 150
	count := 0
	// 1, 3, 5, 7, 9, 11, 13,
	for i := 5; i <= n; i += 3 {
		// if n%i == 0 || n%(i+2) == 0 {
		// 	continue
		// }

		if isPrime(i) && isPrime(n-i) {
			log.Println(i, isPrime(i))
			log.Println(n-i, isPrime(n-i))
			log.Println("=========================================")
			count++
		}

	}
	log.Println(isPrime(13))
}
