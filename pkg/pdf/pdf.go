package pdf

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/pkg/qr"
	"github.com/go-pdf/fpdf"
)

func GeneratePDF(batch *entities.Batch) (*fpdf.Fpdf, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	mux, wg := new(sync.Mutex), new(sync.WaitGroup)
	wg.Add(batch.AmountQrs)
	for i := 0; i < batch.AmountQrs; i++ {
		url := fmt.Sprintf("http://127.0.0.1:8080/api/batch/%s/%d", batch.ID, i)

		go func(i int, wg *sync.WaitGroup, mux *sync.Mutex, pdf *fpdf.Fpdf) {
			defer wg.Done()
			b, err := qr.GenerateQR(i, url)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			buff := bytes.NewReader(b)
			mux.Lock()
			pdf.AddPage()
			pdf.RegisterImageOptionsReader(fmt.Sprintf("%v.png", i), fpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, buff)
			pdf.Image(fmt.Sprintf("%v.png", i), 10, 10, 0, 0, false, "", 0, "")
			mux.Unlock()
		}(i, wg, mux, pdf)

	}

	wg.Wait()
	return pdf, nil
}
