# stock-exporter

## Overview

This project is a proof of concept. `stock-exporter` scrapes data periodically from Yahoo Finance (YFinance) and exports this data in Prometheus metrics format at the /metrics endpoint. This project also publishes an OCI image and Helm chart, designed to be deployed to a Kubernetes cluster.

```sh
$ stocker-exporter 
{"level":"info","ts":1723921832.199917,"caller":"exporter/handler.go:24","msg":"StockPriceExporter listening on 0.0.0.0:9090"}
{"level":"info","ts":1723921832.200129,"caller":"exporter/data.go:44","msg":"Updating stock prices for 7 symbols"}
{"level":"info","ts":1723921832.484951,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"AMZN","price":177.06}
{"level":"info","ts":1723921832.6025732,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"JEPI","price":57.61}
{"level":"info","ts":1723921832.767125,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"JEPQ","price":53.45}
{"level":"info","ts":1723921832.872498,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"MSFT","price":418.47}
{"level":"info","ts":1723921832.944281,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"NVDA","price":124.58}
{"level":"info","ts":1723921833.01983,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"TQQQ","price":68.64}
{"level":"info","ts":1723921833.0904331,"caller":"exporter/data.go:71","msg":"Current price","Symbol":"TSLA","price":216.12}
```
