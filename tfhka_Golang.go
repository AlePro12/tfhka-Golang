package tfhka_Golang

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

)

type Tfhka struct {
	err0r            string
	status           string
	StatusError      []string
	resp             string
	socket           string
	service_port     string
	address          string
	lineasProcesadas int
	arrayS1          []string
	arrayS2          string
	arrayS3          string
	arrayS4          []string
	arrayS5          string
	arrayS6          string
	arrayRX          string
	arrayRZ          string
	LastRes          string
	conn             *net.TCPConn
}


//127.0.0.1:PORT Important
func Tfhka_init(address string, service_port string) (Tfhka, bool) {
	var a = Tfhka{"", "", nil, "", "", service_port, address, 0, nil, "", "", nil, "", "", "", "", "", nil}
	fmt.Println("Dial Complete")
	fmt.Println("Dial to " + address + ":" + service_port)
	seconds := 10
	d := net.Dialer{Timeout: time.Duration(seconds) * time.Second}
	conn, err := d.Dial("tcp", a.address+":"+a.service_port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return a, false
	} else {
		a.conn = conn.(*net.TCPConn)
		return a, true
	}

}
func (a Tfhka) SendCmd(cmd string) bool {
	a.StatusError = a.ReadFpStatus()
	time.Sleep(200 * time.Millisecond)
	//a.StatusError[2] si no es 0,3,1
	if a.StatusError[2] == "0" || a.StatusError[2] == "3" || a.StatusError[2] == "1" {
		var in = "SendCmd():" + cmd + "\000"
		LogFile(in, a.address)
		_, err := a.conn.Write([]byte(in))
		errCheck := CheckError(err, a.address)
		if errCheck == false {
			LogFile("Error Reintentando....: Write: "+cmd, a.address)
			time.Sleep(4 * time.Second)
			_, err = a.conn.Write([]byte(in))
			CheckError(err, a.address)
		}
		tmp := make([]byte, 1024)
		_, err = a.conn.Read(tmp)
		errCheck = CheckError(err, a.address)
		if errCheck == false {
			LogFile("Error Reintentando....: "+string(tmp), a.address)
			time.Sleep(4 * time.Second)
			_, err = a.conn.Read(tmp)
			CheckError(err, a.address)
		}
		a.resp = Substr(string(tmp), 10, 1)
		LogFile("Response: "+string(tmp)+" T/f: "+a.resp, a.address)
		if a.resp == "T" {
			return true
		} else {
			a.StatusError = a.ReadFpStatus()
			LogFile("Error Identificado como: "+a.StatusError[0]+" : "+a.StatusError[2]+a.PrettyError(inum(a.StatusError[2])), a.address)
			return false
		}
	} else {
	
		LogFile("Error Identificado como: "+a.StatusError[0]+" : "+a.StatusError[2]+a.PrettyError(inum(a.StatusError[2])), a.address)
		return false
	}
}
func (a Tfhka) CheckFprinter() bool {
	//fmt.Println("Send Check ...")
	var in = "CheckFprinter():1\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err, a.address)
	tmp := make([]byte, 256)
	//fmt.Println("Write Complete. Reading...")
	_, err = a.conn.Read(tmp)
	CheckError(err, a.address)
	//fmt.Println("Readed: " + string(tmp))
	a.resp = Substr(string(tmp), 10, 1)
	if a.resp == "T" {
		return true
	} else {
		return false
	}
}
func (a Tfhka) ReadFpStatus() []string {
	var in = "ReadFpStatus():1\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err, a.address)
	tmp := make([]byte, 1024)
	_, err = a.conn.Read(tmp)
	CheckError(err, a.address)
	LogFile("Response: "+string(tmp), a.address)
	a.StatusError = strings.Split(Substr(string(tmp), 10, utf8.RuneCountInString(string(tmp))), "|")
	return a.StatusError
}
func (a Tfhka) UploadStatus() []string {
	var in = "UploadStatus():S1\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err, a.address)
	tmp := make([]byte, 256)
	_, err = a.conn.Read(tmp)
	CheckError(err, a.address)
	LogFile("Response: "+string(tmp), a.address)
	a.arrayS1 = strings.Split(Substr(string(tmp), 10, utf8.RuneCountInString(string(tmp))), "|")
	return a.arrayS1
}
func (a Tfhka) UploadStatusA(St string) []string {
	var in = "UploadStatus():" + St + "\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err, a.address)
	tmp := make([]byte, 256)
	_, err = a.conn.Read(tmp)
	CheckError(err, a.address)
	LogFile("Response: "+string(tmp), a.address)
	a.arrayS1 = strings.Split(Substr(string(tmp), 10, utf8.RuneCountInString(string(tmp))), "|")
	return a.arrayS1
}
func (a Tfhka) UploadStatusS4() []string {
	/*
	   PHP CODE
	   	$in = "UploadStatus():S1\0";
	   			$out = "";
	   			socket_write($this->socket, $in, strlen($in));
	   			$out = socket_read($this->socket, 1024);
	   			$this->arrayS1 = explode("|",substr($out,10));
	    			return $this->arrayS1;

	*/
	var in = "UploadStatus():S4\000"
	_, err := a.conn.Write([]byte(in))
	CheckError(err, a.address)
	tmp := make([]byte, 256)
	//fmt.Println("Write Complete. Reading...")
	_, err = a.conn.Read(tmp)
	CheckError(err, a.address)
	//fmt.Println("Readed: " + string(tmp))
	//			$this->arrayS1 = explode("|",substr($out,10));
	LogFile("Response: "+string(tmp), a.address)
	a.arrayS4 = strings.Split(Substr(string(tmp), 10, utf8.RuneCountInString(string(tmp))), "|")
	return a.arrayS4
}

