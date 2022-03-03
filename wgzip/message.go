package wgzip 

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io/ioutil"
	"log"
)

type GateMsg struct {
	Uid string 
	Method int
	Action string
	Data []byte
}

type ServiceMsg struct {
	HttpCode int
	Data []byte
}

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func Compress(s []byte) []byte {
	zipbuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipbuf)
	zipped.Write(s)
	zipped.Close()
	return zipbuf.Bytes()
}

func Decompress(s []byte) []byte {
	rdr, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatal(err)
	}
	rdr.Close()
	return data
}

func DecodeToGateMsg(s []byte) GateMsg {
	p := GateMsg{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func DecodeToServiceMsg(s []byte) ServiceMsg {
	p := ServiceMsg{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

const METHOD_GET = 0;
const METHOD_POST = 1;
const METHOD_PUT = 2;
const METHOD_PATCH = 3;
const METHOD_DELETE = 4;