package dir

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func UpFile() {
	AlocDir()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Não foi possível determinar o diretório, verifique se você tem permissões ou se o arquivo é válido")
	}
	var fileName string
	fmt.Println("Digite o nome do arquivo: ")
	fmt.Scanln(fileName)
	filePath := filepath.Join(homeDir, "easy_data", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		slog.Error("Não foi possível criar o arquivo", "error", err)
		return
	}
	defer file.Close()
	var content string
	fmt.Println("Digite o conteúdo do arquivo: ")
	fmt.Scanf(content)
	_, err = file.WriteString(content)
	if err != nil {
		slog.Error("Não foi possível escrever no arquivo", "error", err)
		return
	}
	slog.Info("Arquivo criado com sucesso", "arquivo", filePath)

}
