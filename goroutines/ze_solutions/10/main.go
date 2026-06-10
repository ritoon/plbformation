package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

const (
	baseURL = "https://www.roumet.com/photos/574/"
)

func main() {
	os.Mkdir("tmp", 0777)
	os.Mkdir("processed", 0777)
	processImage(3, 10)
}

func processImage(nbWorkers, nbLotsToProcess int) {
	var wg sync.WaitGroup
	wg.Add(nbWorkers)

	lotsPerWorker := nbLotsToProcess / nbWorkers
	log.Printf("each worker will process %d lots\n", lotsPerWorker)
	// start workers
	for i := 0; i < nbWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			log.Printf("worker %d started\n", workerID)
			processLot(nbLotsToProcess)
			log.Printf("worker %d finished\n", workerID)
		}(i)
	}

	wg.Wait()
}

func processLot(nbLotsToProcess int) {
	// for each lot
	for i := 1; i < nbLotsToProcess; i++ {
	subLoopForSubImage:
		for j := 0; j < 10; j++ {
			imageName := fmt.Sprintf("%d.jpg", i)
			if j > 0 {
				imageName = fmt.Sprintf("%d-%d.jpg", i, j)
			}
			errTask := make(chan error)
			go func() {
				// get image
				data, err := getImage(baseURL + imageName)
				if err != nil {
					errTask <- err
					return
				}
				if data == nil {
					errTask <- fmt.Errorf("no data")
					return
				}
				// save image
				err = os.WriteFile(fmt.Sprintf("tmp/%v", imageName), data, 0666)
				if err != nil {
					errTask <- err
					return
				}
				// add watermark
				err = watermarkImage(fmt.Sprintf("tmp/%v", imageName), "img/logo.png", fmt.Sprintf("processed/%v", imageName))
				if err != nil {
					errTask <- err
					return
				}
			}() // to create a goroutine for the trace
			select {
			case err := <-errTask:
				if err != nil {
					if err.Error() == "no data" {
						break subLoopForSubImage
					}
					log.Println(err)
				} else {
					log.Printf("image %v processed\n", imageName)
				}
			}
			close(errTask)
		}
	}
}

func getImage(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func watermarkImage(imagePath, watermarkPath, resPath string) error {
	imgOriginal, err := mergi.Import(impexp.NewFileImporter(imagePath))
	if err != nil {
		return err
	}
	// Get the image content by passing image path url or file path
	imgWatermark, err := mergi.Import(impexp.NewFileImporter(watermarkPath))
	if err != nil {
		return err
	}

	// Let's position the watermark left top corner
	p := image.Pt(0, 0)

	resultImage, err := mergi.Watermark(imgWatermark, imgOriginal, p)
	if err != nil {
		return err
	}

	// Let's save the image
	err = mergi.Export(impexp.NewFileExporter(resultImage, resPath))
	if err != nil {
		return err
	}
	return nil
}
