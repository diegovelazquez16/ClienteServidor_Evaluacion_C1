package models

type Product struct {
    ID           int    `json:"id"`
    Nombre       string `json:"nombre"`
    Cantidad     int    `json:"cantidad"`
    CodigoBarras string `json:"codigo_barras"`
}