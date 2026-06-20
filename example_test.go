package pnw_test

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"

	pnw "github.com/Jonhasacat/go-pnw"
)

// apiKey loads PNW_API_KEY from .env or the environment.
// Tests are skipped when the key is absent so the suite passes in CI without credentials.
func apiKey(t *testing.T) string {
	t.Helper()
	_ = godotenv.Load()
	key := os.Getenv("PNW_API_KEY")
	if key == "" {
		t.Skip("PNW_API_KEY not set — skipping integration test")
	}
	return key
}

func TestAlliances(t *testing.T) {
	client := pnw.NewClient(apiKey(t))

	result, err := client.Alliances(context.Background(), pnw.AllianceFilter{First: 5})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Data) == 0 {
		t.Fatal("expected at least one alliance")
	}
	a := result.Data[0]
	t.Logf("Alliance: %s (%s) score=%.2f", a.Name, a.Acronym, a.Score)
}

func TestNations(t *testing.T) {
	client := pnw.NewClient(apiKey(t))

	minCities := 10
	result, err := client.Nations(context.Background(), pnw.NationFilter{
		MinCities: &minCities,
		First:     10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("nations=%d page=1/%d", len(result.Data), result.PaginatorInfo.LastPage)
}

func TestAllNations(t *testing.T) {
	client := pnw.NewClient(apiKey(t))

	var total int
	err := client.AllNations(context.Background(), pnw.NationFilter{First: 500},
		func(page []pnw.Nation) bool {
			total += len(page)
			return total < 1000 // stop after 1000 to keep the test fast
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total nations fetched: %d", total)
}

func TestWars(t *testing.T) {
	client := pnw.NewClient(apiKey(t))

	daysAgo := 7
	result, err := client.Wars(context.Background(), pnw.WarFilter{
		DaysAgo: &daysAgo,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("wars in last 7 days: %d", len(result.Data))
}

func TestTradeprices(t *testing.T) {
	client := pnw.NewClient(apiKey(t))

	result, err := client.Tradeprices(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Data) == 0 {
		t.Fatal("expected trade prices")
	}
	p := result.Data[0]
	t.Logf("food=%.0f steel=%.0f aluminum=%.0f", p.Food, p.Steel, p.Aluminum)
}

func TestRawQuery(t *testing.T) {
	client := pnw.NewClient(apiKey(t))

	var result struct {
		Nations struct {
			Data []struct {
				ID         pnw.ID  `json:"id"`
				NationName string  `json:"nation_name"`
				Score      float64 `json:"score"`
			} `json:"data"`
		} `json:"nations"`
	}

	err := client.Query(context.Background(),
		`query { nations(first: 5) { data { id nation_name score } } }`,
		nil,
		&result,
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range result.Nations.Data {
		t.Logf("%d  %s  %.2f", n.ID, n.NationName, n.Score)
	}
}
