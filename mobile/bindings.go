package lndmobile

import (
	"fmt"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/lightningnetwork/lnd"
	"google.golang.org/grpc/test/bufconn"
)


var (
	bufLightningLis		= bufconn.Listen(100)
	bufWalletUnlockerLis 	= bufconn.Listen(100)
)

func Start(extraArgs string, callback Callback) {
	// Add the extra arguments to os.Args, as that will be parsed during
	// startup
	os.Args = append(os.Args, strings.Fields(extraArgs)...)

	// Call the "real" main in a nested manner so the defers will properly
	// be exectured in the case of a graceful shutdown
	go func() {
		if err := lnd.Main(bufLightningLis, bufWalletUnlockerLis); err != nil {
		//if err := lnd.Main(); err != nil { /* Was throwing error about too many arguments for lnd.Main()*/
			if e, ok := err.(*flags.Error); ok &&
				e.Type == flags.ErrHelp {
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(1)
		}
	}()

	// TODO(halseth): callback when RPC server is running.
	callback.OnResponse([]byte("started"))
}