/*

 	$in = "UploadStatus():S1\0";
			$out = "";
			socket_write($this->socket, $in, strlen($in));
			$out = socket_read($this->socket, 1024);
			$this->arrayS1 = explode("|",substr($out,10));
 			return $this->arrayS1;



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
//log the commands
func CheckError(err error, ip string) bool {
	if err != nil {
	
		fmt.Fprintf(os.Stderr, "Fatal error: ", err.Error())
		LogFile("Fatal error: "+err.Error(), ip)
		return false
	}
	return true
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
func (a Tfhka) CheckCmdError(CmdValid bool, Exit bool) {
	if CmdValid != true {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", a.LastRes)
		if Exit {
			//os.Exit(1)
		} else {

		}
	}
}

var ExecutePath = GetExePath()

func LogFile(texto string, ip string) {
	var CTime = GetTime()

	//fmt.Println(texto)
	//fmt.Println(time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(ExecutePath+"logs/Tfhka_"+ip+"_"+CTime.Format("2006-01-02")+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		//log.Fatal(err)
	}
	defer f.Close()
	//guardar el texto en el archivo con el formato de fecha y hora f.WriteString
	if _, err = f.WriteString(CTime.Format("2006-01-02 15:04:05") + " | " + texto + "\n"); err != nil {
		//log.Fatal(err)
	}
}
func GetTime() time.Time {
	var loc, err = time.LoadLocation("America/Caracas")
	var CTime time.Time
	if err != nil {
		//fmt.Println(err)
		CTime = time.Now()

	} else {
		CTime = time.Now().In(loc)
	}
	return CTime
}
func GetExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex) + "/"
}
func (a Tfhka) PrettyError(errorcode int) string {
	//si es 0 no hay error
	switch errorcode { //todos los errores son fushi
	case 0:
		return "No hay error."
	case 1:
		return "Fin en la entrega de papel."
	case 2:
		return "Error de índole mecánico en la entrega de papel."
	case 3:
		return "Fin en la entrega de papel y error mecánico."
	case 80:
		return "Comando invalido o valor invalido."
	case 84:
		return "Tasa invalida."
	case 88:
		return "No hay asignadas directivas."
	case 92:
		return "Comando invalido."
	case 96:
		return "Error fiscal."
	case 100:
		return "Error de la memoria fiscal."
	case 108:
		return "Memoria fiscal llena."
	case 112:
		return "Buffer completo. (debe enviar el comando de reinicio)"
	case 128:
		return "Error en la comunicación."
	case 137:
		return "No hay respuesta."
	case 144:
		return "Error LRC."
	case 145:
		return "Error interno api."
	case 153:
		return "Error en la apertura del archivo."
	default:
		return "Error desconocido."
	}

}
func inum(n string) int {
	f, _ := strconv.Atoi(n)
	return f
}
