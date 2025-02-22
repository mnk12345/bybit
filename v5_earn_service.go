package bybit

import (
	"github.com/google/go-querystring/query"
)

// V5EarnServiceI :
type V5EarnServiceI interface {
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
