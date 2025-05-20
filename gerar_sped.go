package main

import (
	"fmt"
	"os"
	"time"
)

// Estrutura para armazenar os dados do contribuinte
type Contribuinte struct {
	NomeEmpresa         string
	CNPJ                string
	UF                  string
	InscricaoEstadual   string
	CodigoMunicipio     string
	NomeContador        string
	CPFContador         string
	CRC                 string
	CNPJContador        string
	CEPContador         string
	EnderecoContador    string
	NumeroContador      string
	ComplementoContador string
	BairroContador      string
	TelefoneContador    string
	EmailContador       string
}

// Função para gerar o registro 0000
func gerarRegistro0000(c Contribuinte) string {
	dataInicial := time.Now().Format("02012006")
	dataFinal := time.Now().Format("02012006")
	return fmt.Sprintf("|0000|019|0|%s|%s|%s|%s||%s|%s|%s|||A|1|",
		dataInicial, dataFinal, c.NomeEmpresa, c.CNPJ, c.UF, c.InscricaoEstadual, c.CodigoMunicipio)
}

// Função para gerar o registro 0001
func gerarRegistro0001() string {
	return "|0001|0|"
}

// Função para gerar o registro 0005
func gerarRegistro0005() string {
	return "|0005|LAZIO GELATO|70385510|Quadra SHCS CL QD 115, BLOCO A , LOJA 34|34||ASA SUL|6199106718||lazioasasul@gmail.com|"
}

// Função para gerar o registro 0100
func gerarRegistro0100(c Contribuinte) string {
	return fmt.Sprintf("|0100|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s||%s|%s|",
		c.NomeContador, c.CPFContador, c.CRC, c.CNPJContador, c.CEPContador,
		c.EnderecoContador, c.NumeroContador, c.ComplementoContador, c.BairroContador,
		c.TelefoneContador, c.EmailContador, c.CodigoMunicipio)
}

// Função para gerar o registro 0990
func gerarRegistro0990() string {
	return "|0990|5|"
}

// Função para gerar o registro B001
func gerarRegistroB001() string {
	return "|B001|0|"
}

// Função para gerar o registro B470
func gerarRegistroB470() string {
	return "|B470|0|0|0|0|0|0|0|0|0|0|0|0|0|0|"
}

// Função para gerar o registro B990
func gerarRegistroB990() string {
	return "|B990|3|"
}

// Função para gerar o registro C001
func gerarRegistroC001() string {
	return "|C001|1|"
}

// Função para gerar o registro C990
func gerarRegistroC990() string {
	return "|C990|2|"
}

// Função para gerar o registro D001
func gerarRegistroD001() string {
	return "|D001|1|"
}

// Função para gerar o registro D990
func gerarRegistroD990() string {
	return "|D990|2|"
}

// Função para gerar o registro E001
func gerarRegistroE001() string {
	return "|E001|0|"
}

// Função para gerar o registro E100
func gerarRegistroE100() string {
	dataInicial := time.Now().Format("02012006")
	dataFinal := time.Now().Format("02012006")
	return fmt.Sprintf("|E100|%s|%s|", dataInicial, dataFinal)
}

// Função para gerar o registro E110
func gerarRegistroE110() string {
	return "|E110|0|0|0|0|0|0|0|0|368,07|0|0|0|368,07|0|"
}

// Função para gerar o registro E300
func gerarRegistroE300(c Contribuinte) string {
	dataInicial := time.Now().Format("02012006")
	dataFinal := time.Now().Format("02012006")
	return fmt.Sprintf("|E300|%s|%s|%s|", c.UF, dataInicial, dataFinal)
}

// Função para gerar o registro E310
func gerarRegistroE310() string {
	return "|E310|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|"
}

// Função para gerar o registro E990
func gerarRegistroE990() string {
	return "|E990|6|"
}

// Função para gerar o registro G001
func gerarRegistroG001() string {
	return "|G001|1|"
}

// Função para gerar o registro G990
func gerarRegistroG990() string {
	return "|G990|2|"
}

// Função para gerar o registro H001
func gerarRegistroH001() string {
	return "|H001|1|"
}

// Função para gerar o registro H990
func gerarRegistroH990() string {
	return "|H990|2|"
}

// Função para gerar o registro K001
func gerarRegistroK001() string {
	return "|K001|1|"
}

