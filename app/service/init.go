package service


func Start() {
	// TODO: To setup grpc server and mq consumer.
	//cfg, _ := ini.Load(resource.ConfFilePath)
	//rpcCfg := cfg.Section("rpc")
	//address := rpcCfg.Key("host").String() + ":" + rpcCfg.Key("port").String()
	//
	//
	//listen, err := net.Listen("tcp", address)
	//
	//if err != nil {
	//	log.Fatalf("Failed to listen: %v", err)
	//}
	//
	//s := grpc.NewServer()
	//pb.RegisterTicketServer(s, &server.server{})
	//
	//log.Println("Start grpc server!")
	//
	//if err := s.Serve(listen); err != nil {
	//	log.Fatalf("Failed to serve: %v", err)
	//}
}