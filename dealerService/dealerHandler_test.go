package dealerService

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReadDealer(t *testing.T) {
	dealerDataSetup()
	Convey("Testing ReadDealer", t, func() {
		Convey("ReadDealer should return dealer object", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/dealer", nil)
			if err != nil {
				t.Error(err)
			}
			setHeadersAndContext(req)
			readDealer(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := dealer{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validDealer)
			})
		})
		Convey("ReadDealer return dealer object with specified fields", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/dealer?fields="+dealerFields, nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealer(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := dealer{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validDealerWithFields)
			})
		})
		Convey("ReadDealer should return status code 204, invalid dealer id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/dealer", nil)
			setHeadersAndInvalidDealerIDContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealer(res, req)
			Convey("Response code should be 204", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusNoContent)
				So(apiRes.Meta.Msg, ShouldEqual, "No document found")
			})
		})
		Convey("ReadDealer should return status code 400, invalid tenant name", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/dealer", nil)
			setHeadersAndInvalidTenantContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealer(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
	})
	clearDealerDataSetup()
}

func TestReadFixedOperations(t *testing.T) {
	fixedOperationDataSetup()
	Convey("Testing ReadFixedOperations", t, func() {
		Convey("ReadFixedOperations should return fixed operation object", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperations", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperations(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []fixedOperation{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validFixedOperations)
			})
		})
		Convey("ReadFixedOperations should return fixed operation object with specified fields", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperation?fields="+fixedOperationFields, nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperations(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []fixedOperation{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validFixedOperationsWithFields)
			})
		})
		Convey("ReadFixedOperations should return 'No document found', invalid dealer id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperations", nil)
			setHeadersAndInvalidDealerIDContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperations(res, req)
			Convey("Response code should be 204", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusNoContent)
				So(apiRes.Meta.Msg, ShouldEqual, "No document found")
			})
		})
		Convey("ReadFixedOperations should return 400, invalid tenant name", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperations", nil)
			setHeadersAndInvalidTenantContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperations(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
	})
	clearFixedOperationDataSetup()
}

func TestReadContacts(t *testing.T) {
	contactDataSetup()
	Convey("Testing ReadContacts", t, func() {
		Convey("ReadContacts should return dealer contact object", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/contacts", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerContacts(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []dealerContact{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validContacts)
			})
		})
		Convey("ReadContacts should return dealer contact object with specified fields", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/contacts?fields="+contactFields, nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerContacts(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []dealerContact{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validContactsWithFields)
			})
		})
		Convey("ReadContacts should return 'No document found', invalid dealer id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/contacts", nil)
			setHeadersAndInvalidDealerIDContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerContacts(res, req)
			Convey("response code should be 204", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusNoContent)
				So(apiRes.Meta.Msg, ShouldEqual, "No document found")
			})
		})
		Convey("ReadContacts should return 400, invalid tenant name", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/contacts", nil)
			setHeadersAndInvalidTenantContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerContacts(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
	})
	clearContactDataSetup()
}

func TestReadGoals(t *testing.T) {
	goalDataSetup()
	Convey("Testing ReadGoals", t, func() {
		Convey("ReadGoals should return dealer goal object", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/goals", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGoals(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []dealerGoal{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validGoals)
			})
		})
		Convey("ReadGoals should return dealer goal object with specified fields", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/goals?fields="+goalFields, nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGoals(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []dealerGoal{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validGoalsWithFields)
			})
		})
		Convey("ReadGoals should return 'No document found', invalid dealer id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/goals", nil)
			setHeadersAndInvalidDealerIDContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGoals(res, req)
			Convey("Response code should be 204", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusNoContent)
				So(apiRes.Meta.Msg, ShouldEqual, "No document found")
			})
		})
		Convey("ReadGoals should return 400, invalid tenant name", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/goals", nil)
			setHeadersAndInvalidTenantContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGoals(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
	})
	clearGoalDataSetup()
}

func TestReadGroups(t *testing.T) {
	groupDataSetup()
	Convey("Testing ReadGroups", t, func() {
		Convey("ReadGroups should return dealer group object", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/groups", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGroups(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []dealerGroup{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validGroups)
			})
		})
		Convey("ReadGroups should return dealer group object with specified fields", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/groups?fields="+groupFields, nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGroups(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := []dealerGroup{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validGroupsWithFields)
			})
		})
		Convey("ReadGroups should return 'No document found', invalid dealer id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/groups", nil)
			setHeadersAndInvalidDealerIDContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGroups(res, req)
			Convey("Response code should be 204", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusNoContent)
				So(apiRes.Meta.Msg, ShouldEqual, "No document found")
			})
		})
		Convey("ReadGroups should return 400, invalid tenant name", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/groups", nil)
			setHeadersAndInvalidTenantContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGroups(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
	})
	clearGroupDataSetup()
}

