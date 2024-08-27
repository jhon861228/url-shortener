test:
	go test ./... -v

# Comando para ejecutar los tests con la opción de ver la cobertura
test-cover:
	go test ./... -coverprofile=coverage.out

# Comando para ver la cobertura en el navegador
cover-html: test-cover
	go tool cover -html=coverage.out

# Comando para ejecutar los tests y mostrar la cobertura en la línea de comandos
cover:
	go test ./... -cover

# Comando para limpiar los archivos generados por los tests
clean:
	rm -f coverage.out