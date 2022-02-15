package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
)

const (
	defaultSchemaNamespace  = "th2-schema"
	defaultServiceNamespace = "service"
	monitoringNamespace     = "monitoring"
	rabbitmqSvc             = "rabbitmq-discovery"
	dataProviderSvc         = "rpt-data-provider"
	reportViewerSvc         = "rpt-data-viewer"
	infraMgrSvc             = "infra-mgr"
	infraEditorSvc          = "infra-editor"
	dashboardSvc            = "dashboard-kubernetes-dashboard"
)

var (
	serviceNamespace, schemaNamespace string
	exists                            bool
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	if schemaNamespace, exists = os.LookupEnv("SCHEMA_NAMESPACE"); !exists {
		schemaNamespace = defaultSchemaNamespace
	}
	if serviceNamespace, exists = os.LookupEnv("INFRA_NAMESPACE"); !exists {
		serviceNamespace = defaultServiceNamespace
	}
}

func tearDown() {}

func validFunc(t *testing.T, testCode int, substr string) func(int, string) bool {
	return func(code int, body string) bool {
		if testCode != code {
			logger.Logf(t, "Incorrect response code")
			return false
		}
		if !strings.Contains(body, substr) {
			logger.Logf(t, "Body does not contain "+substr)
			return false
		}
		return true
	}
}

func TestDashboardEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/dashboard/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, 10, 3*time.Second)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

// func TestDashboardRedirectEndpoint(t *testing.T) {
// 	// t.Parallel()
// 	endpoint := "http://localhost:30000/dashboard"
// 	options := k8s.NewKubectlOptions("", "", monitoringNamespace)
// 	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, 10, 3*time.Second)
// 	validator := validFunc(t, 301, "")
// 	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
// }

func TestInfraEditorEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraEditorSvc, 10, 1*time.Second)

	validator := validFunc(t, 200, "<title>Infra editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestRabbitMQEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/rabbitmq/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, rabbitmqSvc, 10, 3*time.Second)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestInfraMgrEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/backend/actuator/health"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraMgrSvc, 10, 1*time.Second)

	validator := validFunc(t, 200, "{\"status\":\"UP\",\"groups\":[\"liveness\",\"readiness\"]}")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 10, 3*time.Second, validator)
}

func TestNamespaceReportEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://localhost:30000/%s/", schemaNamespace)
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	k8s.WaitUntilServiceAvailable(t, options, reportViewerSvc, 30, 10*time.Second)

	validator := validFunc(t, 200, "<title>TH2 Report</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 10, 3*time.Second, validator)
}

//func TestNamespaceDataProviderEndpoint(t *testing.T) {
// t.Parallel()
//	endpoint := fmt.Sprintf("http://localhost:30000/%s/backend/search/events?timestampFrom=0&timestampTo=0", schemaNamespace)
//	options := k8s.NewKubectlOptions("", "", schemaNamespace)
//	k8s.WaitUntilServiceAvailable(t, options, dataProviderSvc, 30, 10*time.Second)
//
//	validator := validFunc(t, 200, "[]")
//	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 10, 10*time.Second, validator)
//}
