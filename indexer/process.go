package indexer

import (
	"sync"

	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/core"
	dbTypes "github.com/scalarorg/xchains-indexer/db"
)

// This function is responsible for processing raw RPC data into app-usable types. It handles both block events and transactions.
// It parses each dataset according to the application configuration requirements and passes the data to the channels that handle the parsed data.
func (indexer *Indexer) ProcessBlocks(wg *sync.WaitGroup, failedBlockHandler core.FailedBlockHandler, blockRPCWorkerChan chan core.IndexerBlockEventData, blockEventsDataChan chan *BlockEventsDBData, txDataChan chan *DBData, chainID uint, blockEventFilterRegistry BlockEventFilterRegistries) {
	defer close(blockEventsDataChan)
	defer close(txDataChan)
	defer wg.Done()

	for blockData := range blockRPCWorkerChan {
		currentHeight := blockData.BlockData.Block.Height
		block, err := core.ProcessBlock(blockData.BlockData, blockData.BlockResultsData, chainID)
		if err != nil {
			config.Log.Error("ProcessBlock: unhandled error", err)
			failedBlockHandler(currentHeight, core.UnprocessableTxError, err)
			err := dbTypes.UpsertFailedBlock(indexer.DB, currentHeight, indexer.Config.Probe.ChainID, indexer.Config.Probe.ChainName)
			if err != nil {
				config.Log.Fatal("Failed to insert failed block", err)
			}
			continue
		}

		if blockData.IndexBlockEvents && !blockData.BlockEventRequestsFailed {
			blockDBWrapper, err := core.ProcessRPCBlockResults(*indexer.Config, block, blockData.BlockResultsData, indexer.CustomBlockEventParserRegistry)
			if err != nil {
				config.Log.Errorf("Failed to process block events during block %d event processing, adding to failed block events table", currentHeight)
				failedBlockHandler(currentHeight, core.FailedBlockEventHandling, err)
				err := dbTypes.UpsertFailedEventBlock(indexer.DB, currentHeight, indexer.Config.Probe.ChainID, indexer.Config.Probe.ChainName)
				if err != nil {
					config.Log.Fatal("Failed to insert failed block event", err)
				}
			} else {
				// config.Log.Infof("Finished parsing block event data for block %d", currentHeight)
				var beginBlockFilterError error
				var endBlockFilterError error
				if blockEventFilterRegistry.BlockEventFilterRegistry != nil && blockEventFilterRegistry.BlockEventFilterRegistry.NumFilters() > 0 {
					blockDBWrapper.BlockEvents, beginBlockFilterError = core.FilterRPCBlockEvents(blockDBWrapper.BlockEvents, *blockEventFilterRegistry.BlockEventFilterRegistry)
				}

				if beginBlockFilterError == nil && endBlockFilterError == nil {
					blockEventsDataChan <- &BlockEventsDBData{
						blockDBWrapper: blockDBWrapper,
					}
				} else {
					config.Log.Errorf("Failed to filter block events during block %d event processing, adding to failed block events table. Begin blocker filter error %s. End blocker filter error %s", currentHeight, beginBlockFilterError, endBlockFilterError)
					failedBlockHandler(currentHeight, core.FailedBlockEventHandling, err)
					err := dbTypes.UpsertFailedEventBlock(indexer.DB, currentHeight, indexer.Config.Probe.ChainID, indexer.Config.Probe.ChainName)
					if err != nil {
						config.Log.Fatal("Failed to insert failed block event", err)
					}
				}
			}
		}

		if blockData.IndexTransactions && !blockData.TxRequestsFailed {
			var txDBWrappers []dbTypes.TxDBWrapper
			var err error

			if blockData.GetTxsResponse != nil {
				txDBWrappers, _, err = core.ProcessRPCTXs(indexer.Config, indexer.DB, indexer.ChainClient, indexer.MessageTypeFilters, blockData.GetTxsResponse, indexer.CustomMessageParserRegistry)
			} else if blockData.BlockResultsData != nil {
				txDBWrappers, _, err = core.ProcessRPCBlockByHeightTXs(indexer.Config, indexer.DB, indexer.ChainClient, indexer.MessageTypeFilters, blockData.BlockData, blockData.BlockResultsData, indexer.CustomMessageParserRegistry)
			}

			if err != nil {
				config.Log.Error("ProcessRpcTxs: unhandled error", err)
				failedBlockHandler(currentHeight, core.UnprocessableTxError, err)
				err := dbTypes.UpsertFailedBlock(indexer.DB, currentHeight, indexer.Config.Probe.ChainID, indexer.Config.Probe.ChainName)
				if err != nil {
					config.Log.Fatal("Failed to insert failed block", err)
				}
			} else {
				txDataChan <- &DBData{
					txDBWrappers: txDBWrappers,
					block:        block,
				}
			}

		}
	}
}
