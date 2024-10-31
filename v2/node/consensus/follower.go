package consensus

import (
	"context"
	"fmt"
	"kwil/node/types"
)

// AcceptProposal checks if the node should download the block corresponding to the proposal.
// This should not be processed by the leader and the sentry nodes.
// Validator should only accept the proposal if it is not already processing a block and
// the proposal is for the next block to be processed.
// If we receive a conflicting proposal, abort the execution of the current proposal and
// start processing the new proposal.
func (ce *ConsensusEngine) AcceptProposal(height int64, blkID, prevBlockID types.Hash) bool {
	ce.log.Info("Accept proposal?", "height", height, "blkID", blkID, "prevHash", prevBlockID)
	ce.state.mtx.Lock()
	defer ce.state.mtx.Unlock()

	ce.updateNetworkHeight(height - 1)

	if ce.role != types.RoleValidator {
		return false
	}

	// initial block must precede genesis
	if height == 1 || prevBlockID.IsZero() {
		return ce.state.lc.blkHash.IsZero()
	}

	// Check if this is the next block to be processed
	if height != ce.state.lc.height+1 {
		ce.log.Info("Block proposal is not for the next height", "blkPropHeight", height, "expected", ce.state.lc.height+1)
		return false
	}

	// Check if the validator is busy processing a block.
	if ce.state.blkProp != nil {
		// check if we are processing a different block, if yes then reset the state.
		if ce.state.blkProp.blkHash != blkID {
			ce.log.Info("Conflicting block proposals, abort block execution and requesting the new block: ", "height", height)
			msg := &resetState{height: height - 1}
			ce.sendConsensusMessage(&consensusMessage{
				MsgType: msg.Type(),
				Msg:     msg,
			})
			return true
		}
		ce.log.Debug("Already processing the block proposal", "height", height, "blkID", blkID)
		return false
	}

	return prevBlockID == ce.state.lc.blkHash
}

// NotifyBlockProposal is used by the p2p stream handler to notify the consensus engine of a new block proposal.
// Only a validator should use this method, not leader or sentry. This method does it's best to ensure that this
// is the next block to be processed, only then it notifies the consensus engine of the block proposal.
// respCb is a callback function used to send the VoteMessage(ack/nack) back to the leader.
func (ce *ConsensusEngine) NotifyBlockProposal(blk *types.Block) {
	// ce.log.Infoln("Notify block proposal", blk.Header.Height, blk.Header.Hash())
	if ce.role == types.RoleLeader {
		return
	}

	ce.state.mtx.Lock()
	defer ce.state.mtx.Unlock()

	if blk.Header.Height != ce.state.lc.height+1 {
		ce.log.Infof("proposal for height %d does not follow %d", blk.Header.Height, ce.state.lc.height)
		return
	}

	if ce.state.blkProp != nil {
		if ce.state.blkProp.blkHash != blk.Header.Hash() {
			ce.log.Info("Conflicting block proposals, abort block execution and requesting the new block: ", "height", blk.Header.Height)
			msg := &resetState{height: blk.Header.Height - 1}
			ce.sendConsensusMessage(&consensusMessage{
				MsgType: msg.Type(),
				Msg:     msg,
			})
			return
		}
		ce.log.Info("block proposal already exists")
		return
	}

	blkProp := &blockProposal{
		height:  blk.Header.Height,
		blkHash: blk.Header.Hash(),
		blk:     blk,
	}

	ce.sendConsensusMessage(&consensusMessage{
		MsgType: blkProp.Type(),
		Msg:     blkProp,
		Sender:  ce.pubKey,
	})
}

// AcceptCommit handles the blockAnnounce message from the leader.
// This should be processed only if this is the next block to be committed by the node.
// This also checks if the node should request the block from its peers. This can happen
// 1. If the node is a sentry node and doesn't have the block.
// 2. If the node is a validator and missed the block proposal message.
func (ce *ConsensusEngine) AcceptCommit(height int64, blkID types.Hash, appHash types.Hash) bool {
	// ce.log.Infoln("Accept commit?", height, blkID, appHash)
	if ce.role == types.RoleLeader {
		return false
	}

	ce.state.mtx.Lock()
	defer ce.state.mtx.Unlock()

	ce.updateNetworkHeight(height)

	ce.log.Info(fmt.Sprintln("Accept commit?", height, blkID, appHash, ce.state.lc.height+1)) // maybe we add log.Infoln?
	if ce.state.lc.height+1 != height {
		// This is not the next block to be committed by the node.
		return false
	}

	blkCommit := &blockAnnounce{
		appHash: appHash,
	}

	if ce.state.blkProp != nil {
		// ensure that we are processing the correct block
		if ce.state.blkProp.blkHash != blkID {
			// Rollback and reprocess the block sent as part of the commit message.
			ce.log.Info("Processing incorrect block, notify consensus engine to abort: ", "height", height)
			msg := &resetState{height: height - 1}
			ce.sendConsensusMessage(&consensusMessage{
				MsgType: msg.Type(),
				Msg:     msg,
			})
			return true // fetch the correct block
		}

		if ce.state.blockRes == nil {
			// still processing the block, ignore the commit message for now and commit when ready.
			return false
		} else {
			// Waiting for the block to be committed, notify the consensus engine to commit the block.
			blkCommit.blk = ce.state.blkProp.blk
			ce.sendConsensusMessage(&consensusMessage{
				MsgType: blkCommit.Type(),
				Msg:     blkCommit,
				Sender:  ce.pubKey,
			})
		}
	} else {
		// either sentry node or slow validator
		// check if this is the first time we are hearing about this block and not already downloaded it.
		blk, _, err := ce.blockStore.Get(blkID)
		if err != nil {
			return true
		}

		blkCommit.blk = blk
		ce.sendConsensusMessage(&consensusMessage{
			MsgType: blkCommit.Type(),
			Msg:     blkCommit,
			Sender:  ce.pubKey,
		})
	}

	return false
}

