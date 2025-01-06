package dir

import (
	"log/slog"
	"os"
	"path/filepath"
)

func UpFile(fileName, content string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Não foi possível determinar o diretório, verifique se você tem permissões ou se o arquivo é válido")
	}

	filePath := filepath.Join(homeDir, "easy_data", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		slog.Error("Não foi possível criar o arquivo", "error", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		slog.Error("Não foi possível escrever no arquivo", "error", err)
		return
	}
	slog.Info("Arquivo criado com sucesso", "arquivo", filePath)

}
