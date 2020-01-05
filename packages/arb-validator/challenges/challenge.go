/*
* Copyright 2020, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package challenges

import (
	"context"
	"errors"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type ChallengeState uint8

const (
	ChallengeContinuing ChallengeState = iota
	ChallengeAsserterWon
	ChallengeAsserterTimedOut
	ChallengeChallengerTimedOut
)

var challengeNoEvents = errors.New("PendingTopChallengeContract notification channel terminated unexpectedly")

func handleNextEvent(note ethbridge.Notification) (outNote ethbridge.Notification, state ChallengeState, err error) {
	switch note.Event.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return note, ChallengeAsserterTimedOut, nil
	case ethbridge.ChallengerTimeoutEvent:
		return note, ChallengeChallengerTimedOut, nil
	}
	return note, 0, nil
}

func getNextEvent(outChan chan ethbridge.Notification) (note ethbridge.Notification, state ChallengeState, err error) {
	note, ok := <-outChan
	if !ok {
		return note, 0, challengeNoEvents
	}
	return handleNextEvent(note)
}

func getNextEventWithTimeout(
	ctx context.Context,
	outChan chan ethbridge.Notification,
	deadline structures.TimeTicks,
	contract ethbridge.ChallengeContract,
) (note ethbridge.Notification, state ChallengeState, err error) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return note, state, errors.New("context cancelled while waiting for event")
		case <-ticker.C:
			currentTime, err := contract.CurrentBlockTime(ctx)
			if err != nil {
				return note, 0, err
			}
			if structures.TimeFromBlockNum(currentTime).Cmp(deadline) >= 0 {
				_, err := contract.TimeoutChallenge(ctx)
				if err != nil {
					return note, 0, err
				}
				ticker.Stop()
			}
		case note, ok := <-outChan:
			if !ok {
				return note, 0, challengeNoEvents
			}
			return handleNextEvent(note)
		}
	}
}