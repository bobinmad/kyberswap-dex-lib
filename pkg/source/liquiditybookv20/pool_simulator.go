package liquiditybookv20

import (
	"math/big"
	"sort"
	"strings"

	"github.com/KyberNetwork/blockchain-toolkit/integer"
	"github.com/KyberNetwork/logger"
	"github.com/goccy/go-json"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
)

type PoolSimulator struct {
	pool.Pool

	blockTimestamp uint64
	feeParams      feeParameters
	activeBinID    uint32
	bins           []Bin
}

func NewPoolSimulator(entityPool entity.Pool) (*PoolSimulator, error) {
	var (
		tokens   = make([]string, 2)
		reserves = make([]*big.Int, 2)

		extra Extra
	)

	if len(entityPool.Reserves) == 2 && len(entityPool.Tokens) == 2 {
		tokens[0] = entityPool.Tokens[0].Address
		reserves[0] = bignumber.NewBig10(entityPool.Reserves[0])

		tokens[1] = entityPool.Tokens[1].Address
		reserves[1] = bignumber.NewBig10(entityPool.Reserves[1])
	}

	err := json.Unmarshal([]byte(entityPool.Extra), &extra)
	if err != nil {
		return nil, err
	}

	info := pool.PoolInfo{
		Address:    strings.ToLower(entityPool.Address),
		ReserveUsd: entityPool.ReserveUsd,
		SwapFee:    nil,
		Exchange:   entityPool.Exchange,
		Type:       entityPool.Type,
		Tokens:     tokens,
		Reserves:   reserves,
		Checked:    false,
	}

	return &PoolSimulator{
		Pool:           pool.Pool{Info: info},
		blockTimestamp: extra.RpcBlockTimestamp,
		feeParams:      extra.FeeParameters,
		activeBinID:    extra.ActiveBinID,
		bins:           extra.Bins,
	}, nil
}

func (p *PoolSimulator) CalcAmountOut(param pool.CalcAmountOutParams) (*pool.CalcAmountOutResult, error) {
	tokenAmountIn := param.TokenAmountIn
	tokenOut := param.TokenOut
	err := p.validateTokens([]string{tokenAmountIn.Token, tokenOut})
	if err != nil {
		return nil, err
	}
	amountIn := tokenAmountIn.Amount
	swapForY := tokenAmountIn.Token == p.Info.Tokens[0]

	swapOutResult, err := p.getSwapOut(amountIn, swapForY)
	if err != nil {
		return nil, err
	}

	return &pool.CalcAmountOutResult{
		TokenAmountOut: &pool.TokenAmount{
			Token:  tokenOut,
			Amount: swapOutResult.AmountOut,
		},
		Fee: &pool.TokenAmount{
			Token:  tokenAmountIn.Token,
			Amount: swapOutResult.Fee,
		},
		Gas: defaultGas,
		SwapInfo: SwapInfo{
			BinsReserveChanges: swapOutResult.BinsReserveChanges,
			NewFeeParameters:   swapOutResult.FeeParameters,
			NewActiveID:        swapOutResult.NewActiveID,
		},
	}, nil
}

func (p *PoolSimulator) getSwapOut(amountIn *big.Int, swapForY bool) (*getSwapOutResult, error) {
	var (
		id                 = p.activeBinID
		amountOut          = integer.Zero()
		swapFee            = integer.Zero()
		binsReserveChanges []binReserveChanges
	)

	// All fields are value type, so we can copy directly.
	fp := p.feeParams
	fp.updateVariableFeeParameters(p.blockTimestamp, id)

	for {
		binArrIdx, err := p.findBinArrIndex(id)
		if err != nil {
			return nil, err
		}
		bin := p.bins[binArrIdx]
		if !bin.isEmptyForSwap(!swapForY) {
			amountInToBin, amountOutOfBin, totalFee, _, err := bin.getAmounts(&fp, id, swapForY, amountIn)
			if err != nil {
				return nil, err
			}

			swapFee = new(big.Int).Add(swapFee, totalFee)

			amountIn = new(big.Int).Sub(amountIn, new(big.Int).Add(amountInToBin, totalFee))
			amountOut = new(big.Int).Add(amountOut, amountOutOfBin)

			newBinReserveChanges := newBinReserveChanges(
				id, !swapForY, amountInToBin, amountOutOfBin,
			)
			binsReserveChanges = append(binsReserveChanges, newBinReserveChanges)
		}

		if amountIn.Cmp(integer.Zero()) == 0 {
			break
		}

		nextID, err := p.getNextNonEmptyBin(swapForY, id)
		if err != nil {
			return nil, err
		}

		id = nextID
	}

	ret := getSwapOutResult{
		AmountOut:          amountOut,
		Fee:                swapFee,
		BinsReserveChanges: binsReserveChanges,
		FeeParameters:      fp,
		NewActiveID:        id,
	}

	return &ret, nil
}

