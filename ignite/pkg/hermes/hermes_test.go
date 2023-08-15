package hermes_test

import (
	"context"
	"testing"

	"github.com/ignite/cli/ignite/pkg/hermes"
)

func TestHermes(t *testing.T) {
	ctx := context.Background()
	h, err := hermes.New()
	if err != nil {
		t.Fatal(err)
	}
	defer h.Cleanup()

	// Create the default config and add chains
	c := hermes.DefaultConfig()
	err = c.AddChain("mars-1", "http://0.0.0.0:26649", "http://localhost:9082")
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddChain("venus-1", "http://0.0.0.0:26659", "http://localhost:9092")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Save(); err != nil {
		t.Fatal(err)
	}

	// Add hermes keys
	err = h.AddMnemonic(
		ctx,
		"mars-1",
		"letter column benefit acoustic evidence false trim cave jump pluck awesome lion",
	)
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddMnemonic(
		ctx,
		"venus-1",
		"jeans payment lock client result enemy bullet rug crush deny month salad",
	)
	if err != nil {
		t.Fatal(err)
	}

	// create clients
	err = h.CreateClient(ctx, "mars-1", "venus-1")
	if err != nil {
		t.Fatal(err)
	}
	err = h.CreateClient(ctx, "venus-1", "mars-1")
	if err != nil {
		t.Fatal(err)
	}

	// create connection
	err = h.CreateConnection(ctx, "mars-1", "07-tendermint-0", "07-tendermint-0")
	if err != nil {
		t.Fatal(err)
	}

	// create and query channel
	err = h.CreateChannel(ctx, "mars-1", "connection-0", "transfer", "transfer")
	if err != nil {
		t.Fatal(err)
	}
	err = h.QueryChannels(ctx, true, "mars-1")
	if err != nil {
		t.Fatal(err)
	}

	// start hermes
	err = h.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