// TODO: Can we club this and AcceptCommit into a single method?
// NotifyBlockCommit is used by the p2p stream handler to notify the consensus engine of a new block commit.
// It validates blk height, appHash and blkHash and only then notifies the consensus engine to commit the block.
func (ce *ConsensusEngine) NotifyBlockCommit(blk *types.Block, appHash types.Hash) {
	// ce.log.Infoln("Notify block commit", blk.Header.Height, blk.Header.Hash(), appHash)
	if ce.role == types.RoleLeader {
		// Leader can also use this in blocksync mode, when it tries to replay the blocks or catchup with the network.
		return
	}

	ce.log.Info("Notify block commit", "height", blk.Header.Height, "hash", blk.Header.Hash(), "appHash", appHash)

	ce.state.mtx.Lock()
	ce.state.mtx.Unlock()

	if ce.state.lc.height+1 != blk.Header.Height {
		return
	}

	if ce.state.blkProp != nil && ce.state.blockRes == nil {
		// ensure that we are processing correct block
		if ce.state.blkProp.blkHash != blk.Header.Hash() {
			// Rollback and reprocess the block sent as part of the commit message.
			msg := &resetState{height: blk.Header.Height - 1}
			ce.log.Info("Processing incorrect block, notify consensus engine to abort: ", "height", blk.Header.Height-1)
			ce.sendConsensusMessage(&consensusMessage{
				MsgType: msg.Type(),
				Msg:     msg,
			})
			// sending the correct block to the consensus engine below.
		} else {
			ce.log.Infof("still processing the block, ignore the commit message for now and commit when ready")
		}
		return
	}

	// Notify the consensus engine to commit the block in the below scenarios:
	// 1. Sentry node: Execute the block, validate the appHash and commit the block.
	// 2. Validator:
	// - No blockProposal received: Execute the block, validate the appHash and commit the block.
	// - Incorrect Block received: Rollback and reprocess the block sent as part of the commit message.
	// - Incorrect AppHash: Halt the node.
	blkCommit := &blockAnnounce{
		blk:     blk,
		appHash: appHash,
	}

	ce.sendConsensusMessage(&consensusMessage{
		MsgType: blkCommit.Type(),
		Msg:     blkCommit,
		Sender:  ce.pubKey,
	})
	// ce.log.Infoln("Notified consensus engine to commit the block", blk.Header.Height, blk.Header.Hash(), appHash)
}

// Accept ResetState message from the leader in the following scenarios:
// 1. If we are currently processing block at height +1 and the leader wants to abort the block processing of height +1.
// 2. If the block at height+1 is not already committed, else its a stale.
func (ce *ConsensusEngine) NotifyResetState(height int64) {
	ce.state.mtx.Lock()
	defer ce.state.mtx.Unlock()

	ce.log.Info("Notify reset state? ", "height", height)
	// Leader is the sender of the reset message, only sentry and validators should accept this message.
	if ce.role != types.RoleValidator {
		return
	}

	// Potentially stale reset message or in the future, ignore it.
	if ce.state.lc.height+1 != height {
		ce.log.Info("Stale reset message, ignoring it.", "height", height, "current", ce.state.lc.height+1)
		return
	}

	// Ensure that you are processing a block at height +1
	if ce.state.blkProp == nil {
		ce.log.Info("Not processing any block proposal at the moment, reset message ignored.")
		return // nothing to reset
	}

	msg := &resetState{height: height - 1}
	ce.sendConsensusMessage(&consensusMessage{
		MsgType: msg.Type(),
		Msg:     msg,
	})
}

