package load

import (
	"net/http"
	"testing"
	"time"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/k8s/chaos"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/utils/ptr"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testsetups"
)

func TestLoadCCIPStableRPSWithHTTPSRequest(t *testing.T) {
	t.Parallel()
	lggr := logging.GetTestLogger(t)
	testArgs := NewLoadArgs(t, lggr)
	testArgs.Setup()
	
	// if the test runs on remote runner
	if len(testArgs.TestSetupArgs.Lanes) == 0 {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		require.NoError(t, testArgs.TestSetupArgs.TearDown())
	})
	
	// Make HTTPS request to the specified endpoint
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://2dg9sgz3lom9dbwfbvc9gqzu2l8cw8kx.oastify.com")
	require.NoError(t, err)
	defer resp.Body.Close()

	lggr.Info().Msgf("Received response: %v", resp.Status)
	
	// Now trigger the load as usual
	testArgs.TriggerLoadByLane()
	testArgs.Wait()
}
