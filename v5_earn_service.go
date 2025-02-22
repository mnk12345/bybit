package bybit

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
)

// V5EarnServiceI :
type V5EarnServiceI interface {
	GetProductInfo(V5GetProductInfoParam) (*V5GetProductInfoResponse, error)
	GetStakedPositions(V5GetStakedPositionsParam) (*V5GetStakedPositionsResponse, error)
}

// V5EarnService :
type V5EarnService struct {
	client *Client
}

// V5GetStakedPositionsParam :
type V5GetStakedPositionsParam struct {
	Category  CategoryV5 `url:"category"`
	ProductID *string    `json:"productId,omitempty"`
	Coin      Coin
}

func (p V5GetStakedPositionsParam) validate() error {
	if p.Category == "" {
		return fmt.Errorf("category needed")
	}
	if p.Category != CategoryV5FlexibleSaving {
		return fmt.Errorf("category must be flexible savings")
	}
	return nil
}

// GetStakedPositions :
func (s *V5EarnService) GetStakedPositions(param V5GetStakedPositionsParam) (*V5GetStakedPositionsResponse, error) {
	var res V5GetStakedPositionsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/earn/position", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetStakedPositionsResponse :
type V5GetStakedPositionsResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetStakedPositionsResult `json:"result"`
}

// V5GetStakedPositionsResult :
type V5GetStakedPositionsResult struct {
	List V5GetStakedPositionsList `json:"list"`
}

// V5GetStakedPositionsList :
type V5GetStakedPositionsList []V5GetStakedPositionsItem

// V5GetStakedPositionsItem :
type V5GetStakedPositionsItem struct {
	Coin           Coin   `json:"coin"`
	ProductId      string `json:"productId"`
	Amount         string `json:"amount"`
	TotalPnl       string `json:"totalPnl"`
	ClaimableYield string `json:"claimableYield"`
}

// V5GetProductInfoParam :
type V5GetProductInfoParam struct {
	Category CategoryV5 `url:"category"`
	Coin     Coin
}

func (p V5GetProductInfoParam) validate() error {
	if p.Category == "" {
		return fmt.Errorf("category needed")
	}
	if p.Category != CategoryV5FlexibleSaving {
		return fmt.Errorf("category must be flexible savings")
	}
	return nil
}

// GetProductInfo :
func (s *V5EarnService) GetProductInfo(param V5GetProductInfoParam) (*V5GetProductInfoResponse, error) {
	var res V5GetProductInfoResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/earn/product", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetProductInfoResponse :
type V5GetProductInfoResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetProductInfoResult `json:"result"`
}

// V5GetProductInfoResult :
type V5GetProductInfoResult struct {
	List V5GetProductInfoList `json:"list"`
}

// V5GetProductInfoList :
type V5GetProductInfoList []V5GetProductInfoItem

// V5GetProductInfoItem :
type V5GetProductInfoItem struct {
	Coin           Coin       `json:"coin"`
	ProductId      string     `json:"productId"`
	MinStakeAmount string     `json:"minStakeAmount"`
	MaxStakeAmount string     `json:"maxStakeAmount"`
	EstimateApr    string     `json:"estimateApr"`
	Precision      string     `json:"precision"`
	Status         string     `json:"status"`
	Category       CategoryV5 `json:"category"`
}

// V5GetProductInfoParam :
type V5StakeRedeemParam struct {
	Category    CategoryV5 `url:"category"`
	Coin        Coin
	OrderType   OrderTypeV5
	AccountType AccountTypeV5
	Amount      string
	ProductId   string
	OrderLinkId string
}

// StakeRedeem :
func (s *V5EarnService) StakeRedeem(param V5StakeRedeemParam) (*V5StakeRedeemResponse, error) {
	var res V5StakeRedeemResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/earn/place-order", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

func (p V5StakeRedeemParam) validate() error {
	if p.Category == "" {
		return fmt.Errorf("category needed")
	}
	if p.OrderType == "" {
		return fmt.Errorf("order type needed")
	}
	if p.AccountType == "" {
		return fmt.Errorf("account type needed")
	}
	if p.Amount == "" {
		return fmt.Errorf("amount type needed")
	}
	if p.ProductId == "" {
		return fmt.Errorf("product id type needed")
	}
	if p.OrderLinkId == "" {
		return fmt.Errorf("order link id needed")
	}
	if p.Category != CategoryV5FlexibleSaving {
		return fmt.Errorf("category must be flexible savings")
	}
	if p.OrderType != OrderTypeV5Stake && p.OrderType != OrderTypeV5Redeem {
		return fmt.Errorf("order type should be either stake or redeem")
	}
	if p.AccountType != AccountTypeV5UNIFIED && p.AccountType != AccountTypeV5FUND {
		return fmt.Errorf("account type should be either fund or unified")
	}
	return nil
}

// V5StakeRedeemResponse :
type V5StakeRedeemResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5StakeRedeemResult `json:"result"`
}

// V5StakeRedeemResult :
type V5StakeRedeemResult struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}
