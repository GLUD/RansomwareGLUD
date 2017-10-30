/* _Defer_ se usa para asegurar que una función es llamada posteriormente durante la ejecución del programa,
  generalmente con propósitos de limpieza. `Defer` se usa regularmente donde en otros lenguajes donde se utilizaría `ensure` y `finally`
*/
package main

import "fmt"
import "os"
import "log"

func main() {

  /* Con 'createFile' Se obtiene el objeto archivo, se DIFIERE el cierre del archivo con closeFile.
      Esto último se ejecutará después de que 'writeFile' terminó de ejecutarse (al término de la función main). */
  f := createFile("/tmp/defer.txt")
  defer closeFile(f)
  writeFile(f)
  /* Cambiando los permisos del archivo */
  if err := os.Chmod("/tmp/defer.txt", 0777); err != nil {
		log.Fatal(err)
	}
  fmt.Println("Permisos del Archivo cambiados")
}

/* Crear un archivo usando defer */
func createFile(p string) *os.File {
  f, err := os.Create(p)
  fmt.Println("Archivo "+p+" creado")
  if err != nil {
    panic(err)
  }
  return f
}

func writeFile(f *os.File) {
  fmt.Println("En el archivo se está escribiendo ")
  fmt.Fprintln(f, "Writting")
}

func closeFile(f *os.File) {
  fmt.Println("Cerrando el archivo ")
  f.Close()
}
