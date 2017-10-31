package dealerService

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	//"strings"
	"testing"
	//"github.com/smartystreets/goconvey/convey"
	//"github.com/go-kit/kit/transport/grpc"

	"github.com/golang/go/src/pkg/bytes"
	//grpc2 "google.golang.org/grpc"
	"bitbucket.org/tekion/tbaas/tapi"
)

var (
	TenantName = "Test"
	DealerID   = "3"
	ClientID   = "web"
	ctx        = apiContext.APIContext{Tenant: TenantName, DealerID: DealerID, ClientID: ClientID}
)

func TestReadContacts(t *testing.T) {
	contactDataSetup()
	Convey("Testing readContacts", t, func() {
		Convey("readContacts should return dealer contact objects", func() {
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
		Convey("readContacts should return dealer contact objects with specified fields", func() {
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
		Convey("readContacts should return 400, invalid tenant name", func() {
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
		Convey("readContacts should return 204, no dealer contact data in db", func() {
			clearContactDataSetup()
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/contacts", nil)
			setHeadersAndContext(req)
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
	})
}

func TestReadGoals(t *testing.T) {
	goalDataSetup()
	Convey("Testing readGoals", t, func() {
		Convey("readGoals should return dealer goal objects", func() {
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
		Convey("readGoals should return dealer goal objects with specified fields", func() {
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
		Convey("readGoals should return 'No document found'", func() {
			clearGoalDataSetup()
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/goals", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readDealerGoals(res, req)
			Convey("Response code should be 204, no dealer goal data in db", func() {
				apiRes := apiResponse{}
				err := json.NewDecoder(res.Body).Decode(&apiRes)
				if err != nil {
					t.Error(err)
				}
				So(apiRes.Meta.Code, ShouldEqual, http.StatusNoContent)
				So(apiRes.Meta.Msg, ShouldEqual, "No document found")
			})
		})
	})
}

func TestReadGroups(t *testing.T) {
	groupDataSetup()
	Convey("Testing readGroups", t, func() {
		Convey("readGroups should return dealer group objects", func() {
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
		Convey("readGroups should return dealer group objects with specified fields", func() {
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
		Convey("readGroups should return 400, invalid tenant name", func() {
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
		Convey("readGroups should return 204, no dealer group data in db", func() {
			clearGroupDataSetup()
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/groups", nil)
			setHeadersAndContext(req)
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
	})
}

func TestReadDealer(t *testing.T) {
	dealerDataSetup()
	Convey("Testing readDealer", t, func() {
		Convey("readDealer should return dealer object", func() {
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
		Convey("readDealer should return dealer object with specified fields", func() {
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
		Convey("readDealer should return status code 400, invalid tenant name", func() {
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
		Convey("readDealer should return status code 204, no dealer data in db. `impossible case ?`", func() {
			clearDealerDataSetup()
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/dealer", nil)
			setHeadersAndContext(req)
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
	})
}

func TestReadFixedOperation(t *testing.T) {
	fixedOperationDataSetup()
	Convey("Testing readFixedOperation", t, func() {
		Convey("readFixedOperation should return fixed operation object", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperation", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperation(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
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
		Convey("readFixedOperation should return fixed operation object with specified fields", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperation?fields="+fixedOperationFields, nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperation(res, req)
			Convey("Response code should be 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
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
		Convey("readFixedOperation should return 400, invalid tenant name", func() {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperation", nil)
			setHeadersAndInvalidTenantContext(req)
			if err != nil {
				t.Error(err)
			}
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
		Convey("readFixedOperation should return 204, no fixed operation data in db", func() {
			clearFixedOperationDataSetup()
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/tdealer/fixedoperation", nil)
			setHeadersAndContext(req)
			if err != nil {
				t.Error(err)
			}
			readFixedOperation(res, req)
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
	})
}

func TestReadContact(t *testing.T) {
	contactDataSetup()
	Convey("Testing readContact", t, func() {
		Convey("readContact should return dealer contact object", func() {
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
		Convey("readContact should return dealer contact object with specified fields", func() {
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
		Convey("readContact should return 400, incorrect tenant", func() {
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
		Convey("readContact should return 204, invalid dealer contact id", func() {
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
func SetHeader() http.Header {
	header := http.Header{}
	header.Set(TenantName, ctx.Tenant)
	header.Set(DealerID, ctx.DealerID)
	header.Set(tapi.TekionAPIToken, "TestToken")
	header.Set(ClientID, ctx.ClientID)

	return header
}

func TestDealerUpdate(t *testing.T) {
	insertDealerData()
	Update_api_check := dealer{
		ID:       "3",
		Name:     "TestingDealer",
		MakeCode: []string{"make1"},

		DoingBusinessAsName:      "business_name_as",
		StateIssuedNumber:        "234556",
		ManufacturerIssuedNumber: "24456677",
		Website:                  "http://dealer.com",
		TimeZone:                 "us-pecific",
		Currency:                 "usd",
		TenantID:                 "88",
		Phone:                    "8983833939",
	}
	Convey("check if inserted data is updates", func() {
		res := httptest.NewRecorder()
		inputObject, _ := json.Marshal(Update_api_check)
		req, err := http.NewRequest("POST", "/tdealer/updatedealer", bytes.NewBuffer(inputObject))
		req.Header = SetHeader()
		context.Set(req, "apiContext", ctx)
		if nil != err {
			t.Error(err)
		}
		updateDealer(res, req)
		Convey("Response code should be 200", func() {
			So(res.Code, ShouldEqual, http.StatusOK)
		})
	})
	clearDealerData()
}
