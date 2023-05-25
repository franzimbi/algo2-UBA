package procesamiento

// este paquete son funciones usadas para procesar todo lo que recibe el programa
// no es un TDA ni un encapsulamiento de funciones y/o datos
// simplemente es para hacer mas leible y entendible el programa desarrollado
import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	Errores "tp1/errores"
	. "tp1/votos"
)

const (
	TIPO_VOTO_INVALIDO       = -1
	_POSICION_NOMBRE_PARTIDO = 0
	_POSICION_PRESIDENTE     = 1
	_POSICION_GOBERNADOR     = 2
	_POSICION_INTENDENTE     = 3
	_CANTIDAD_POSTULANTES    = 3
)

func ProcesarLinea(linea *bufio.Scanner) (parametros []string) {
	lineaLeida := linea.Text()
	return strings.Split(lineaLeida, " ")
}

func LeerCsvPartidos(archivo string) ([]Partido, error) {
	csvPartidos, err := os.Open(archivo)
	if err != nil {
		csvPartidos.Close()
		return nil, &Errores.ErrorLeerArchivo{}
	}
	PartidosReader := csv.NewReader(csvPartidos)
	PartidosReader.LazyQuotes = true
	var partidos []Partido
	partidos = append(partidos, CrearVotosEnBlanco())
	for {
		rec, err := PartidosReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			csvPartidos.Close()
			return nil, &Errores.ErrorLeerArchivo{}
		}
		partidos = append(partidos, CrearPartido(rec[_POSICION_NOMBRE_PARTIDO], [_CANTIDAD_POSTULANTES]string{rec[_POSICION_PRESIDENTE], rec[_POSICION_GOBERNADOR], rec[_POSICION_INTENDENTE]}))
	}
	csvPartidos.Close()
	return partidos, nil
}

func LeerTxtPadron(archivo string) ([]Votante, error) {
	txtPadron, err := os.Open(archivo)
	if err != nil {
		txtPadron.Close()
		return nil, &Errores.ErrorLeerArchivo{}
	}
	padronReader := bufio.NewScanner(txtPadron)
	padronReader.Split(bufio.ScanLines)
	var padronInt []int

	for padronReader.Scan() {
		dni, err := strconv.Atoi(padronReader.Text())
		if err != nil {
			txtPadron.Close()
			return nil, &Errores.ErrorLeerArchivo{}
		}
		padronInt = append(padronInt, dni)
	}
	txtPadron.Close()
	sort.Ints(padronInt)
	padron := make([]Votante, len(padronInt))
	for i := 0; i < len(padronInt); i++ {
		padron[i] = CrearVotante(padronInt[i])
	}
	return padron, nil
}

func busquedaBinariaPadron(padron []Votante, pos, dni int) int {
	tam := len(padron)
	if tam <= 0 {
		return -1
	}
	medio := tam / 2
	dniAux := padron[medio].LeerDNI()
	if dniAux == dni {
		return pos + medio
	}
	if dniAux < dni {
		return busquedaBinariaPadron(padron[medio+1:], pos+medio+1, dni)
	} else {
		return busquedaBinariaPadron(padron[:medio], pos, dni)
	}
}

func PadronValidar(padron []Votante, dni string) (int, error) {
	nroDni, err := strconv.Atoi(dni)
	if err != nil || nroDni <= 0 {
		return 0, &Errores.DNIError{}
	}
	pos := busquedaBinariaPadron(padron, 0, nroDni)
	if pos == -1 {
		return -1, &Errores.DNIFueraPadron{}
	} else {
		return pos, nil
	}
}
func tipoVotoAEnum(tipo string) TipoVoto {
	switch tipo {
	case "Presidente":
		return PRESIDENTE
	case "Gobernador":
		return GOBERNADOR
	case "Intendente":
		return INTENDENTE
	default:
		return TIPO_VOTO_INVALIDO
	}
}

func VotoValidar(cantPartidos int, voto string, nroLista string) (TipoVoto, int, error) {
	tipoVoto := tipoVotoAEnum(voto)
	if tipoVoto == TIPO_VOTO_INVALIDO {
		return TIPO_VOTO_INVALIDO, TIPO_VOTO_INVALIDO, &Errores.ErrorTipoVoto{}
	}
	nroListaInt, err := strconv.Atoi(nroLista)
	if err != nil || nroListaInt < 0 || nroListaInt >= cantPartidos {
		return TIPO_VOTO_INVALIDO, TIPO_VOTO_INVALIDO, &Errores.ErrorAlternativaInvalida{}
	}
	return tipoVoto, nroListaInt, nil
}
