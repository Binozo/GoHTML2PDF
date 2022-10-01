package chrome

import (
	"os"
	"os/exec"
)

const CHROME = "/usr/bin/chromium-browser"

func ConvertHTMLToPDF(html string) ([]byte, error) {
	file, err := os.Create("/source.html")
	if err != nil {
		return nil, err
	}

	if _, err := file.Write([]byte(html)); err != nil {
		return nil, err
	}
	cmd := exec.Command(CHROME, "--headless", "--no-sandbox", "--disable-gpu", "--no-margins", "--print-to-pdf-no-header", "--print-to-pdf=result.pdf", "file:///source.html")
	_, err = cmd.Output()

	if err != nil {
		return nil, err
	}

	os.Remove("/source.html")
	dat, err := os.ReadFile("/result.pdf")
	os.Remove("/result.pdf")

	return dat, err
}
