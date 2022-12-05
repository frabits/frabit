/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package main

import "fmt"

type Agent struct {
	frabit string
}

func New(f string) *Agent {
	return &Agent{
		frabit: f,
	}
}

func RunAgent(mysqlURI string) {
	fmt.Println("start  agent process as sidecar")
}
