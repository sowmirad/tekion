package dealerService

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestFetchOne(t *testing.T) {
	dealerDataSetup()
	Convey("Testing readOne func", t, func() {
		Convey("Testing for happy path flow", func() {
			var data dealer
			err := fetchOne(validContext, getDealerCollectionName(), bson.M{"_id": validDealerID}, nil, &data)
			Convey("Error returned by readOne should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validDealer)
			})
		})
		Convey("Testing for happy path flow with fields", func() {
			var data dealer
			err := fetchOne(validContext, getDealerCollectionName(), bson.M{"_id": validDealerID}, dealerFieldsSlice, &data)
			Convey("Error returned by readOne should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validDealerWithFields)
			})
		})
		Convey("Testing for invalid tenant", func() {
			var data interface{}
			err := fetchOne(invalidTenantContext, getDealerCollectionName(), bson.M{"_id": validDealerID}, nil, &data)
			Convey("Error returned by readOne should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "The tenant InvalidTenant is not connected")
			})
		})
		Convey("Testing for invalid dealer id", func() {
			var data interface{}
			err := fetchOne(validContext, getDealerCollectionName(), bson.M{"_id": invalidDealerID}, nil, &data)
			Convey("Error returned by readOne should not be nil", func() {
				So(err.Error(), ShouldEqual, "not found")
			})
		})
	})
	clearDealerDataSetup()
}

func TestFetchFixedOperations(t *testing.T) {
	fixedOperationDataSetup()
	Convey("Testing readFixedOperations func", t, func() {
		Convey("Testing for happy path flow", func() {
			var data []FixedOperation
			err := fetchFixedOperations(validContext, bson.M{"dealerID": validDealerID}, nil, &data)
			Convey("Error returned by readFixedOperations should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validFixedOperations)
			})
		})
		Convey("Testing for happy path flow with fields", func() {
			var data []FixedOperation
			err := fetchFixedOperations(validContext, bson.M{"dealerID": validDealerID}, fixedOperationFieldsSlice, &data)
			Convey("Error returned by readFixedOperations should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validFixedOperationsWithFields)
			})
		})
		Convey("Testing for invalid tenant", func() {
			var data []FixedOperation
			err := fetchFixedOperations(invalidTenantContext, bson.M{"dealerID": validDealerID}, nil, &data)
			Convey("Error returned by readFixedOperations should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "The tenant InvalidTenant is not connected")
			})
		})
		Convey("Testing for invalid dealer id", func() {
			var data []FixedOperation
			err := fetchFixedOperations(validContext, bson.M{"dealerID": invalidDealerID}, nil, &data)
			Convey("Error returned by readFixedOperations should not be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, data)
			})
		})
	})
	clearFixedOperationDataSetup()
}

func TestFetchDealerContacts(t *testing.T) {
	contactDataSetup()
	Convey("Testing readDealerContacts func", t, func() {
		Convey("Testing for happy path flow", func() {
			var data []dealerContact
			err := fetchDealerContacts(validContext, bson.M{"dealerID": validDealerID}, nil, &data)
			Convey("Error returned by readDealerContacts should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validContacts)
			})
		})
		Convey("Testing for happy path flow with fields", func() {
			var data []dealerContact
			err := fetchDealerContacts(validContext, bson.M{"dealerID": validDealerID}, contactFieldsSlice, &data)
			Convey("Error returned by readDealerContacts should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validContactsWithFields)
			})
		})
		Convey("Testing for invalid tenant", func() {
			var data []dealerContact
			err := fetchDealerContacts(invalidTenantContext, bson.M{"dealerID": validDealerID}, nil, &data)
			Convey("Error returned by readDealerContacts should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "The tenant InvalidTenant is not connected")
			})
		})
		Convey("Testing for invalid dealer id", func() {
			var data []dealerContact
			err := fetchDealerContacts(validContext, bson.M{"dealerID": invalidDealerID}, nil, &data)
			Convey("Error returned by readDealerContacts should not be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, data)
			})
		})
	})
	clearContactDataSetup()
}

func TestFetchDealerGoals(t *testing.T) {
	goalDataSetup()
	Convey("Testing readDealerGoals func", t, func() {
		Convey("Testing for happy path flow", func() {
			var data []dealerGoal
			err := fetchDealerGoals(validContext, bson.M{"dealerID": validDealerID}, nil, &data)
			Convey("Error returned by readDealerGoals should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validGoals)
			})
		})
		Convey("Testing for happy path flow with fields", func() {
			var data []dealerGoal
			err := fetchDealerGoals(validContext, bson.M{"dealerID": validDealerID}, goalFieldsSlice, &data)
			Convey("Error returned by readDealerGoals should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validGoalsWithFields)
			})
		})
		Convey("Testing for invalid tenant", func() {
			var data []dealerGoal
			err := fetchDealerGoals(invalidTenantContext, bson.M{"dealerID": validDealerID}, nil, &data)
			Convey("Error returned by readDealerGoals should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "The tenant InvalidTenant is not connected")
			})
		})
		Convey("Testing for invalid dealer id", func() {
			var data []dealerGoal
			err := fetchDealerGoals(validContext, bson.M{"dealerID": invalidDealerID}, nil, &data)
			Convey("Error returned by readDealerGoals should not be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, data)
			})
		})
	})
	clearGoalDataSetup()
}

func TestFetchDealerGroups(t *testing.T) {
	groupDataSetup()
	Convey("Testing readDealerGroups func", t, func() {
		Convey("Testing for happy path flow", func() {
			var data []dealerGroup
			err := fetchDealerGroups(validContext, bson.M{"dealers": validDealerID}, nil, &data)
			Convey("Error returned by readDealerGroups should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validGroups)
			})
		})
		Convey("Testing for happy path flow with fields", func() {
			var data []dealerGroup
			err := fetchDealerGroups(validContext, bson.M{"dealers": validDealerID}, groupFieldsSlice, &data)
			Convey("Error returned by readDealerGroups should be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, validGroupsWithFields)
			})
		})
		Convey("Testing for invalid tenant", func() {
			var data []dealerGroup
			err := fetchDealerGroups(invalidTenantContext, bson.M{"dealers": validDealerID}, nil, &data)
			Convey("Error returned by readDealerGroups should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "The tenant InvalidTenant is not connected")
			})
		})
		Convey("Testing for invalid dealer id", func() {
			var data []dealerGroup
			err := fetchDealerGroups(validContext, bson.M{"dealers": invalidDealerID}, nil, &data)
			Convey("Error returned by readDealerGroups should not be nil", func() {
				So(err, ShouldBeNil)
				So(data, ShouldResemble, data)
			})
		})
	})
	clearGroupDataSetup()
}
