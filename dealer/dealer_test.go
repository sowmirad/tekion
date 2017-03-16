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
		Convey("Get dealer info for valid dealerId", func() {
			ctx := apiContext.APIContext{Tenant: "Buck"}
			loanerVehicleresponse, err := GetDealerByID(ctx, validDealerID)

			Convey("Verify existing dealer (before new fields added) for the data", func() {
				So(loanerVehicleresponse, ShouldNotBeNil)
				So(loanerVehicleresponse.SkillSet, ShouldBeEmpty)
				So(loanerVehicleresponse.ServiceGroup, ShouldBeEmpty)
				Convey("error should not  be nil", func() {
					So(err, ShouldEqual, nil)
				})

			})
		})
		Convey("Get dealer info for invalid dealerId", func() {
			ctx := apiContext.APIContext{Tenant: "Buck"}
			loanerAgreementResponse, err := GetDealerByID(ctx, invalidDealerID)

			Convey("Verify dealer response for  invalidDealerID", func() {
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
		Convey("Get dealer response if error in finding data for dealerId ", func() {
			object := "qewrwr2131243124"
			ctx := apiContext.APIContext{Tenant: "Buck"}
			loanerAgreementResponse, err := GetDealerByID(ctx, object)
			Convey("Verify the error for dealerId  response", func() {
				Convey("error should not be nil", func() {
					So(err, ShouldNotBeNil)
					So(loanerAgreementResponse, ShouldNotBeNil)
				})

			})

		})

	})
}

func TestInsertMethod(t *testing.T) {
	dealerInputeObject := Dealer{DealerName:"Seaside Infiniti", SkillSet:[]string{"Engine"},ServiceGroup:[]string{""},}
	Convey("Create a context for right db i.e Buck", t, func() {
		Convey("Get dealer info for valid context", func() {
			ctx := apiContext.APIContext{Tenant: "Buck"}
			err := dealerInputeObject.Insert(ctx)

			Convey("Verify dealer response for  valid context", func() {
				Convey("error should not  be nil", func() {
					So(err, ShouldNotBeEmpty)
				})

			})
		})

		Convey("Create a context for wrong tenant i.e Buck1", func() {
			ctx := apiContext.APIContext{Tenant: "Buck1"}
			Convey("Get error response ", func() {
				err := dealerInputeObject.Insert(ctx)
				Convey("error should not be empty", func() {
					So(err, ShouldNotBeEmpty)
				})

			})

		})
		Convey("Create a context for wrong dealerId i.e 123asdas", func() {
			ctx := apiContext.APIContext{DealerID: "123asdas"}
			Convey("Get error response ", func() {
				err := dealerInputeObject.Insert(ctx)
				Convey("error should not be empty", func() {
					So(err, ShouldNotBeEmpty)
				})

			})

		})

	})
}
