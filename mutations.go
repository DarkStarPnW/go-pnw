package pnw

import "context"

// BankWithdrawInput holds the parameters for a bank withdrawal mutation.
// This mutation is only available to verified bots.
type BankWithdrawInput struct {
	ReceiverID   int
	ReceiverType int // 1 = nation, 2 = alliance
	Note         string
	Money        float64
	Coal         float64
	Oil          float64
	Uranium      float64
	Iron         float64
	Bauxite      float64
	Lead         float64
	Gasoline     float64
	Munitions    float64
	Steel        float64
	Aluminum     float64
	Food         float64
}

// BankDepositInput holds the parameters for a bank deposit mutation.
// This mutation is only available to verified bots.
type BankDepositInput struct {
	Note      string
	Money     float64
	Coal      float64
	Oil       float64
	Uranium   float64
	Iron      float64
	Bauxite   float64
	Lead      float64
	Gasoline  float64
	Munitions float64
	Steel     float64
	Aluminum  float64
	Food      float64
}

// BankWithdraw withdraws resources from the alliance bank to a nation or alliance.
// Requires a verified bot API key.
func (c *Client) BankWithdraw(ctx context.Context, input BankWithdrawInput) (*BankRec, error) {
	q := `mutation BankWithdraw(
		$receiver: Int!, $receiver_type: Int!, $note: String,
		$money: Float, $coal: Float, $oil: Float, $uranium: Float,
		$iron: Float, $bauxite: Float, $lead: Float, $gasoline: Float,
		$munitions: Float, $steel: Float, $aluminum: Float, $food: Float
	) {
		bankWithdraw(
			receiver: $receiver, receiver_type: $receiver_type, note: $note,
			money: $money, coal: $coal, oil: $oil, uranium: $uranium,
			iron: $iron, bauxite: $bauxite, lead: $lead, gasoline: $gasoline,
			munitions: $munitions, steel: $steel, aluminum: $aluminum, food: $food
		) {
			id date sid stype rid rtype note
			money coal oil uranium iron bauxite lead
			gasoline munitions steel aluminum food
		}
	}`

	vars := map[string]any{
		"receiver":      input.ReceiverID,
		"receiver_type": input.ReceiverType,
		"note":          input.Note,
		"money":         input.Money,
		"coal":          input.Coal,
		"oil":           input.Oil,
		"uranium":       input.Uranium,
		"iron":          input.Iron,
		"bauxite":       input.Bauxite,
		"lead":          input.Lead,
		"gasoline":      input.Gasoline,
		"munitions":     input.Munitions,
		"steel":         input.Steel,
		"aluminum":      input.Aluminum,
		"food":          input.Food,
	}

	var out struct {
		BankWithdraw BankRec `json:"bankWithdraw"`
	}
	if err := c.do(ctx, q, vars, &out); err != nil {
		return nil, err
	}
	return &out.BankWithdraw, nil
}

// BankDeposit deposits resources into the alliance bank from the authenticated nation.
// Requires a verified bot API key.
func (c *Client) BankDeposit(ctx context.Context, input BankDepositInput) (*BankRec, error) {
	q := `mutation BankDeposit(
		$note: String,
		$money: Float, $coal: Float, $oil: Float, $uranium: Float,
		$iron: Float, $bauxite: Float, $lead: Float, $gasoline: Float,
		$munitions: Float, $steel: Float, $aluminum: Float, $food: Float
	) {
		bankDeposit(
			note: $note,
			money: $money, coal: $coal, oil: $oil, uranium: $uranium,
			iron: $iron, bauxite: $bauxite, lead: $lead, gasoline: $gasoline,
			munitions: $munitions, steel: $steel, aluminum: $aluminum, food: $food
		) {
			id date sid stype rid rtype note
			money coal oil uranium iron bauxite lead
			gasoline munitions steel aluminum food
		}
	}`

	vars := map[string]any{
		"note":      input.Note,
		"money":     input.Money,
		"coal":      input.Coal,
		"oil":       input.Oil,
		"uranium":   input.Uranium,
		"iron":      input.Iron,
		"bauxite":   input.Bauxite,
		"lead":      input.Lead,
		"gasoline":  input.Gasoline,
		"munitions": input.Munitions,
		"steel":     input.Steel,
		"aluminum":  input.Aluminum,
		"food":      input.Food,
	}

	var out struct {
		BankDeposit BankRec `json:"bankDeposit"`
	}
	if err := c.do(ctx, q, vars, &out); err != nil {
		return nil, err
	}
	return &out.BankDeposit, nil
}