// Função para gerar o registro K990
func gerarRegistroK990() string {
	return "|K990|2|"
}

// Função para gerar o registro 1001
func gerarRegistro1001() string {
	return "|1001|0|"
}

// Função para gerar o registro 1010
func gerarRegistro1010() string {
	return "|1010|N|N|N|N|N|N|N|N|N|N|N|N|N|"
}

// Função para gerar o registro 1990
func gerarRegistro1990() string {
	return "|1990|3|"
}

// Função para gerar o registro 9001
func gerarRegistro9001() string {
	return "|9001|0|"
}

// Função para gerar os registros 9900
func gerarRegistros9900() []string {
	return []string{
		"|9900|0000|1|",
		"|9900|0001|1|",
		"|9900|0005|1|",
		"|9900|0100|1|",
		"|9900|0990|1|",
		"|9900|1001|1|",
		"|9900|1010|1|",
		"|9900|1990|1|",
		"|9900|9001|1|",
		"|9900|9990|1|",
		"|9900|9999|1|",
		"|9900|B001|1|",
		"|9900|B470|1|",
		"|9900|B990|1|",
		"|9900|C001|1|",
		"|9900|C990|1|",
		"|9900|D001|1|",
		"|9900|D990|1|",
		"|9900|E001|1|",
		"|9900|E100|1|",
		"|9900|E110|1|",
		"|9900|E300|1|",
		"|9900|E310|1|",
		"|9900|E990|1|",
		"|9900|G001|1|",
		"|9900|G990|1|",
		"|9900|H001|1|",
		"|9900|H990|1|",
		"|9900|K001|1|",
		"|9900|K990|1|",
		"|9900|9900|31|",
	}
}

// Função para gerar o registro 9990
func gerarRegistro9990() string {
	return "|9990|34|"
}

// Função para gerar o registro 9999
func gerarRegistro9999() string {
	return "|9999|61|"
}

func main() {
	// Criar instância do contribuinte com os dados necessários
	contribuinte := Contribuinte{
		NomeEmpresa:         "LAZIO GELATO PIZZARIA E RESTAURANTE LTDA",
		CNPJ:                "53325471000136",
		UF:                  "DF",
		InscricaoEstadual:   "0826873700133",
		CodigoMunicipio:     "5300108",
		NomeContador:        "KARLA CINTHIAN MEIRA",
		CPFContador:         "07706420639",
		CRC:                 "1-DF-026799/O-5",
		CNPJContador:        "20650846000184",
		CEPContador:         "71200010",
		EnderecoContador:    "Setor SIA, Trecho 01, Lote 630 a 780",
		NumeroContador:      "102",
		ComplementoContador: "SALA",
		BairroContador:      "GUARÁ",
		TelefoneContador:    "61981646410",
		EmailContador:       "karlameiracontadora@gmail.com",
	}

	// Criar arquivo
	arquivo, err := os.Create("sped.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer arquivo.Close()

	// Escrever registros no arquivo
	registros := []string{
		gerarRegistro0000(contribuinte),
		gerarRegistro0001(),
		gerarRegistro0005(),
		gerarRegistro0100(contribuinte),
		gerarRegistro0990(),
		gerarRegistroB001(),
		gerarRegistroB470(),
		gerarRegistroB990(),
		gerarRegistroC001(),
		gerarRegistroC990(),
		gerarRegistroD001(),
		gerarRegistroD990(),
		gerarRegistroE001(),
		gerarRegistroE100(),
		gerarRegistroE110(),
		gerarRegistroE300(contribuinte),
		gerarRegistroE310(),
		gerarRegistroE990(),
		gerarRegistroG001(),
		gerarRegistroG990(),
		gerarRegistroH001(),
		gerarRegistroH990(),
		gerarRegistroK001(),
		gerarRegistroK990(),
		gerarRegistro1001(),
		gerarRegistro1010(),
		gerarRegistro1990(),
		gerarRegistro9001(),
	}

	// Adicionar registros 9900
	registros = append(registros, gerarRegistros9900()...)

	// Adicionar registros finais
	registros = append(registros, gerarRegistro9990(), gerarRegistro9999())

	// Escrever todos os registros no arquivo
	for _, registro := range registros {
		_, err := arquivo.WriteString(registro + "\n")
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}
	}

	fmt.Println("Arquivo SPED gerado com sucesso!")
}