// ProcessBlockProposal is used by the validator's consensus engine to process the new block proposal message.
// This method is used to validate the received block, execute the block and generate appHash and
// report the result back to the leader.
func (ce *ConsensusEngine) processBlockProposal(_ context.Context, blkPropMsg *blockProposal) error {
	ce.log.Info("Processing block proposal", "height", blkPropMsg.blk.Header.Height, "header", blkPropMsg.blk.Header)
	if ce.role != types.RoleValidator {
		ce.log.Warn("Only validators can process block proposals")
		return nil
	}

	ce.state.mtx.Lock()
	defer ce.state.mtx.Unlock()

	if ce.state.blkProp != nil {
		return fmt.Errorf("We are in the process of executing a block, can't accept a new block proposal.")
	}

	if err := ce.validateBlock(blkPropMsg.blk); err != nil {
		go ce.ackBroadcaster(false, blkPropMsg.height, blkPropMsg.blkHash, nil)
		ce.log.Error("Error validating block, sending NACK", "error", err)
		return err
	}
	ce.state.blkProp = blkPropMsg

	if err := ce.executeBlock(); err != nil {
		// TODO: what to do if the block execution fails? Send NACK?
		ce.log.Error("Error executing block", "error", err)
		return err
	}

	// Broadcast the result back to the leader
	ce.log.Info("Sending ack to the leader", "height", blkPropMsg.height,
		"hash", blkPropMsg.blkHash, "appHash", ce.state.blockRes.appHash)
	go ce.ackBroadcaster(true, blkPropMsg.height, blkPropMsg.blkHash, &ce.state.blockRes.appHash)

	return nil
}

// This is triggered in response to the blockAnn message from the leader.
// This method is used by the sentry and the validators nodes to commit the specified block.
// If the validator node processed a different block, it should rollback and reprocess the block.
// Validator nodes can skip the block execution and directly commit the block if they have already processed the block.
// The nodes should only commit the block if the appHash is valid, else halt the node.
func (ce *ConsensusEngine) commitBlock(blk *types.Block, appHash types.Hash) error {
	// ce.log.Infoln("processing Commit block", blk.Header.Height, blk.Header.Hash(), appHash)
	if ce.role == types.RoleLeader {
		return nil
	}

	ce.state.mtx.Lock()
	defer ce.state.mtx.Unlock()

	// Three different scenarios are possible here:
	// 1. Sentry node: Execute the block, validate the appHash and commit the block.
	// 2. Validator:
	// - No blockProposal received: Execute the block, validate the appHash and commit the block.
	// - Incorrect Block received: Rollback and reprocess the block sent as part of the commit message.
	// - Incorrect AppHash: Halt the node.

	if ce.role == types.RoleSentry {
		// go ce.processAndCommit(blk, appHash)
		return ce.processAndCommit(blk, appHash)
	}

	// You are a validator
	if ce.state.blkProp == nil {
		// No block proposal received, execute the block, validate the appHash and commit the block.
		// go ce.processAndCommit(blk, appHash)
		return ce.processAndCommit(blk, appHash)
	}

	// ensure that you are processing the correct block
	if ce.state.blkProp.blkHash != blk.Header.Hash() {
		// Rollback and reprocess the block sent as part of the commit message.
		ce.resetState()
		// TODO: somehow signal the current block processing to halt and reprocess the new block.
		return ce.processAndCommit(blk, appHash)
	}

	if ce.state.blockRes == nil {
		// Still processing the block, return and commit when ready.
		return nil
	}

	if ce.state.blockRes.appHash != appHash {
		ce.log.Error("Incorrect AppHash, halting the node.", "got", appHash, "expected", ce.state.blockRes.appHash)
		close(ce.haltChan)
		return nil
	}

	// Commit the block
	if err := ce.commit(); err != nil {
		ce.log.Errorf("Error committing block: %v", err)
		return err
	}

	// Move to the next state
	ce.nextState()
	return nil
}

// processAndCommit: used by the sentry nodes and slow validators to process and commit the block.
// This is used when the acks are not required to be sent back to the leader, essentially in catchup mode.
func (ce *ConsensusEngine) processAndCommit(blk *types.Block, appHash types.Hash) error {
	ce.log.Info("Processing committed block", "height", blk.Header.Height, "hash", blk.Header.Hash(), "appHash", appHash)
	if err := ce.validateBlock(blk); err != nil {
		ce.log.Errorf("Error validating block: %v", err)
		return err // who gets this error?
	}
	ce.state.blkProp = &blockProposal{
		height:  blk.Header.Height,
		blkHash: blk.Header.Hash(),
		blk:     blk,
		// respCb is not required here as we are not sending acks back to the leader.
	}

	if err := ce.executeBlock(); err != nil {
		ce.log.Errorf("Error executing block: %v", err)
		return err
	}

	if ce.state.blockRes.appHash != appHash {
		// Incorrect AppHash, halt the node.
		ce.log.Error("processAndCommit: Incorrect AppHash", "received", appHash, "have", ce.state.blockRes.appHash)
		// ce.haltChan <- struct{}{}
		close(ce.haltChan)
		return fmt.Errorf("appHash mismatch, expected: %s, received: %s", appHash, ce.state.blockRes.appHash)
	}

	// Commit the block if the appHash is valid
	if err := ce.commit(); err != nil {
		ce.log.Errorf("Error committing block: %v", err)
		return err
	}

	// Move to the next state
	ce.nextState()
	return nil
}