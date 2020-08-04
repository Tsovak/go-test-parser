package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tsovak/go-test-parser/models"
)

func TestSingleSuccessful(t *testing.T) {
	lines := []string{
		`{"Time":"2020-08-01T05:05:45.438058+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T05:05:45.438424+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== RUN   TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T05:05:45.459045+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== PAUSE TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T05:05:45.459075+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T05:05:45.459099+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T05:05:45.459105+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== CONT  TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T05:07:03.818099+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"--- PASS: TestAccAzureRMResourceGroup_basic (78.38s)\n"}`,
		`{"Time":"2020-08-01T05:07:03.818215+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Elapsed":78.38}`,
		`{"Time":"2020-08-01T05:07:03.818313+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"PASS\n"}`,
		`{"Time":"2020-08-01T05:07:03.828007+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"ok  \tgithub.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests\t78.429s\n"}`,
		`{"Time":"2020-08-01T05:07:03.860807+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Elapsed":78.462}`,
	}
	results := testParser(lines)
	require.Len(t, results, 1, "Expected a single result")

	result := results[0]
	require.Equal(t, result.Result, models.Successful, "Expected the result to be Successful")
	require.Equal(t, result.Duration, 78.38, "Expected the duration to be 78.38")

	expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_basic
=== PAUSE TestAccAzureRMResourceGroup_basic
=== CONT  TestAccAzureRMResourceGroup_basic
--- PASS: TestAccAzureRMResourceGroup_basic (78.38s)
`

	require.Equal(t, expectedStdOut, result.StdOut)
	require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
}

func TestSingleIgnored(t *testing.T) {
	lines := []string{
		`{"Time":"2020-08-01T12:48:32.302781+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport"}`,
		`{"Time":"2020-08-01T12:48:32.30319+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Output":"=== RUN   TestAccAzureRMResourceGroup_requiresImport\n"}`,
		`{"Time":"2020-08-01T12:48:32.30322+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Output":"--- SKIP: TestAccAzureRMResourceGroup_requiresImport (0.00s)\n"}`,
		`{"Time":"2020-08-01T12:48:32.303229+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Output":"    resource_arm_resource_group_test.go:36: Skipping since resources aren't required to be imported\n"}`,
		`{"Time":"2020-08-01T12:48:32.303239+01:00","Action":"skip","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Elapsed":0}`,
		`{"Time":"2020-08-01T12:48:32.303264+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"PASS\n"}`,
		`{"Time":"2020-08-01T12:48:32.304518+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"ok  \tgithub.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests\t0.029s\n"}`,
		`{"Time":"2020-08-01T12:48:32.308532+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Elapsed":0.033}`,
	}

	results := testParser(lines)
	require.Len(t, results, 1)

	result := results[0]
	require.Equal(t, result.Result, models.Ignored)
	require.Equal(t, result.Duration, 0.0)

	expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_requiresImport
--- SKIP: TestAccAzureRMResourceGroup_requiresImport (0.00s)
    resource_arm_resource_group_test.go:36: Skipping since resources aren't required to be imported
`
	require.Equal(t, expectedStdOut, result.StdOut)
	require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
}

func TestSingleFailed(t *testing.T) {
	lines := []string{
		`{"Time":"2020-08-01T12:44:04.182228+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:44:04.182571+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== RUN   TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:44:04.206855+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== PAUSE TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:44:04.206929+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:44:04.206961+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:44:04.206969+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== CONT  TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:44:04.333766+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"--- FAIL: TestAccAzureRMResourceGroup_basic (0.15s)\n"}`,
		`{"Time":"2020-08-01T12:44:04.333798+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"    testing.go:569: Step 0 error: config is invalid: Invalid resource type: The provider provider.azurerm does not support resource type \"azurerm_resource_group\".\n"}`,
		`{"Time":"2020-08-01T12:44:04.333822+01:00","Action":"fail","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Elapsed":0.15}`,
		`{"Time":"2020-08-01T12:44:04.333843+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"FAIL\n"}`,
		`{"Time":"2020-08-01T12:44:04.336708+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"FAIL\tgithub.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests\t0.180s\n"}`,
		`{"Time":"2020-08-01T12:44:04.33689+01:00","Action":"fail","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Elapsed":0.18}`,
	}
	results := testParser(lines)
	require.Len(t, results, 1, "Expected a single result")

	result := results[0]
	require.Equal(t, result.Result, models.Failed, "Expected the result to be Failed")
	require.Equal(t, 0.15, result.Duration)

	expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_basic
=== PAUSE TestAccAzureRMResourceGroup_basic
=== CONT  TestAccAzureRMResourceGroup_basic
--- FAIL: TestAccAzureRMResourceGroup_basic (0.15s)
    testing.go:569: Step 0 error: config is invalid: Invalid resource type: The provider provider.azurerm does not support resource type "azurerm_resource_group".
`
	require.Equal(t, expectedStdOut, result.StdOut)
	require.Equal(t, expectedStdOut, result.StdOut)
	require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
}

func TestSingleFailedMultipleLines(t *testing.T) {
	lines := []string{
		`{"Time":"2020-08-01T12:51:47.588077+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:51:47.588388+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== RUN   TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:51:47.608038+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== PAUSE TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:51:47.608092+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:51:47.608107+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:51:47.608113+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== CONT  TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:51:47.611735+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"--- FAIL: TestAccAzureRMResourceGroup_basic (0.02s)\n"}`,
		`{"Time":"2020-08-01T12:51:47.611781+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"    testing.go:569: Step 0 error: Error initializing context: 2 problems:\n"}`,
		`{"Time":"2020-08-01T12:51:47.611796+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        \n"}`,
		`{"Time":"2020-08-01T12:51:47.611827+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        - Could not satisfy plugin requirements: \n"}`,
		`{"Time":"2020-08-01T12:51:47.611848+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        Plugin reinitialization required. Please run \"terraform init\".\n"}`,
		`{"Time":"2020-08-01T12:51:47.611856+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        \n"}`,
		`{"Time":"2020-08-01T12:51:47.611953+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        Plugins are external binaries that Terraform uses to access and manipulate\n"}`,
		`{"Time":"2020-08-01T12:51:47.611968+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        resources. The configuration provided requires plugins which can't be located,\n"}`,
		`{"Time":"2020-08-01T12:51:47.611975+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        don't satisfy the version constraints, or are otherwise incompatible.\n"}`,
		`{"Time":"2020-08-01T12:51:47.612092+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        \n"}`,
		`{"Time":"2020-08-01T12:51:47.612108+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        Terraform automatically discovers provider requirements from your\n"}`,
		`{"Time":"2020-08-01T12:51:47.612136+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        configuration, including providers used in child modules. To see the\n"}`,
		`{"Time":"2020-08-01T12:51:47.612146+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        requirements and constraints from each module, run \"terraform providers\".\n"}`,
		`{"Time":"2020-08-01T12:51:47.612153+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        \n"}`,
		`{"Time":"2020-08-01T12:51:47.61216+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"        - provider \"azurerm\" is not available\n"}`,
		`{"Time":"2020-08-01T12:51:47.612173+01:00","Action":"fail","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Elapsed":0.02}`,
		`{"Time":"2020-08-01T12:51:47.612194+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"FAIL\n"}`,
		`{"Time":"2020-08-01T12:51:47.613587+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"FAIL\tgithub.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests\t0.053s\n"}`,
		`{"Time":"2020-08-01T12:51:47.613711+01:00","Action":"fail","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Elapsed":0.053}`,
	}
	results := testParser(lines)
	require.Len(t, results, 1)

	result := results[0]
	require.Equal(t, result.Result, models.Failed)
	require.Equal(t, result.Duration, 0.02)

	expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_basic
=== PAUSE TestAccAzureRMResourceGroup_basic
=== CONT  TestAccAzureRMResourceGroup_basic
--- FAIL: TestAccAzureRMResourceGroup_basic (0.02s)
    testing.go:569: Step 0 error: Error initializing context: 2 problems:
        
        - Could not satisfy plugin requirements: 
        Plugin reinitialization required. Please run "terraform init".
        
        Plugins are external binaries that Terraform uses to access and manipulate
        resources. The configuration provided requires plugins which can't be located,
        don't satisfy the version constraints, or are otherwise incompatible.
        
        Terraform automatically discovers provider requirements from your
        configuration, including providers used in child modules. To see the
        requirements and constraints from each module, run "terraform providers".
        
        - provider "azurerm" is not available
`
	require.Equal(t, expectedStdOut, result.StdOut)
	require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
}

func TestSinglePanic(t *testing.T) {
	lines := []string{
		`{"Time":"2020-08-01T12:57:33.387711+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:57:33.388053+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== RUN   TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:57:33.40788+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== PAUSE TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:57:33.40794+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:57:33.407959+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T12:57:33.407964+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== CONT  TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T12:57:33.407978+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"--- FAIL: TestAccAzureRMResourceGroup_basic (0.02s)\n"}`,
		`{"Time":"2020-08-01T12:57:33.410433+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"panic: Hello [recovered]\n"}`,
		`{"Time":"2020-08-01T12:57:33.41049+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\tpanic: Hello\n"}`,
		`{"Time":"2020-08-01T12:57:33.4105+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\n"}`,
		`{"Time":"2020-08-01T12:57:33.410521+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"goroutine 23 [running]:\n"}`,
		`{"Time":"2020-08-01T12:57:33.410535+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"testing.tRunner.func1(0xc000130300)\n"}`,
		`{"Time":"2020-08-01T12:57:33.410564+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/usr/local/Cellar/go/1.13.5/libexec/src/testing/testing.go:874 +0x3a3\n"}`,
		`{"Time":"2020-08-01T12:57:33.410578+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"panic(0x459bc80, 0x52ea040)\n"}`,
		`{"Time":"2020-08-01T12:57:33.410589+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/usr/local/Cellar/go/1.13.5/libexec/src/runtime/panic.go:679 +0x1b2\n"}`,
		`{"Time":"2020-08-01T12:57:33.410682+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests.TestAccAzureRMResourceGroup_basic.func1()\n"}`,
		`{"Time":"2020-08-01T12:57:33.410695+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/Users/tharvey/code/src/github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests/resource_arm_resource_group_test.go:19 +0x39\n"}`,
		`{"Time":"2020-08-01T12:57:33.410715+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"github.com/hashicorp/terraform-plugin-sdk/helper/resource.Test(0x53d4a80, 0xc000130300, 0x0, 0x4dc16f0, 0xc0008977a0, 0x0, 0x0, 0x4dc17a8, 0xc0006eb728, 0x2, ...)\n"}`,
		`{"Time":"2020-08-01T12:57:33.410724+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/Users/tharvey/code/pkg/mod/github.com/hashicorp/terraform-plugin-sdk@v1.1.1/helper/resource/testing.go:482 +0x179b\n"}`,
		`{"Time":"2020-08-01T12:57:33.410742+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"github.com/hashicorp/terraform-plugin-sdk/helper/resource.ParallelTest(0x53d4a80, 0xc000130300, 0x0, 0x4dc16f0, 0xc0008977a0, 0x0, 0x0, 0x4dc17a8, 0xc0006eb728, 0x2, ...)\n"}`,
		`{"Time":"2020-08-01T12:57:33.430902+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/Users/tharvey/code/pkg/mod/github.com/hashicorp/terraform-plugin-sdk@v1.1.1/helper/resource/testing.go:444 +0x83\n"}`,
		`{"Time":"2020-08-01T12:57:33.430934+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests.TestAccAzureRMResourceGroup_basic(0xc000130300)\n"}`,
		`{"Time":"2020-08-01T12:57:33.430943+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/Users/tharvey/code/src/github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests/resource_arm_resource_group_test.go:18 +0x457\n"}`,
		`{"Time":"2020-08-01T12:57:33.430981+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"testing.tRunner(0xc000130300, 0x4dc16f8)\n"}`,
		`{"Time":"2020-08-01T12:57:33.430997+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/usr/local/Cellar/go/1.13.5/libexec/src/testing/testing.go:909 +0xc9\n"}`,
		`{"Time":"2020-08-01T12:57:33.431007+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"created by testing.(*T).Run\n"}`,
		`{"Time":"2020-08-01T12:57:33.431032+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"\t/usr/local/Cellar/go/1.13.5/libexec/src/testing/testing.go:960 +0x350\n"}`,
		`{"Time":"2020-08-01T12:57:33.431123+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"FAIL\tgithub.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests\t0.071s\n"}`,
		`{"Time":"2020-08-01T12:57:33.43165+01:00","Action":"fail","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Elapsed":0.071}`,
	}
	results := testParser(lines)
	require.Len(t, results, 1)

	result := results[0]
	require.Equal(t, result.Result, models.Failed)
	require.Equal(t, result.Duration, 0.071)

	expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_basic
=== PAUSE TestAccAzureRMResourceGroup_basic
=== CONT  TestAccAzureRMResourceGroup_basic
--- FAIL: TestAccAzureRMResourceGroup_basic (0.02s)
panic: Hello [recovered]
	panic: Hello

goroutine 23 [running]:
testing.tRunner.func1(0xc000130300)
	/usr/local/Cellar/go/1.13.5/libexec/src/testing/testing.go:874 +0x3a3
panic(0x459bc80, 0x52ea040)
	/usr/local/Cellar/go/1.13.5/libexec/src/runtime/panic.go:679 +0x1b2
github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests.TestAccAzureRMResourceGroup_basic.func1()
	/Users/tharvey/code/src/github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests/resource_arm_resource_group_test.go:19 +0x39
github.com/hashicorp/terraform-plugin-sdk/helper/resource.Test(0x53d4a80, 0xc000130300, 0x0, 0x4dc16f0, 0xc0008977a0, 0x0, 0x0, 0x4dc17a8, 0xc0006eb728, 0x2, ...)
	/Users/tharvey/code/pkg/mod/github.com/hashicorp/terraform-plugin-sdk@v1.1.1/helper/resource/testing.go:482 +0x179b
github.com/hashicorp/terraform-plugin-sdk/helper/resource.ParallelTest(0x53d4a80, 0xc000130300, 0x0, 0x4dc16f0, 0xc0008977a0, 0x0, 0x0, 0x4dc17a8, 0xc0006eb728, 0x2, ...)
	/Users/tharvey/code/pkg/mod/github.com/hashicorp/terraform-plugin-sdk@v1.1.1/helper/resource/testing.go:444 +0x83
github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests.TestAccAzureRMResourceGroup_basic(0xc000130300)
	/Users/tharvey/code/src/github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests/resource_arm_resource_group_test.go:18 +0x457
testing.tRunner(0xc000130300, 0x4dc16f8)
	/usr/local/Cellar/go/1.13.5/libexec/src/testing/testing.go:909 +0xc9
created by testing.(*T).Run
	/usr/local/Cellar/go/1.13.5/libexec/src/testing/testing.go:960 +0x350
FAIL	github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests	0.071s
`
	require.Equal(t, expectedStdOut, result.StdOut)
	require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
}

func TestMultipleTestsSamePackage(t *testing.T) {
	lines := []string{
		`{"Time":"2020-08-01T04:45:17.228802+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T04:45:17.231582+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== RUN   TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T04:45:17.25869+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== PAUSE TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T04:45:17.258756+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T04:45:17.25878+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport"}`,
		`{"Time":"2020-08-01T04:45:17.258799+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Output":"=== RUN   TestAccAzureRMResourceGroup_requiresImport\n"}`,
		`{"Time":"2020-08-01T04:45:17.259549+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Output":"--- SKIP: TestAccAzureRMResourceGroup_requiresImport (0.00s)\n"}`,
		`{"Time":"2020-08-01T04:45:17.259594+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Output":"    resource_arm_resource_group_test.go:36: Skipping since resources aren't required to be imported\n"}`,
		`{"Time":"2020-08-01T04:45:17.259629+01:00","Action":"skip","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_requiresImport","Elapsed":0}`,
		`{"Time":"2020-08-01T04:45:17.259665+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears"}`,
		`{"Time":"2020-08-01T04:45:17.259683+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears","Output":"=== RUN   TestAccAzureRMResourceGroup_disappears\n"}`,
		`{"Time":"2020-08-01T04:45:17.277204+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears","Output":"=== PAUSE TestAccAzureRMResourceGroup_disappears\n"}`,
		`{"Time":"2020-08-01T04:45:17.277263+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears"}`,
		`{"Time":"2020-08-01T04:45:17.27729+01:00","Action":"run","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags"}`,
		`{"Time":"2020-08-01T04:45:17.277306+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags","Output":"=== RUN   TestAccAzureRMResourceGroup_withTags\n"}`,
		`{"Time":"2020-08-01T04:45:17.300809+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags","Output":"=== PAUSE TestAccAzureRMResourceGroup_withTags\n"}`,
		`{"Time":"2020-08-01T04:45:17.301086+01:00","Action":"pause","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags"}`,
		`{"Time":"2020-08-01T04:45:17.301151+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic"}`,
		`{"Time":"2020-08-01T04:45:17.301196+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"=== CONT  TestAccAzureRMResourceGroup_basic\n"}`,
		`{"Time":"2020-08-01T04:45:17.302446+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags"}`,
		`{"Time":"2020-08-01T04:45:17.302486+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags","Output":"=== CONT  TestAccAzureRMResourceGroup_withTags\n"}`,
		`{"Time":"2020-08-01T04:45:17.302503+01:00","Action":"cont","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears"}`,
		`{"Time":"2020-08-01T04:45:17.302519+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears","Output":"=== CONT  TestAccAzureRMResourceGroup_disappears\n"}`,
		`{"Time":"2020-08-01T04:46:27.433877+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears","Output":"--- PASS: TestAccAzureRMResourceGroup_disappears (70.15s)\n"}`,
		`{"Time":"2020-08-01T04:46:34.422914+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_disappears","Elapsed":70.15}`,
		`{"Time":"2020-08-01T04:46:34.422995+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Output":"--- PASS: TestAccAzureRMResourceGroup_basic (77.15s)\n"}`,
		`{"Time":"2020-08-01T04:47:04.104737+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_basic","Elapsed":77.15}`,
		`{"Time":"2020-08-01T04:47:04.104807+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags","Output":"--- PASS: TestAccAzureRMResourceGroup_withTags (106.83s)\n"}`,
		`{"Time":"2020-08-01T04:47:04.104869+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Test":"TestAccAzureRMResourceGroup_withTags","Elapsed":106.83}`,
		`{"Time":"2020-08-01T04:47:04.104901+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"PASS\n"}`,
		`{"Time":"2020-08-01T04:47:04.11591+01:00","Action":"output","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Output":"ok  \tgithub.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests\t106.926s\n"}`,
		`{"Time":"2020-08-01T04:47:04.189611+01:00","Action":"pass","Package":"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/tests","Elapsed":107}`,
	}

	results := testParser(lines)
	require.Len(t, results, 4)

	for _, result := range results {
		switch result.TestName {
		case "TestAccAzureRMResourceGroup_basic":
			require.Equal(t, result.Result, models.Successful)
			require.Equal(t, result.Duration, 77.15)

			expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_basic
=== PAUSE TestAccAzureRMResourceGroup_basic
=== CONT  TestAccAzureRMResourceGroup_basic
--- PASS: TestAccAzureRMResourceGroup_basic (77.15s)
`
			require.Equal(t, expectedStdOut, result.StdOut)
			require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
			continue

		case "TestAccAzureRMResourceGroup_disappears":
			require.Equal(t, result.Result, models.Successful)
			require.Equal(t, result.Duration, 70.15)

			expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_disappears
=== PAUSE TestAccAzureRMResourceGroup_disappears
=== CONT  TestAccAzureRMResourceGroup_disappears
--- PASS: TestAccAzureRMResourceGroup_disappears (70.15s)
`
			require.Equal(t, expectedStdOut, result.StdOut)
			require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
			continue

		case "TestAccAzureRMResourceGroup_requiresImport":
			require.Equal(t, result.Result, models.Ignored)
			require.Equal(t, result.Duration, 0.0)
			expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_requiresImport
--- SKIP: TestAccAzureRMResourceGroup_requiresImport (0.00s)
    resource_arm_resource_group_test.go:36: Skipping since resources aren't required to be imported
`
			require.Equal(t, expectedStdOut, result.StdOut)
			require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
			continue

		case "TestAccAzureRMResourceGroup_withTags":
			require.Equal(t, result.Result, models.Successful)
			require.Equal(t, result.Duration, 106.83)

			expectedStdOut := `=== RUN   TestAccAzureRMResourceGroup_withTags
=== PAUSE TestAccAzureRMResourceGroup_withTags
=== CONT  TestAccAzureRMResourceGroup_withTags
--- PASS: TestAccAzureRMResourceGroup_withTags (106.83s)
`
			require.Equal(t, expectedStdOut, result.StdOut)
			require.Len(t, result.StdErr, 0, "Expected the stderr to be empty")
			continue

		default:
			t.Fatalf("Unknown Test %q", result.TestName)
		}
	}
}

func TestMultipleTestsDifferentPackages(t *testing.T) {
	lines := multipleTestsDifferentPackageConfig
	results := testParser(lines)
	require.Len(t, results, 11, "Expected 11 tests")

	uniquePackages := make(map[string]struct{}, 0)
	for _, r := range results {
		uniquePackages[r.Package] = struct{}{}
	}
	require.Len(t, uniquePackages, 2, "Expected 2 packages")
}

func testParser(lines []string) []models.TestResult {
	results := make([]models.TestResult, 0)
	parser := NewResultsParser(func(result models.TestResult) {
		results = append(results, result)
	})

	for _, v := range lines {
		parser.ParseLine(v, true)
	}

	return results
}
