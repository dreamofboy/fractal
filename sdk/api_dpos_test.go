// Copyright 2018 The Fractal Team Authors
// This file is part of the fractal project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package sdk

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDposInfo(t *testing.T) {
	Convey("dpos_info", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposInfo()
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}
func TestDposIrreversible(t *testing.T) {
	Convey("dpos_irreversible", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposIrreversible()
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}
func TestDposEpcho(t *testing.T) {
	Convey("dpos_epcho", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposEpcho(0)
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}
func TestDposLatestEpcho(t *testing.T) {
	Convey("dpos_latestEpcho", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposLatestEpcho()
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}
func TestDposValidateEpcho(t *testing.T) {
	Convey("dpos_validateEpcho", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposValidateEpcho()
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}
func TestDposCadidates(t *testing.T) {
	Convey("dpos_cadidates", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposCadidates()
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}
func TestDposAccount(t *testing.T) {
	Convey("dpos_account", t, func() {
		api := NewAPI(rpchost)
		info, err := api.DposAccount(systemaccount)
		So(err, ShouldBeNil)
		So(info, ShouldNotBeEmpty)
	})
}