func (p *PoolSimulator) UpdateBalance(params pool.UpdateBalanceParams) {
	swapInfo, ok := params.SwapInfo.(SwapInfo)
	if !ok {
		logger.WithFields(logger.Fields{
			"address": p.Info.Address,
		}).Warn("invalid swap info")
	}

	// update total reserves
	for idx, reserve := range p.Info.Reserves {
		if strings.EqualFold(p.Info.Tokens[idx], params.TokenAmountIn.Token) {
			p.Info.Reserves[idx] = new(big.Int).Add(reserve, params.TokenAmountIn.Amount)
		}
		if strings.EqualFold(p.Info.Tokens[idx], params.TokenAmountOut.Token) {
			p.Info.Reserves[idx] = new(big.Int).Sub(reserve, params.TokenAmountOut.Amount)
		}
	}

	// active bin ID
	p.activeBinID = swapInfo.NewActiveID

	// fee
	p.feeParams = swapInfo.NewFeeParameters

	// update reserves of bins
	totalBinReserveChanges := make(map[uint32]binReserveChanges)
	for _, b := range swapInfo.BinsReserveChanges {
		changes, ok := totalBinReserveChanges[b.BinID]
		if !ok {
			changes = binReserveChanges{
				BinID:      b.BinID,
				AmountXIn:  integer.Zero(),
				AmountXOut: integer.Zero(),
				AmountYIn:  integer.Zero(),
				AmountYOut: integer.Zero(),
			}
		}
		changes.AmountXIn = new(big.Int).Add(changes.AmountXIn, b.AmountXIn)
		changes.AmountXOut = new(big.Int).Add(changes.AmountXOut, b.AmountXOut)
		changes.AmountYIn = new(big.Int).Add(changes.AmountYIn, b.AmountYIn)
		changes.AmountYOut = new(big.Int).Add(changes.AmountYOut, b.AmountYOut)

		totalBinReserveChanges[b.BinID] = changes
	}
	newBins := []Bin{}
	for _, b := range p.bins {
		newBin := Bin{
			ID:       b.ID,
			ReserveX: new(big.Int).Set(b.ReserveX),
			ReserveY: new(big.Int).Set(b.ReserveY),
		}

		changes, ok := totalBinReserveChanges[newBin.ID]
		if ok {
			newBin.ReserveX = new(big.Int).Add(new(big.Int).Sub(newBin.ReserveX, changes.AmountXOut), changes.AmountXIn)
			newBin.ReserveY = new(big.Int).Add(new(big.Int).Sub(newBin.ReserveY, changes.AmountYOut), changes.AmountYIn)
		}

		if !newBin.isEmpty() {
			newBins = append(newBins, newBin)
		}
	}
	sort.Slice(newBins, func(i, j int) bool {
		return newBins[i].ID < newBins[j].ID
	})
	p.bins = newBins
}

func (t *PoolSimulator) GetMetaInfo(_ string, _ string) interface{} {
	return nil
}

func (p *PoolSimulator) validateTokens(tokens []string) error {
	for _, t := range tokens {
		if p.GetTokenIndex(t) < 0 {
			return ErrInvalidToken
		}
	}
	return nil
}

func (p *PoolSimulator) findBinArrIndex(binID uint32) (uint32, error) {
	if len(p.bins) == 0 {
		return 0, ErrNotFoundBinID
	}

	var (
		l = 0
		r = len(p.bins)
	)

	for r-l > 1 {
		m := (r + l) >> 1
		if p.bins[m].ID <= binID {
			l = m
		} else {
			r = m
		}
	}

	if p.bins[l].ID != binID {
		return 0, ErrNotFoundBinID
	}

	return uint32(l), nil
}

func (p *PoolSimulator) getNextNonEmptyBin(swapForY bool, id uint32) (uint32, error) {
	if swapForY {
		return p.findFirstRight(id)
	}

	return p.findFirstLeft(id)
}

func (p *PoolSimulator) findFirstRight(id uint32) (uint32, error) {
	idx, err := p.findBinArrIndex(id)
	if err != nil {
		return 0, err
	}
	if idx == 0 {
		return 0, ErrNotFoundBinID
	}
	return p.bins[idx-1].ID, nil
}

func (p *PoolSimulator) findFirstLeft(id uint32) (uint32, error) {
	idx, err := p.findBinArrIndex(id)
	if err != nil {
		return 0, err
	}
	if idx == uint32(len(p.bins)-1) {
		return 0, ErrNotFoundBinID
	}
	return p.bins[idx+1].ID, nil
}
