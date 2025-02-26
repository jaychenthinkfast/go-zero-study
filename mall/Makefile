#svc
api[after,order,product,reply]:
	cd apps/app/api && go run api.go

admin_api_rest:
	cd apps/product/admin && go run admin.go

order:
	cd apps/order/rpc && go run order.go

product:
	cd apps/product/rpc && go run product.go

reply:
	cd apps/reply/rpc && go run reply.go

#test
redis_flushall:
	redis-cli flushall

test_order:
	curl http://127.0.0.1:8888/v1/order/list?uid=123

test_productdetail:
	curl http://127.0.0.1:8888/v1/product/detail?product_id=1

test_imageupload:
	curl --location 'http://127.0.0.1:8889/v1/upload/image' \
    --form 'image=@"/path/test.png"'

test_Product:
	grpcurl -plaintext -d '{"product_id": 1}' 127.0.0.1:8081 product.Product.Product

test_ProductList:
	grpcurl -plaintext -d '{"category_id": 8}' 127.0.0.1:8081 product.Product.ProductList

test_OperationProducts:
	grpcurl -plaintext -d '{}' 127.0.0.1:8081 product.Product.OperationProducts
#code generation
rpc[reply]:
	cd apps/reply/rpc && goctl rpc protoc reply.proto --go_out=. --go-grpc_out=. --zrpc_out=.

rpc[product]:
	cd apps/product/rpc && goctl rpc protoc product.proto --go_out=. --go-grpc_out=. --zrpc_out=.

rpc[order]:
	cd apps/order/rpc && goctl rpc protoc order.proto --go_out=. --go-grpc_out=. --zrpc_out=.

api[app]:
	cd apps/app/api && goctl api go -api api.api -dir .

admin_api[product]:
	cd apps/product/admin && goctl api go -api admin.api -dir .

model[product]:
	goctl model mysql datasource -url="user:pass@tcp(127.0.0.1:3307)/product" --dir="./apps/product/rpc/internal/model" -cache=true -table="*"