package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/ethermint/tests"
	"github.com/stretchr/testify/suite"
)

type CSRTestSuite struct {
	suite.Suite
	owner     string
	contracts []string
	id        uint64
	account   string
}

func TestCSRSuite(t *testing.T) {
	suite.Run(t, new(CSRTestSuite))
}

func (suite *CSRTestSuite) SetupTest() {
	suite.owner = sdk.AccAddress(tests.GenerateAddress().Bytes()).String()
	suite.contracts = []string{tests.GenerateAddress().String(), tests.GenerateAddress().String(),
		tests.GenerateAddress().String(), tests.GenerateAddress().String()}
	suite.id = 0
	suite.account = sdk.AccAddress(tests.GenerateAddress().Bytes()).String()
}

func (suite *CSRTestSuite) TestCSR() {
	testCases := []struct {
		msg        string
		csr        CSR
		expectPass bool
	}{
		{
			"Create CSR object - pass",
			CSR{
				Owner:     suite.owner,
				Contracts: suite.contracts,
				Id:        suite.id,
				Account:   suite.account,
			},
			true,
		},
		{
			"Create CSR object with 0 smart contracts - fail",
			CSR{
				Owner:     suite.owner,
				Contracts: []string{},
				Id:        suite.id,
				Account:   suite.account,
			},
			false,
		},
		{
			"Create CSR object with invalid owner address - fail",
			CSR{
				Owner:     "",
				Contracts: suite.contracts,
				Id:        suite.id,
				Account:   suite.account,
			},
			false,
		},
		{
			"Create CSR object with invalid account address - fail",
			CSR{
				Owner:     suite.owner,
				Contracts: suite.contracts,
				Id:        suite.id,
				Account:   "",
			},
			false,
		},
		{
			"Create CSR object with invalid smart contract addresses - fail",
			CSR{
				Owner:     suite.owner,
				Contracts: append(suite.contracts, ""),
				Id:        suite.id,
				Account:   suite.account,
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.msg, func() {
			err := tc.csr.Validate()

			if tc.expectPass {
				suite.Require().NoError(err, tc.msg)
			} else {
				suite.Require().Error(err, tc.msg)
			}
		})
	}
}
