package main

func (s *server) routes() {
	s.router.HandleFunc("/servers", s.handleGetServers()).Methods("GET")
	s.router.HandleFunc("/servers", s.handlePostServers()).Methods("POST")
	s.router.HandleFunc("/servers/:serverId", s.handleGetServer()).Methods("GET")
	s.router.HandleFunc("/servers/:serverId", s.handlePutServer()).Methods("PUT")
}