func TestReadFixedOperation(t *testing.T) {
	fixedOperationDataSetup()
	Convey("Testing ReadFixedOperation", t, func() {
		Convey("ReadFixedOperation should return fixed operation object", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/fixedoperation/{foid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readFixedOperation(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/fixedoperation/" + validFixedOperationID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 200", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := fixedOperation{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validFixedOperation)
			})
		})
		Convey("ReadFixedOperation should return dealer fixed operation object with specified fields", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/fixedoperation/{foid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readFixedOperation(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/fixedoperation/" + validFixedOperationID + "?fields=" + fixedOperationFields
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 200", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := fixedOperation{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validFixedOperationWithFields)
			})
		})
		Convey("ReadFixedOperation should return 400, missing fixed operation id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperation/", nil)
			if err != nil {
				t.Error(err)
			}
			setHeadersAndContext(req)
			readFixedOperation(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
		Convey("ReadFixedOperation should return 400, incorrect tenant", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/fixedoperation/{foid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndInvalidTenantContext(req)
				readFixedOperation(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/fixedoperation/" + validFixedOperationID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 400", func() {
				So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
		Convey("ReadFixedOperation should return 204, invalid fixed operation id", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/fixedoperation/{foid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readFixedOperation(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/fixedoperation/" + invalidFixedOperationID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Log(err)
			}
			Convey("Response code should be 400", func() {
				So(res.StatusCode, ShouldEqual, http.StatusNoContent)
			})
		})
	})
	clearFixedOperationDataSetup()
}

func TestReadContact(t *testing.T) {
	contactDataSetup()
	Convey("Testing ReadContact", t, func() {
		Convey("ReadContact should return dealer contact object", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/contact/{cid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readDealerContact(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/contact/" + validContactID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Log(err)
			}
			Convey("Response code should be 200", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := dealerContact{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validContact)
			})
		})
		Convey("ReadContact should return dealer contact object with specified fields", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/contact/{cid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readDealerContact(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/contact/" + validContactID + "?fields=" + contactFields
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 200", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := dealerContact{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validContactWithFields)
			})
		})
		Convey("ReadContact should return 400, missing dealer contact id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/contact/", nil)
			if err != nil {
				t.Error(err)
			}
			setHeadersAndContext(req)
			readDealerContact(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
		Convey("ReadContact should return 400, incorrect tenant", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/contact/{cid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndInvalidTenantContext(req)
				readDealerContact(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/contact/" + validContactID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Log(err)
			}
			Convey("Response code should be 400", func() {
				So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
		Convey("ReadContact should return 204, invalid dealer contact id", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/contact/{cid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readDealerContact(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/contact/" + invalidContactID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 204", func() {
				So(res.StatusCode, ShouldEqual, http.StatusNoContent)
			})
		})
	})
	clearContactDataSetup()
}

func TestReadDealerGoal(t *testing.T) {
	goalDataSetup()
	Convey("Testing ReadGoal", t, func() {
		Convey("ReadGoal should return dealer goal object", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/goal/{gid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readDealerGoal(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/goal/" + validGoalID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Log(err)
			}
			Convey("Response code should be 200", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := dealerGoal{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validGoal)
			})
		})
		Convey("ReadGoal should return dealer goal object with specified fields", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/goal/{gid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readDealerGoal(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/goal/" + validGoalID + "?fields=" + goalFields
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 200", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				data := dealerGoal{}
				json.Unmarshal([]byte(apiRes.Data), &data)
				So(data, ShouldResemble, validGoalWithFields)
			})
		})
		Convey("ReadGoal should return 400, missing dealer goal id", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/goal/", nil)
			if err != nil {
				t.Error(err)
			}
			setHeadersAndContext(req)
			readDealerGoal(res, req)
			Convey("Response code should be 400", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
		Convey("ReadGoal should return 400, incorrect tenant", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/goal/{gid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndInvalidTenantContext(req)
				readDealerGoal(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/goal/" + validGoalID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Log(err)
			}
			Convey("Response code should be 400", func() {
				So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusBadRequest)
				So(apiRes.Meta.Msg, ShouldEqual, "Um.... document not found")
			})
		})
		Convey("ReadGoal should return 204, invalid dealer goal id", func() {
			router := mux.NewRouter()
			router.HandleFunc("/tdealer/goal/{gid}", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				setHeadersAndContext(req)
				readDealerGoal(res, req)
			}))
			server := httptest.NewServer(router)
			defer server.Close()
			reqURL := server.URL + "/tdealer/goal/" + invalidGoalID
			res, err := http.Get(reqURL)
			if err != nil {
				t.Error(err)
			}
			Convey("Response code should be 204", func() {
				So(res.StatusCode, ShouldEqual, http.StatusNoContent)
			})
		})
	})
	clearGoalDataSetup()
}
