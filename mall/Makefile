api[after,order,product]:
	cd apps/app/api && go run api.go

order:
	cd apps/order/rpc && go run order.go

product:
	cd apps/product/rpc && go run product.go

test[after,all]:
	curl http://127.0.0.1:8888/v1/order/list?uid=123