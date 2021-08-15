package gtls

import (
	"testing"
	"time"
)

func TestCert_GenerateCrtAndKey(t *testing.T) {
	option := &Options{
		KeyType:      KeyTypeRSA,
		Organization: "upatry",
		Duration:     365 * 24 * time.Hour,
		Hosts:        []string{"github.com"},
	}
	cert, err := New(option)
	if err != nil {
		panic(err)
	}

	crtWrapper, keyWrapper, err := cert.GenerateCrtAndKey(option)
	if err != nil {
		panic(err)
	}

	if err = crtWrapper.File("./server.crt"); err != nil {
		panic(err)
	}
	if err = keyWrapper.File("./server.key"); err != nil {
		panic(err)
	}
}

func TestTlsServer(t *testing.T) {
	//http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
	//	writer.Write([]byte("我喜欢你"))
	//})
	//if err := http.ListenAndServeTLS(":8888", "./server.crt", "./server.key", nil); err != nil {
	//	panic(err)
	//}
}
