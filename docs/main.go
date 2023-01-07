/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	startTime, err := time.Parse("20060102", "20200407")
	if err != nil {
		fmt.Println(err)
	}
	/*
		endTime, err := time.Parse("20060102", "20220309")
		if err != nil {
			fmt.Println(err)
		}

	*/
	// currentTime := time.(endTime-startTime)
	target := startTime.Add(time.Hour * 24 * 365 * 2)
	fmt.Printf("Time %v\n", target)

}
