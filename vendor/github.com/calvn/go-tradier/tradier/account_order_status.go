package tradier

import "fmt"

func (s *AccountService) OrderStatus(accountId, orderId string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountId, orderId)
	return s.AccountRequest(u)
}
