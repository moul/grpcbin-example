run:
	grpcc -i -p ./grpcbin.proto --address grpcb.in:9000 --eval 'client.dummyUnary({f_string:"hello world",f_int32:42}, printReply)'

install:
	wget -qN https://github.com/moul/pb/raw/master/grpcbin/grpcbin.proto
	npm install -g grpcc
