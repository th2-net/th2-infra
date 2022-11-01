package test

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/json"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultSchemaNamespace  = "th2-schema"
	defaultVhost            = "th2"
	defaultServiceNamespace = "service"
	monitoringNamespace     = "monitoring"
	rabbitmqPod             = "rabbitmq-0"
	rabbitmqSvc             = "rabbitmq-discovery"
	rabbitmqUser            = "th2"
	rabbitmqPassword        = "test"
	cassandraPod            = "cassandra-0"
	dataProviderSvc         = "rpt-data-provider"
	reportViewerSvc         = "rpt-data-viewer"
	infraMgrSvc             = "infra-mgr"
	infraOperatorSvc        = "infra-operator"
	infraMgrAppName         = "infra-mgr"
	infraOperatorAppName    = "infra-operator"
	infraEditorSvc          = "infra-editor"
	dashboardSvc            = "th2-infra-dashboard"
	retries                 = 10
	timeout                 = 5 * time.Second
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
			logger.Log(t, "Incorrect response code")
			return false
		}
		if !strings.Contains(body, substr) {
			logger.Log(t, "Body does not contain "+substr)
			return false
		}
		return true
	}
}

//func TestErrorsInMgr(t *testing.T) {
//	mgrLogs, mgrErr := getLogsFromApp(t, infraMgrAppName)
//	failIfErrExist(t, mgrErr, "kubectl logs infra-mgr failed")
//	assert.Truef(t, errorFreeLogs(mgrLogs), "infra-mgr contains errors")
//}
//
//func TestErrorsInOperator(t *testing.T) {
//	operatorLogs, operatorErr := getLogsFromApp(t, infraOperatorAppName)
//	t.Log(operatorLogs) //TODO: to remove
//	failIfErrExist(t, operatorErr, "kubectl logs infra-operator failed")
//	assert.Truef(t, errorFreeLogs(operatorLogs), "infra-operator contains errors")
//}

func TestDashboardEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/dashboard/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, retries, timeout)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestDashboardRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/dashboard"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, retries, timeout)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestInfraEditorEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraEditorSvc, retries, timeout)

	validator := validFunc(t, 200, "<title>Infra editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestInfraEditorRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraEditorSvc, retries, timeout)

	validator := validFunc(t, 200, "<title>Infra editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestRabbitMQEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/rabbitmq/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestRabbitMQRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/rabbitmq"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestInfraMgrEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/backend/actuator/health"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, cassandraPod, retries, timeout)

	validator := validFunc(t, 200, "{\"status\":\"UP\",\"groups\":[\"liveness\",\"readiness\"]}")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

//func TestNamespaceReportEndpoint(t *testing.T) {
//	// t.Parallel()
//	endpoint := fmt.Sprintf("http://localhost:30000/%s/", schemaNamespace)
//	options := k8s.NewKubectlOptions("", "", schemaNamespace)
//	k8s.WaitUntilServiceAvailable(t, options, reportViewerSvc, retries, timeout)
//
//	validator := validFunc(t, 200, "<title>TH2 Report</title>")
//	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
//}

func TestNamespaceDataProviderEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://localhost:30000/%s/backend/filters/sse-events", schemaNamespace)
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dataProviderSvc, retries, timeout)

	validator := validFunc(t, 200, "[\"attachedMessageId\",\"type\",\"name\",\"body\",\"status\"]")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestRabbitAllQueues(t *testing.T) {
	endpoint := fmt.Sprintf("http://%s:%s@localhost:30000/rabbitmq/api/queues/%s", rabbitmqUser, rabbitmqPassword, defaultVhost)
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)
	status, response := http_helper.HttpGet(t, endpoint, nil)
	if status != 200 {
		t.Fatalf("queues endpoint didn't return success (200 status code), instead it was: %d", status)
	}
	jsonByteArr := []byte(response)
	var jsonMaps []map[string]interface{}

	if err := json.Unmarshal(jsonByteArr, &jsonMaps); err != nil {
		t.Fatalf("Error occured during unmarshaling json: %s", err.Error())
	}

	makeName := func(resourceName string, pinName string) string {
		return fmt.Sprintf("link[%s:%s:%s]", defaultSchemaNamespace, resourceName, pinName)
	}

	expectedQueues := []string{
		makeName("mstore", "mstore-pin"),
		makeName("estore", "estore-pin"),
		makeName("codec-fix", "in_codec_encode"),
		makeName("codec-fix", "in_codec_decode"),
		makeName("codec-fix", "in_codec_general_encode"),
		makeName("codec-fix", "in_codec_general_decode"),
		makeName("rpt-data-provider", "from_codec"),
	}

	var actualQueues []string

	for _, jsonMap := range jsonMaps {
		actualQueues = append(actualQueues, fmt.Sprint(jsonMap["name"]))
	}

	sort.Strings(expectedQueues)
	sort.Strings(actualQueues)

	assert.Equal(t, expectedQueues, actualQueues)

}

//func TestRabbitMQQueues(t *testing.T) {
//	// t.Parallel()
//	endpoint := fmt.Sprintf("http://%s:%s@localhost:30000/rabbitmq/api/queues/%s/link%%5B%s%%3Arpt-data-provider%%3Afrom_codec%%5D",
//		rabbitmqUser, rabbitmqPassword, defaultVhost, schemaNamespace,
//	)
//	options := k8s.NewKubectlOptions("", "", serviceNamespace)
//	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)
//	expectedString := fmt.Sprintf("\"name\":\"link[%s:rpt-data-provider:from_codec]\"", schemaNamespace)
//	validator := validFunc(t, 200, expectedString)
//	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
//}

func TestPodAnnotations(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	filters := metav1.ListOptions{
		LabelSelector: "app=rpt-data-viewer",
	}
	pods := k8s.ListPods(t, options, filters)
	assert.Equal(t, "test-annotation", pods[0].ObjectMeta.Annotations["e2e"])
	assert.Equal(t, "test-common-annotation", pods[0].ObjectMeta.Annotations["e2ecommon"])
}

func TestPodCommonAnnotationsOnly(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	filters := metav1.ListOptions{
		LabelSelector: "app=rpt-data-provider",
	}
	pods := k8s.ListPods(t, options, filters)
	assert.Empty(t, pods[0].ObjectMeta.Annotations["e2e"])
	assert.Equal(t, "test-common-annotation", pods[0].ObjectMeta.Annotations["e2ecommon"])
}

func TestCodecFixPod(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	filters := metav1.ListOptions{
		LabelSelector: "app=codec-fix",
	}
	pods := k8s.ListPods(t, options, filters)
	codecFixPodName := pods[0].ObjectMeta.Name
	k8s.WaitUntilPodAvailable(t, options, codecFixPodName, retries, timeout)
}

func TestTH2Main(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/"
	validator := validFunc(t, 200, "<title>Welcome to th2</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestTH2MainAsset(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/assets/js/axios.min.js"
	validator := validFunc(t, 200, "/* axios v0.27.2 | (c) 2022 by Matt Zabriskie */")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}
