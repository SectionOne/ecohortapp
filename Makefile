BINARY_NAME=EcoHortApp.exe
APP_NAME=EcoHortApp
VERSION=1.0.1
BUILD_NO=2
APP_ID=cat.cibernarium.ecohortapp.preferences

## construim: construim el binari i paquet de la app
build:
	del ${BINARY_NAME}
	fyne package -appID ${APP_ID} -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release

## execució: construim i executem l'aplicació
run:
	env DB_PATH="./sql.db" go run .

## neteja: executa go clean i borra els binaris
clean:
	@echo "Borrant..."
	@go clean
	@del ${BINARY_NAME}
	@echo "Borrat!"

## test: executa tots els tests
test:
	go test -v ./...