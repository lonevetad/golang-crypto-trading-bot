// Copyright © 2017 Alessandro Sanino <saninoale@gmail.com>
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
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package strategies

import (
	"time"

	"github.com/saniales/golang-crypto-trading-bot/environment"
	"github.com/saniales/golang-crypto-trading-bot/exchanges"
)

// IntervalStrategy is an interval based strategy.
type IntervalStrategy struct {
	Model    StrategyModel
	Interval time.Duration
}

// Name returns the name of the strategy.
func (is IntervalStrategy) Name() string {
	return is.Model.Name
}

// String returns a string representation of the object.
func (is IntervalStrategy) String() string {
	return is.Name()
}

// Apply executes Cyclically the On Update, basing on provided interval.
func (is IntervalStrategy) Apply(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) {
	var err error
	var delta time.Duration
	var timeBefore time.Time
	if is.Model.Setup != nil {
		err = is.Model.Setup(wrappers, markets)
		if err != nil && is.Model.OnError != nil {
			is.Model.OnError(err)
		}
	}
	for err == nil {
		howMuchToSleep = time.Now()
		err = is.Model.OnUpdate(wrappers, markets)
		if err != nil && is.Model.OnError != nil {
			is.Model.OnError(err)
		}
		howMuchToSleep = time.Now().Sub(timeBefore) - is.Interval //it's in time.Duration

		if(howMuchToSleep > 0){
			time.Sleep(howMuchToSleep)
		}
	}
	if is.Model.TearDown != nil {
		err = is.Model.TearDown(wrappers, markets)
		if err != nil && is.Model.OnError != nil {
			is.Model.OnError(err)
		}
	}
}
