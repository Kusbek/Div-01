package apiserver

//Options ...
type Options struct {
	Address string
}

//Start ....
func Start(opts Options) error {
	srv := newServer()
	return srv.ListenAndServe(opts.Address)
}
