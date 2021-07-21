package tfhka_Golang

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Tfhka struct {
	err0r            string
	status           string
	StatusError      string
	resp             string
	socket           string
	service_port     string
	address          string
	lineasProcesadas int
	arrayS1          string
	arrayS2          string
	arrayS3          string
	arrayS4          string
	arrayS5          string
	arrayS6          string
	arrayRX          string
	arrayRZ          string
	conn             *net.TCPConn
}

//127.0.0.1:PORT Important
func Tfhka_init(address string, service_port string) (Tfhka, bool) {
	var a = Tfhka{"", "", "", "", "", service_port, address, 0, "", "", "", "", "", "", "", "", nil}
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", a.address+":"+a.service_port)
	//CheckError(err)
	//fmt.Println("Loading...")
	//conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	//CheckError(err)
	fmt.Println("Dial Complete")
	seconds := 5
	d := net.Dialer{Timeout: time.Duration(seconds) * time.Second}
	conn, err := d.Dial("tcp", a.address+":"+a.service_port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		a.conn = conn.(*net.TCPConn)
		return a, false
		//os.Exit(1)
	} else {
		a.conn = conn.(*net.TCPConn)
		return a, true
	}

}
func (a Tfhka) SendCmd(cmd string) bool {
	//fmt.Println("Send Command: (" + cmd + ") ...")
	var in = "SendCmd():" + cmd + "\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err)
	tmp := make([]byte, 256)
	//fmt.Println("Write Complete. Reading...")
	_, err = a.conn.Read(tmp)
	CheckError(err)
	//fmt.Println("Readed: " + string(tmp))
	a.resp = Substr(string(tmp), 10, 1)
	if a.resp == "T" {
		return true
	} else {
		return false
	}
}
func (a Tfhka) CheckFprinter(cmd string) bool {
	//fmt.Println("Send Check ...")
	var in = "CheckFprinter():1\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err)
	tmp := make([]byte, 256)
	//fmt.Println("Write Complete. Reading...")
	_, err = a.conn.Read(tmp)
	CheckError(err)
	//fmt.Println("Readed: " + string(tmp))
	a.resp = Substr(string(tmp), 10, 1)
	if a.resp == "T" {
		return true
	} else {
		return false
	}
}

/*
function SendCmd($cmd = "")
{Read()
	$in = "SendCmd():".$cmd."\0";
	$out="";
	socket_write($this->socket, $in, strlen($in));
	$out = socket_read($this->socket, 1024);
	$this->resp= substr($out,10,1);
	if($this->resp==="T"){
		return true;
	}else{
		return false;
	}
}CheckFprinter():1
*/
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
