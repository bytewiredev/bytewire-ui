package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/bytewiredev/bytewire/pkg/engine"
	"github.com/bytewiredev/bytewire/pkg/router"
	"github.com/bytewiredev/bytewire/pkg/style"

	ui "github.com/bytewiredev/bytewire-ui"
	"github.com/bytewiredev/bytewire-ui/showcase/pages"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	r := router.New().
		Handle("/", pages.Home).
		Handle("/buttons", pages.Buttons).
		Handle("/forms", pages.Forms).
		Handle("/data", pages.Data).
		Handle("/feedback", pages.Feedback).
		Handle("/navigation", pages.Navigation).
		Handle("/typography", pages.Typography).
		Handle("/layout", pages.Layout)

	srv := engine.NewServer(":4553", nil, r.Mount,
		engine.WithWebSocketFallback(),
		engine.WithSSR(),
		engine.WithStaticDir("static"),
		engine.WithCSS(ui.BaseCSS+ui.SpinnerCSS+ui.SkeletonCSS+style.GenerateAll()),
		engine.WithHTTPAddr(":9091"),
	)

	log.Println("Showcase running on http://localhost:9091")
	if err := srv.ListenAndServe(ctx); err != nil {
		log.Fatal(err)
	}
}
