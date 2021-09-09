# mygotool
自定义工具类

# use
tickers.SetupTickManager(
	examples.GetTickerConfig(),//这里写上你要执行配置
)

defer tickers.TickManager().Stop() //停止
