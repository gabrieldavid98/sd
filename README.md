# Sistemas Distribuidos

Este proyecto muestra como se calcula el número PI usando la fórmula de Leibniz, el proyecto permite hacer el cálculo sin hacer uso de hilos, haciendo uso de hilos y finalmente haciendo uso de otros nodos para de manera distribuida calcular cada uno de los componentes necesarios para encontrar el número PI

## Endpoints
| Endpoint                           | Parámetros                                                                                            | Descripción                                                                                     |
|------------------------------------|-------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| `GET /nodes`                         |                                                                                                       | Retorna la cantidad de nodos registrados                                                        |
| `GET /compute-pi-single/:limit`      | `limit`: representa la cantidad de iteraciones  que se tiene que hacer (Valor por defecto: 90 millones) | Retorna el número PI calculado sin usar hilos                                                   |
| `GET /compute-pi-multicore/:limit`   | `limit`: representa la cantidad de iteraciones que se tiene que hacer (Valor por defecto: 90 millones)  | Retorna el número PI calculado usando hilos                                                     |
| `GET /compute-pi-distributed/:limit` | `limit`: representa la cantidad de iteraciones que se tiene que hacer (Valor por defecto: 90 millones)  | Retorna el número PI usando nodos externos para distribuir la carga del cálculo de dicho número |

## Para ejecutar el proyecto
El proyecto corre sobre docker usando docker compose para tener un cluster de nodos. A continuación se muestran las instrucciones para correr el proyecto en Windows, Linux/Mac

### Requerimientos
- Docker
- Docker Compose

### Windows
El proyecto se puede correr en esta plataforma, usando PowerShell

#### PowerShell (requiere tener activo la ejecución de scripts de powershell)
En la carpeta raíz del proyecto ejecutar el siguiente comando en una sesión de PowerShell:
```
.\build.ps1 -Run
```

#### Manualmente
En la carpeta raíz del proyecto ejecutar el siguiente comando en una terminal:
```
docker-compose up --force-recreate --build
```

### Linux/Mac
Para correr el proyecto en estas plataformas solo hay que correr manualmente el siguiente comando

#### Manualmente
En la carpeta raíz del proyecto ejecutar el siguiente comando en una terminal:
```
docker-compose up --force-recreate --build
```