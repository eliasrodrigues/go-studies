package models

type Telefone struct {
	DDD    string `json:"ddd"`
	Numero string `json:"numero"`
	IsFax  bool   `json:"is_fax"`
}

type Socio struct {
	NomeSocio            string `json:"nome_socio"`
	CnpjCpfSocio         string `json:"cnpj_cpf_socio"`
	QualificacaoSocio    string `json:"qualificacao_socio"`
	DataEntradaSociedade string `json:"data_entrada_sociedade"`
	IdentificadorSocio   string `json:"identificador_socio"`
	FaixaEtaria          string `json:"faixa_etaria"`
}

type Empresa struct {
	CNPJ                  string     `json:"cnpj"`
	RazaoSocial           string     `json:"razao_social"`
	NomeFantasia          string     `json:"nome_fantasia"`
	SituacaoCadastral     string     `json:"situacao_cadastral"`
	DataSituacaoCadastral string     `json:"data_situacao_cadastral"`
	MatrizFilial          string     `json:"matriz_filial"`
	DataInicioAtividade   string     `json:"data_inicio_atividade"`
	CNAEPrincipal         string     `json:"cnae_principal"`
	CNAEsSecundarios      []string   `json:"cnaes_secundarios"`
	CNAEsSecundariosCount int        `json:"cnaes_secundarios_count"`
	NaturezaJuridica      string     `json:"natureza_juridica"`
	Logradouro            string     `json:"logradouro"`
	Numero                string     `json:"numero"`
	Complemento           string     `json:"complemento"`
	Bairro                string     `json:"bairro"`
	CEP                   string     `json:"cep"`
	UF                    string     `json:"uf"`
	Municipio             string     `json:"municipio"`
	Email                 string     `json:"email"`
	Telefones             []Telefone `json:"telefones"`
	CapitalSocial         string     `json:"capital_social"`
	PorteEmpresa          string     `json:"porte_empresa"`
	OpcaoSimples          *bool      `json:"opcao_simples"`
	DataOpcaoSimples      *string    `json:"data_opcao_simples"`
	OpcaoMEI              *bool      `json:"opcao_mei"`
	DataOpcaoMEI          *string    `json:"data_opcao_mei"`
	QSA                   []Socio    `json:"QSA"`
}
