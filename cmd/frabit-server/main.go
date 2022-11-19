/*

  (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
  GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

  This file is part of Frabit

*/

package main

import (
	"fmt"
	"time"

	_ "github.com/frabit-io/frabit/server"
)

func main() {
	fmt.Printf("Datatime:%v\n", time.Now().Format(time.RFC3339))
}
