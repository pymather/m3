	"github.com/m3db/m3db/x/xio"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"
	if maxLRUBlocks := cfg.Cache.SeriesConfiguration().MaxLRUBlocks; maxLRUBlocks > 0 {
		runtimeOpts = runtimeOpts.SetMaxWiredBlocks(maxLRUBlocks)
	}
	segmentReaderPool := xio.NewSegmentReaderPool(
	contextPool := context.NewPool(context.NewOptions().
		SetContextPoolOptions(contextPoolOpts).
		SetFinalizerPoolOptions(closersPoolOpts))
	var identifierPool ident.Pool
		identifierPool = ident.NewPool(
		identifierPool = ident.NewNativePool(

	if opts.SeriesCachePolicy() == series.CacheLRU {
		runtimeOpts := opts.RuntimeOptionsManager()
		wiredList := block.NewWiredList(runtimeOpts, iopts, opts.ClockOptions())
		blockOpts = blockOpts.SetWiredList(wiredList)
	}