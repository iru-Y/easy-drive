package dir

import (
	"log/slog"
	"os"
	"path/filepath"
)

func AlocDir() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Não foi possível alocar espaço na home")
	}
	dir := filepath.Join(homeDir, "easy_data")
	perm := os.FileMode(0755)
	slog.Info("Permissões de escrita e leitura adicionadas...")
	err = os.MkdirAll(dir, perm)
	if err != nil {
		slog.Error("Não foi possível criar o diretório, verifique se você tem as permissões necessárias")
	} else {
		slog.Info("Iniciando a criação de diretórios...")
	}
}
