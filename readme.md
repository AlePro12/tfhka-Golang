#Tfhka_Golang
****Integracion de impresora fiscal Tfhka Venezuela mediante el uso de Sockets.****

Ya esta API esta programada en PHP pero aqui esta en Golang.

Es cliente Servidor y enviara los datos por sockets TCP a la PC donde 
esta instalada la impresora con DPOS_SOCKET_P.exe corriendo.

*NOTA: Por ahora es una BETA esta en pruebas*

#USO
**Ejemplo basico de conexion con el socket**

    ip="192.168.0.2" // Direccion de la PC que tiene la impresora fiscal y que este corriendo DPOS_SOCKET.exe
    Thka, Is_connected := tfhka_Golang.Tfhka_init(ip, "8090")
	if Is_connected == true {
	Result := Thka.CheckFprinter()
    if Result == true {
    fmt.Println("Impresora lista para recibir un comando")
    }else{
    fmt.Println("La impresora no esta disp.")
    }
    }else{
    fmt.Println("No se pudo conectar con la Impresora")
    }
Es necesario que el host este ejecutando DPOS_SOCKET.exe ya que este es el que hace la conexion entre estas.

**Ejemplo de envio de comandos (Imprimir Factura Fiscal)**

    ip="192.168.0.2" // Direccion de la PC que tiene la impresora fiscal y que este corriendo DPOS_SOCKET.exe
    Thka, Is_connected := tfhka_Golang.Tfhka_init(ip, "8090")
	if Is_connected == true {
	Result := Thka.CheckFprinter()
    if Result == true {
    var CmdValid bool
    CmdValid = Thka.SendCmd("i01NOMBRE:ALEJANDRO SANCHEZ")
    CmdValid = Thka.SendCmd("i02CEDULA:V-9999999")
    CmdValid = Thka.SendCmd("i03DIRECCION:CABIMAS")
    CmdValid = Thka.SendCmd("!011431034500001000PEPITO XXL 180GR")
    CmdValid = Thka.SendCmd("!012138793100001000SERVILLETA SARY X 150UNI")
    CmdValid = Thka.SendCmd("101")
    }else{
    fmt.Println("La impresora no esta disp.")
    }
    }else{
    fmt.Println("No se pudo conectar con la Impresora")
    }
