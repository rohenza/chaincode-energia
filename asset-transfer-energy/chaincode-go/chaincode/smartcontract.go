package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID             string `json:"ID"`
	Consumidor          string `json:"color"`
	Consumo float64 `json:"consumo"`
	Valor float64 `json:"valor"`
	Estado string `json:"estado"`
	Numero string `json:"numero"`
	DataAbertura string `json:"dataAbertura"`
	DataFechamento string `json:"dataFechamento"`
}

// Define as características do Meter.
type Meter struct {
	Id string `json:"id"`
	IdType string `json:"idType"`
	IdNameSpace string `json:"idNameSpace"`
}

// Define o bloco de cada medida.
type iReading struct {
	EndTime string `json:"endTime"`
	Value float64 `json:"value"`
	Flags string `json:"flags"`
}

// Define a o bloco de intervalo de medição.
type Medicao struct {
	Meter   Meter `json:"meter"`
	ReadingTypeId  string `json:"readingTypeId"`
	iReading iReading `json:"iReading"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "fatura1", Consumidor: "Rodrigo", Consumo: 15.5, Valor: 65.0, Estado: "Aberta", Numero: "8453244", DataAbertura: "01/07/2020", DataFechamento: ""},
		{ID: "fatura2", Consumidor: "Thais", Consumo: 12, Valor: 60.0, Estado: "Aberta", Numero: "756456456", DataAbertura: "01/07/2020", DataFechamento: ""},
		{ID: "fatura3", Consumidor: "Marcelo", Consumo: 13, Valor: 61.0, Estado: "Aberta", Numero: "4564567", DataAbertura: "01/07/2020", DataFechamento: ""},
		{ID: "fatura4", Consumidor: "Maria", Consumo: 14, Valor: 62.0, Estado: "Aberta", Numero: "46456453", DataAbertura: "01/07/2020", DataFechamento: ""},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, consumidor string, consumo float64, valor float64, estado string, numero string, dataAbertura string, dataFechamento string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID: id,
		Consumidor: consumidor,
		Consumo: consumo, 
		Valor: valor, 
		Estado: estado, 
		Numero: numero, 
		DataAbertura: dataAbertura, 
		DataFechamento: dataFechamento,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset *Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, consumidor string, consumo float64, valor float64, estado string, numero string, dataAbertura string, dataFechamento string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ID: id,
		Consumidor: consumidor,
		Consumo: consumo, 
		Valor: valor, 
		Estado: estado, 
		Numero: numero, 
		DataAbertura: dataAbertura, 
		DataFechamento: dataFechamento,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func (s *SmartContract) CreateMedicao(ctx contractapi.TransactionContextInterface, id string, idFatura string, idMeter string, idType string, idNameSpace string, endTime string, flags string, value float64, readingTypeId string) error {
	var meter = Meter{Id: idMeter, IdType: idType, IdNameSpace: idNameSpace}
	
	var iReading = iReading{EndTime: endTime, Flags: flags, Value: value}

	var medicao = Medicao{Meter: meter, iReading: iReading, ReadingTypeId: readingTypeId}

	medicaoAsBytes, _ := json.Marshal(medicao)
	ctx.GetStub().PutState(id, medicaoAsBytes)

	faturaAsBytes, _ := ctx.GetStub().GetState(idFatura)
	fatura := Asset{}

	json.Unmarshal(faturaAsBytes, &fatura)
	 	
	fatura.Consumo = fatura.Consumo + value

	faturaAsBytes, _ = json.Marshal(fatura)

	return ctx.GetStub().PutState(id, faturaAsBytes)

}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset *Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}
