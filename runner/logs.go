package runner

import (
	"database/sql"
	"sync"

	"github.com/pganalyze/collector/grant"
	"github.com/pganalyze/collector/input"
	"github.com/pganalyze/collector/input/postgres"
	"github.com/pganalyze/collector/input/system/google_cloudsql"
	"github.com/pganalyze/collector/input/system/selfhosted"
	"github.com/pganalyze/collector/output"
	"github.com/pganalyze/collector/state"
	"github.com/pganalyze/collector/util"
	"github.com/pkg/errors"
)

func downloadLogsForServer(server state.Server, globalCollectionOpts state.CollectionOpts, logger *util.Logger) (state.PersistedLogState, bool, error) {
	newLogState := server.LogPrevState

	grant, err := grant.GetLogsGrant(server, globalCollectionOpts, logger)
	if err != nil {
		return newLogState, false, errors.Wrap(err, "could not get log grant")
	}

	if !grant.Valid {
		if globalCollectionOpts.TestRun {
			logger.PrintError("  Failed - Log Insights feature not available on this pganalyze plan, or log data limit exceeded. You may need to upgrade, see https://pganalyze.com/pricing")
		} else {
			logger.PrintVerbose("Skipping log data: Feature not available on this pganalyze plan, or log data limit exceeded")
		}
		return newLogState, false, nil
	}

	var connection *sql.DB
	if server.Config.EnableLogExplain {
		connection, err = postgres.EstablishConnection(server, logger, globalCollectionOpts, "")
		if err != nil {
			return newLogState, false, errors.Wrap(err, "failed to connect to database")
		}

		defer connection.Close()
	}

	transientLogState, persistedLogState, err := input.DownloadLogs(server, server.LogPrevState, connection, globalCollectionOpts, logger)
	if err != nil {
		transientLogState.Cleanup()
		return newLogState, false, errors.Wrap(err, "could not collect logs")
	}

	newLogState = persistedLogState

	err = output.UploadAndSendLogs(server, grant, globalCollectionOpts, logger, transientLogState)
	if err != nil {
		transientLogState.Cleanup()
		return newLogState, false, errors.Wrap(err, "failed to upload/send logs")
	}

	transientLogState.Cleanup()
	return newLogState, true, nil
}

// TestLogsForAllServers - Test log download/tailing
func TestLogsForAllServers(servers []state.Server, globalCollectionOpts state.CollectionOpts, logger *util.Logger) (hasFailedServers bool, hasSuccessfulLocalServers bool) {
	if !globalCollectionOpts.TestRun {
		return
	}

	for _, server := range servers {
		if server.Config.DisableLogs {
			continue
		}

		prefixedLogger := logger.WithPrefixAndRememberErrors(server.Config.SectionName)

		if server.Config.LogLocation != "" {
			prefixedLogger.PrintInfo("Testing log collection (local)...")
			err := selfhosted.TestLogTail(server, globalCollectionOpts, prefixedLogger)
			if err != nil {
				hasFailedServers = true
				prefixedLogger.PrintError("ERROR - Could not tail logs for server: %s", err)
			} else {
				prefixedLogger.PrintInfo("  Local log test successful")
				hasSuccessfulLocalServers = true
			}
		} else if server.Config.GcpCloudSQLInstanceID != "" && server.Config.GcpPubsubSubscription != "" {
			prefixedLogger.PrintInfo("Testing log collection (Google Cloud SQL)...")
			err := google_cloudsql.LogTestRun(server, globalCollectionOpts, prefixedLogger)
			if err != nil {
				hasFailedServers = true
				prefixedLogger.PrintError("ERROR - Could not get Pub/Sub log output for server: %s", err)
			} else {
				prefixedLogger.PrintInfo("  Log test successful")
			}
		} else if server.Config.AwsDbInstanceID != "" {
			prefixedLogger.PrintInfo("Testing log collection (RDS)...")
			_, _, err := downloadLogsForServer(server, globalCollectionOpts, prefixedLogger)
			if err != nil {
				hasFailedServers = true
				prefixedLogger.PrintError("Could not download logs for server: %s", err)
			}
		}
	}

	return
}

// DownloadLogsFromAllServers - Downloads logs from all servers that are remote systems and sends them to the pganalyze service
func DownloadLogsFromAllServers(servers []state.Server, globalCollectionOpts state.CollectionOpts, logger *util.Logger) {
	var wg sync.WaitGroup

	if !globalCollectionOpts.CollectLogs {
		return
	}

	for idx := range servers {
		if servers[idx].Config.DisableLogs || (servers[idx].Grant.Valid && !servers[idx].Grant.Config.EnableLogs) {
			continue
		}

		if servers[idx].Config.AwsDbInstanceID == "" {
			continue
		}

		wg.Add(1)
		go func(server *state.Server) {
			prefixedLogger := logger.WithPrefixAndRememberErrors(server.Config.SectionName)

			server.LogStateMutex.Lock()
			newLogState, success, err := downloadLogsForServer(*server, globalCollectionOpts, prefixedLogger)
			if err != nil {
				server.LogStateMutex.Unlock()
				prefixedLogger.PrintError("Could not collect logs for server: %s", err)
				if server.Config.ErrorCallback != "" {
					go runCompletionCallback("error", server.Config.ErrorCallback, server.Config.SectionName, "logs", err, prefixedLogger)
				}
			} else if success {
				server.LogPrevState = newLogState
				server.LogStateMutex.Unlock()
				if server.Config.SuccessCallback != "" {
					go runCompletionCallback("success", server.Config.SuccessCallback, server.Config.SectionName, "logs", nil, prefixedLogger)
				}
			}
			wg.Done()
		}(&servers[idx])
	}

	wg.Wait()

	return
}
