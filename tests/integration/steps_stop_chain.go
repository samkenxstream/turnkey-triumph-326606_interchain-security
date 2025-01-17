package main

import "time"

// submits a consumer-removal proposal and removes the chain
func stepsStopChain(consumerName string, propNumber uint) []Step {
	s := []Step{
		{
			action: submitConsumerRemovalProposalAction{
				chain:          chainID("provi"),
				from:           validatorID("bob"),
				deposit:        10000001,
				consumerChain:  chainID(consumerName),
				stopTimeOffset: 0 * time.Millisecond,
			},
			state: State{
				chainID("provi"): ChainState{
					ValBalances: &map[validatorID]uint{
						validatorID("bob"): 9489999999,
					},
					Proposals: &map[uint]Proposal{
						propNumber: ConsumerRemovalProposal{
							Deposit:  10000001,
							Chain:    chainID(consumerName),
							StopTime: 0,
							Status:   "PROPOSAL_STATUS_VOTING_PERIOD",
						},
					},
					ConsumerChains: &map[chainID]bool{"consu": true}, // consumer chain not yet removed
				},
			},
		},
		{
			action: voteGovProposalAction{
				chain:      chainID("provi"),
				from:       []validatorID{validatorID("alice"), validatorID("bob"), validatorID("carol")},
				vote:       []string{"yes", "yes", "yes"},
				propNumber: propNumber,
			},
			state: State{
				chainID("provi"): ChainState{
					Proposals: &map[uint]Proposal{
						propNumber: ConsumerRemovalProposal{
							Deposit:  10000001,
							Chain:    chainID(consumerName),
							StopTime: 0,
							Status:   "PROPOSAL_STATUS_PASSED",
						},
					},
					ValBalances: &map[validatorID]uint{
						validatorID("bob"): 9500000000,
					},
					ConsumerChains: &map[chainID]bool{}, // Consumer chain is now removed
				},
			},
		},
	}

	return s
}
