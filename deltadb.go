// (deltadb.go) - Contains the change of state commands of ./mfc
// Copyright (C) 2021 MaxflowO2, the only author of Max Flow Chain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"encoding/hex"
	"github.com/MaxflowO2/mfc/K12"
	"io/ioutil"
)

func hashDB() string {
	// get payload
	data, err := ioutil.ReadFile("blockchain.db")
	if err != nil {
		panic(err)
	}
	value := K12.Sum256(data)
	return hex.EncodeToString(value)
}
