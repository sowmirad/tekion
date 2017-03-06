package dealer
import (
	"bitbucket.org/tekion/tbaas/apiContext"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetDealerByIDMethod(t *testing.T) {
	validDealerID := "3"
	invalidDealerID := "2131242adasf"
	Convey("Create a context for right db i.e Buck", t, func() {
		Convey("Get loanerVehicle agreement for valid RoId", func() {
			ctx := apiContext.APIContext{Tenant: "Buck"}
			loanerVehicleresponse, err := GetDealerByID(ctx, validDealerID)

			Convey("Verify loanerVehicle response for empty loanerVehicleId", func() {
				So(loanerVehicleresponse, ShouldNotBeNil)

				Convey("error should not  be nil", func() {
					So(err, ShouldEqual,nil)
				})

			})
		})
		Convey("Get loanerVehicle agreement for invalid RoId", func() {
			ctx := apiContext.APIContext{Tenant: "Buck"}
			loanerAgreementResponse, err := GetDealerByID(ctx, invalidDealerID)

			Convey("Verify loanerVehicle response for loanerVehicleId", func() {
				So(loanerAgreementResponse, ShouldNotBeNil)

				Convey("error should not be nil", func() {
					So(err, ShouldNotBeNil)
				})

			})

		})
		Convey("Create a context for wrong tenant i.e Buck1", func() {
			ctx := apiContext.APIContext{Tenant: "Buck1"}
			Convey("Get error response ", func() {
				loanerAgreementResponse, err := GetDealerByID(ctx, validDealerID)
				Convey("error should not be empty", func() {
					So(err, ShouldNotBeEmpty)
					So(loanerAgreementResponse, ShouldNotBeNil)
				})

			})

		})
		Convey("Get loanerAgreement response if error in finding data for loanerVehicleApiPayload ", func() {
			object := "qewrwr2131243124"
			ctx := apiContext.APIContext{Tenant: "Buck"}
			loanerAgreementResponse, err := GetDealerByID(ctx, object)
			Convey("Verify the error  loanerVehicle  response", func() {
				Convey("error should not be nil", func() {
					So(err, ShouldNotBeNil)
					So(loanerAgreementResponse, ShouldNotBeNil)
				})

			})

		})

	})
}