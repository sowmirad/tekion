package tdealerService
import (
	"bitbucket.org/tekion/tbaas/apiContext"
	"github.com/gorilla/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDealerByID(t *testing.T) {

	// Testing handler function
	req, err := http.NewRequest("GET", "http://localhost:8079/tdealer/getDealerById", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := apiContext.APIContext{Tenant: "Buck", DealerID: "3"}
	context.Set(req, "apiContext", ctx)
	req.Header.Add("tekion-api-token", "TestToken")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	GetDealerByID(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}