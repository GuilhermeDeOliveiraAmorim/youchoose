package util

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = oldOutput
	}()

	code := 123
	message := "Teste de mensagem"
	from := "Origem"
	layer := "Camada"
	typeLog := "TipoLog"
	NewLoggerError(code, message, from, layer, typeLog)

	w.Close()
	output, _ := readAll(r)

	var logOutput map[string]interface{}
	err := json.Unmarshal(output, &logOutput)
	if err != nil {
		t.Errorf("Erro ao decodificar a saída JSON: %v", err)
	}
}

func readAll(r *os.File) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	return buf.Bytes(), err
}
