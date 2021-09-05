package specification

import (
	"errors"
	"testing"

	"Sharykhin/rent-car/domain"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsCarEnginePowerCorrectSpecification(t *testing.T) {
	Convey("Given a zero power", t, func() {
		var power uint64 = 0
		Convey("When the power must be more than zero", func() {
			err := IsCarEnginePowerCorrectSpecification(power)
			Convey("Then an error is returned", func() {
				So(errors.Is(err, ErrPowerIsZero), ShouldBeTrue)
				Convey("And it must be considered as validation kind", func() {
					if err, ok := err.(*domain.Error); ok {
						So(err.Code, ShouldEqual, domain.ValidationErrorCode)
					} else {
						t.Errorf("failed to convert error into custom domain error: %v", err)
					}
				})
			})
		})
	})
	Convey("Given a non-zero power", t, func() {
		var power uint64 = 100
		Convey("When the power must be more than zero", func() {
			err := IsCarEnginePowerCorrectSpecification(power)
			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
